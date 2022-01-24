package main

import (
	"log"

	"github.com/temporalio/samples-go/trading"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	tradingHost = "localhost:50051"
	walletHost  = "localhost:50052"
	xcHost      = "localhost:50053"
)

var grpcConns []*grpc.ClientConn = make([]*grpc.ClientConn, 0)

func main() {
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	tradingConn, err := grpc.Dial(tradingHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection: %v\n", err)
	}
	grpcConns = append(grpcConns, tradingConn)

	walletConn, err := grpc.Dial(walletHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection: %v\n", err)
	}
	grpcConns = append(grpcConns, walletConn)

	xcConn, err := grpc.Dial(xcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection: %v\n", err)
	}
	grpcConns = append(grpcConns, xcConn)

	defer (func() {
		c.Close()
		for _, conn := range grpcConns {
			conn.Close()
		}
	})()

	tradingActivities := trading.NewTradingActivities(tradingConn, walletConn, xcConn)

	w := worker.New(c, "trading", worker.Options{})

	w.RegisterWorkflow(trading.TradingWorkflow)
	w.RegisterActivity(tradingActivities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
