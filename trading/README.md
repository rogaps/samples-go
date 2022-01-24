# Trading

## Commands

### Trading

#### Create Trading

```
grpcurl -plaintext \
  -proto trading/proto/trading/trading.proto \
  -d '{ "base_asset_type": "btc", "quote_asset_type": "idr", "side": "BUY", "amount": "10000", "amount_asset_type": "idr" }' \
  localhost:50051 trading.TradingService/CreateTrading
```

#### Commit Trading

```
grpcurl -plaintext \
  -proto trading/proto/trading/trading.proto \
  -d '{ "transaction_id": "f1de5b7e-3777-446a-8995-4e0217e82a7f" }' \
  localhost:50051 trading.TradingService/CommitTrading
```

### Wallet

```

```

### XC

#### Get Exchange Amount

```
grpcurl -plaintext \
  -proto trading/proto/xc/xc.proto \
  -d '{ "base_asset_type": "btc", "quote_asset_type": "idr", "side": "BUY", "amount": "10000", "amount_asset_type": "idr" }' \
  localhost:50053 xc.XcService/GetExchangeAmount
```

#### Place Order

```
grpcurl -plaintext \
  -proto trading/proto/xc/xc.proto \
  -d '{"base_asset_type": "btc", "quote_asset_type": "idr", "side": "BUY", "base_amount": "1000", "quote_amount": "1000"}'
  localhost:50053 xc.XcService/PlaceOrder
```
