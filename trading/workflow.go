package trading

import (
	"time"

	"github.com/temporalio/samples-go/trading/proto/trading"
	"github.com/temporalio/samples-go/trading/proto/wallet"
	"github.com/temporalio/samples-go/trading/proto/xc"
	"go.temporal.io/sdk/workflow"
)

func TradingWorkflow(ctx workflow.Context, request *trading.CreateTradingResponse) (err error) {
	logger := workflow.GetLogger(ctx)

	var tradingActivities *TradingActivities

	// normal activity options
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})

	err = workflow.ExecuteActivity(ctx, tradingActivities.CreateTrading).Get(ctx, nil)
	if err != nil {
		return NewNonRetryableError(err.Error())
	}

	// commit trading timeout
	waitCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 60 * time.Second,
	})
	err = workflow.ExecuteActivity(waitCtx, tradingActivities.WaitForCommitTrading).Get(waitCtx, nil)
	if err != nil {
		logger.Error("failed to wait for commit trading", "Error", err)
		return err
	}

	// wallet create transaction activity
	walletReq := &wallet.CreateTradingRequest{
		BaseAssetType:  request.BaseAssetType,
		QuoteAssetType: request.QuoteAssetType,
		Side:           wallet.TradingSide(request.Side),
		BaseAmount:     request.Amount,
		QuoteAmount:    request.Total,
	}
	var walletRes *wallet.CreateTradingResponse
	err = workflow.ExecuteActivity(ctx, tradingActivities.CreateWalletTransaction, walletReq).Get(ctx, &walletRes)
	if err != nil {
		logger.Error("failed to create wallet transaction", "Error", err)
		return err
	}

	// xc activity
	xcReq := &xc.PlaceOrderRequest{
		BaseAssetType:  request.BaseAssetType,
		QuoteAssetType: request.QuoteAssetType,
		Side:           xc.TradingSide(request.Side),
		BaseAmount:     request.Amount,
		QuoteAmount:    request.Total,
	}
	var xcRes *xc.PlaceOrderResponse
	err = workflow.ExecuteActivity(ctx, tradingActivities.PlaceTradingOrder, xcReq).Get(ctx, &xcRes)
	if err != nil {
		err = workflow.ExecuteActivity(ctx, tradingActivities.RejectWalletTransaction).Get(ctx, nil)
		logger.Error("failed to place order", "Error", err)
		return err
	}

	// wallet confirm transaction
	confirmReq := &wallet.ConfirmTransactionRequest{
		TransactionId: walletRes.TransactionId,
	}
	err = workflow.ExecuteActivity(ctx, tradingActivities.ConfirmWalletTransaction, confirmReq).Get(ctx, nil)
	if err != nil {
		logger.Error("failed to confirm transaction", "Error", err)
		return err
	}

	return nil
}
