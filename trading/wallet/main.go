package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/google/uuid"
	pb "github.com/temporalio/samples-go/trading/proto/wallet"
	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 50052, "The server port")
	balance = flag.Float64("balance", 100000, "The IDR balance")
)

type server struct {
	pb.UnimplementedWalletServiceServer
	balances     map[string]float64
	transactions map[string]*pb.CreateTradingRequest
}

func NewServer(idr float64) *server {
	balances := map[string]float64{"idr": idr}
	return &server{
		balances:     balances,
		transactions: make(map[string]*pb.CreateTradingRequest),
	}
}

func (s *server) getBalance(assetType string) float64 {
	assetType = strings.ToLower(assetType)
	if balance, ok := s.balances[assetType]; ok {
		return balance
	} else {
		s.balances[strings.ToLower(assetType)] = 0
		return 0
	}
}

func (s *server) setBalance(assetType string, balance float64) {
	assetType = strings.ToLower(assetType)
	s.balances[assetType] = balance
}

func (s *server) validateBalance(in *pb.CreateTradingRequest) error {
	switch in.Side {
	case pb.TradingSide_BUY:
		quoteBalance := s.getBalance(in.QuoteAssetType)
		quoteAmount, err := strconv.ParseFloat(in.QuoteAmount, 64)
		if err != nil {
			return err
		}
		if quoteBalance < quoteAmount {
			return errors.New("insufficient balance")
		}
	case pb.TradingSide_SELL:
		baseBalance := s.getBalance(in.BaseAssetType)
		baseAmount, err := strconv.ParseFloat(in.BaseAmount, 64)
		if err != nil {
			return err
		}
		if baseBalance < baseAmount {
			return errors.New("insufficient balance")
		}
	}
	return nil
}

func (s *server) CreateTradingTransaction(ctx context.Context, in *pb.CreateTradingRequest) (*pb.CreateTradingResponse, error) {
	if err := s.validateBalance(in); err != nil {
		return nil, err
	}

	transactionId := uuid.NewString()
	s.transactions[transactionId] = in

	return &pb.CreateTradingResponse{
		TransactionId: transactionId,
		Status:        pb.TransactionStatus_PENDING,
	}, nil
}

func (s *server) ConfirmTransaction(ctx context.Context, in *pb.ConfirmTransactionRequest) (*pb.ConfirmTransactionResponse, error) {
	transaction, ok := s.transactions[in.TransactionId]
	if !ok {
		return nil, fmt.Errorf("transaction with id %s not found", in.TransactionId)
	}
	if err := s.validateBalance(transaction); err != nil {
		return nil, err
	}
	return &pb.ConfirmTransactionResponse{
		TransactionId: in.TransactionId,
		Status:        pb.TransactionStatus_SUCCESS,
	}, nil
}

func (s *server) RejectTransaction(ctx context.Context, in *pb.ConfirmTransactionRequest) (*pb.ConfirmTransactionResponse, error) {
	return &pb.ConfirmTransactionResponse{
		TransactionId: in.TransactionId,
		Status:        pb.TransactionStatus_FAILED,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := NewServer(*balance)
	pb.RegisterWalletServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
