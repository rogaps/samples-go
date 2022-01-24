package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/temporalio/samples-go/trading"
	pb "github.com/temporalio/samples-go/trading/proto/trading"
	"github.com/temporalio/samples-go/trading/proto/xc"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedTradingServiceServer
	xcClient       xc.XcServiceClient
	redis          *redis.Client
	temporalClient client.Client
}

func NewServer(xcConn *grpc.ClientConn, redis *redis.Client, temporalClient client.Client) *server {
	return &server{
		xcClient:       xc.NewXcServiceClient(xcConn),
		redis:          redis,
		temporalClient: temporalClient,
	}
}

func (s *server) CreateTrading(ctx context.Context, in *pb.CreateTradingRequest) (*pb.CreateTradingResponse, error) {
	exchangeAmount, err := s.xcClient.GetExchangeAmount(ctx, &xc.GetExchangeAmountRequest{
		BaseAssetType:   in.BaseAssetType,
		QuoteAssetType:  in.QuoteAssetType,
		Side:            xc.TradingSide(in.Side),
		Amount:          in.Amount,
		AmountAssetType: in.AmountAssetType,
	})
	if err != nil {
		return nil, err
	}

	transactionId := uuid.NewString()
	retryPolicy := &temporal.RetryPolicy{
		NonRetryableErrorTypes: []string{"NonRetryableError", "create trading"},
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:          "trading_" + transactionId,
		TaskQueue:   "trading",
		RetryPolicy: retryPolicy,
	}

	res := &pb.CreateTradingResponse{
		TransactionId:   transactionId,
		BaseAssetType:   in.BaseAssetType,
		QuoteAssetType:  in.QuoteAssetType,
		Side:            in.Side,
		Amount:          exchangeAmount.BaseAmount,
		AmountAssetType: in.BaseAssetType,
		Total:           exchangeAmount.QuoteAmount,
		TotalAssetType:  in.QuoteAssetType,
	}

	we, err := s.temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, trading.TradingWorkflow, res)
	if err != nil {
		return nil, err
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	return res, nil
}

func (s *server) CommitTrading(ctx context.Context, in *pb.CommitTradingRequest) (*pb.CommitTradingResponse, error) {
	taskToken, err := s.redis.Get(ctx, fmt.Sprintf("transaction:%s:task-token", in.TransactionId)).Result()
	if err != nil {
		return nil, err
	}
	err = s.temporalClient.CompleteActivity(ctx, []byte(taskToken), nil, nil)
	if err != nil {
		log.Fatalf("Failed to complete activity with error: %s", err)
		return nil, err
	}
	return &pb.CommitTradingResponse{
		TransactionId: in.TransactionId,
		Status:        pb.TransactionStatus_PENDING,
	}, nil
}

func (s *server) RegisterActivity(ctx context.Context, in *pb.RegisterActivityRequest) (*pb.RegisterActivityResponse, error) {
	err := s.redis.Set(ctx, fmt.Sprintf("transaction:%s:task-token", in.Id), in.TaskToken, 0).Err()
	if err != nil {
		return nil, err
	}
	return &pb.RegisterActivityResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// grpc connection
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection: %v", err)
	}

	// redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// temporal client
	temporalClient, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalf("failed to create temporal client: %v", err)
	}

	defer (func() {
		conn.Close()
		rdb.Close()
		temporalClient.Close()
	})()

	server := NewServer(conn, rdb, temporalClient)
	pb.RegisterTradingServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
