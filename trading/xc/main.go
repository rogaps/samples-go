package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/google/uuid"
	pb "github.com/temporalio/samples-go/trading/proto/xc"
	"google.golang.org/grpc"
)

var (
	port   = flag.Int("port", 50053, "The server port")
	rate   = flag.Float64("rate", 15000, "The exchange rate")
	spread = flag.Float64("spread", 0.01, "The spread between buy and sell in percentage")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedXcServiceServer
}

func CalculateQuoteAmount(baseAmountStr string, side pb.TradingSide, rate float64, spread float64) (string, error) {
	baseAmount, err := strconv.ParseFloat(baseAmountStr, 64)
	if err != nil {
		return "", err
	}
	quoteAmount := baseAmount * rate
	halfSpread := spread / 2
	switch side {
	case pb.TradingSide_BUY:
		quoteAmount = quoteAmount - (halfSpread * quoteAmount)
	case pb.TradingSide_SELL:
		quoteAmount = quoteAmount + (halfSpread * quoteAmount)
	default:
		return "", errors.New("invalid side")
	}
	return fmt.Sprintf("%f", quoteAmount), nil
}

func CalculateBaseAmount(quoteAmountStr string, side pb.TradingSide, rate float64, spread float64) (string, error) {
	quoteAmount, err := strconv.ParseFloat(quoteAmountStr, 64)
	if err != nil {
		return "", err
	}
	baseAmount := quoteAmount / rate
	halfSpread := spread / 2
	switch side {
	case pb.TradingSide_BUY:
		baseAmount = baseAmount - (halfSpread * baseAmount)
	case pb.TradingSide_SELL:
		baseAmount = baseAmount + (halfSpread * baseAmount)
	default:
		return "", errors.New("invalid side")
	}
	return fmt.Sprintf("%f", baseAmount), nil
}

// GetExchangeAmount implements XcService.GetExchangeAmount
func (s *server) GetExchangeAmount(ctx context.Context, in *pb.GetExchangeAmountRequest) (out *pb.GetExchangeAmountResponse, err error) {
	var baseAmount, quoteAmount string
	switch in.AmountAssetType {
	case in.BaseAssetType:
		baseAmount = in.Amount
		quoteAmount, err = CalculateQuoteAmount(in.Amount, in.Side, *rate, *spread)
		if err != nil {
			return out, err
		}
	case in.QuoteAssetType:
		quoteAmount = in.Amount
		baseAmount, err = CalculateBaseAmount(in.Amount, in.Side, *rate, *spread)
		if err != nil {
			return out, err
		}
	default:
		return out, errors.New("invalid amount asset type")
	}

	out = &pb.GetExchangeAmountResponse{
		BaseAssetType:  in.BaseAssetType,
		QuoteAssetType: in.QuoteAssetType,
		Side:           in.Side,
		BaseAmount:     baseAmount,
		QuoteAmount:    quoteAmount,
	}
	return
}

func (s *server) PlaceOrder(ctx context.Context, in *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	return &pb.PlaceOrderResponse{
		OrderId: uuid.NewString(),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterXcServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
