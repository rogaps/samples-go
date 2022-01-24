package trading

import (
	"context"
	"errors"
	"log"

	"github.com/temporalio/samples-go/trading/proto/trading"
	"github.com/temporalio/samples-go/trading/proto/wallet"
	"github.com/temporalio/samples-go/trading/proto/xc"
	"go.temporal.io/sdk/activity"
	"google.golang.org/grpc"
)

type TradingRequest struct {
	BaseAssetType  string
	QuoteAssetType string
	Side           string
	AssetType      string
	Amount         string
}

type AssetAmount struct {
	AssetType string
	Amount    string
}

type TradingResponse struct {
	BaseAssetType  string
	QuoteAssetType string
	Side           string
	Amount         AssetAmount
	SubTotal       AssetAmount
	Total          AssetAmount
}

type Transaction struct {
	TransactionID string
}

type TradingActivities struct {
	tradingClient trading.TradingServiceClient
	walletClient  wallet.WalletServiceClient
	xcClient      xc.XcServiceClient
}

func NewTradingActivities(tradingConn *grpc.ClientConn, walletConn *grpc.ClientConn, xcConn *grpc.ClientConn) *TradingActivities {
	return &TradingActivities{
		tradingClient: trading.NewTradingServiceClient(tradingConn),
		walletClient:  wallet.NewWalletServiceClient(walletConn),
		xcClient:      xc.NewXcServiceClient(xcConn),
	}
}

func (a *TradingActivities) CreateTrading(ctx context.Context, request *trading.CreateTradingRequest) error {
	return NewNonRetryableError("create trading")
}

func (a *TradingActivities) WaitForCommitTrading(ctx context.Context, transactionId string) error {
	if len(transactionId) == 0 {
		return errors.New("transaction id is empty")
	}
	logger := activity.GetLogger(ctx)

	// register activity to trading-server
	activityInfo := activity.GetInfo(ctx)
	req := &trading.RegisterActivityRequest{
		Id:        transactionId,
		TaskToken: string(activityInfo.TaskToken),
	}
	_, err := a.tradingClient.RegisterActivity(ctx, req)
	if err != nil {
		logger.Warn("WaitForCommitTrading failed to register activity.", "Error", err)
	}

	logger.Info("WaitForCommitTrading successfully registered activity.", "TransactionId", transactionId)
	return activity.ErrResultPending
}

func (a *TradingActivities) CreateWalletTransaction(ctx context.Context, request *wallet.CreateTradingRequest) (*wallet.CreateTradingResponse, error) {
	logger := activity.GetLogger(ctx)
	res, err := a.walletClient.CreateTradingTransaction(ctx, request)
	if err != nil {
		logger.Warn("CreateWalletTransaction failed to create trading transaction.", "Error", err)
		return nil, err
	}
	logger.Info("CreateWalletTransaction successfully created trading transaction", "TransactionId", res.TransactionId)
	return res, err
}

func (a *TradingActivities) PlaceTradingOrder(ctx context.Context, request *xc.PlaceOrderRequest) (*xc.PlaceOrderResponse, error) {
	logger := activity.GetLogger(ctx)
	res, err := a.xcClient.PlaceOrder(ctx, request)
	if err != nil {
		logger.Warn("PlaceTradingOrder failed to place order.", "Error", err)
		return nil, err
	}
	logger.Info("PlaceTradingOrder successfully placed order", "OrderId", res.OrderId)
	return res, err
}

func (a *TradingActivities) ConfirmWalletTransaction(ctx context.Context, request *wallet.ConfirmTransactionRequest) error {
	logger := activity.GetLogger(ctx)
	_, err := a.walletClient.ConfirmTransaction(ctx, request)
	if err != nil {
		logger.Error("failed to confirm transaction", "Error", err)
		return err
	}
	return nil
}

func (a *TradingActivities) RejectWalletTransaction(ctx context.Context, param TradingRequest) error {
	log.Println("RejectWalletTransactionActivity")
	return nil
}
