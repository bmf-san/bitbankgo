package api

import (
	"fmt"
)

const (
	// PublicAPIEndpoint is a base endpoint for public api.
	PublicAPIEndpoint = "https://public.bitbank.cc"
	// PrivateAPIEndpoint is a base endpoint for private api.
	PrivateAPIEndpoint = "https://api.bitbank.cc"

	// TickerEndpoint is a /{pair}/ticker.
	TickerEndpoint = PublicAPIEndpoint + TickerPath
	// DepthEndpoint is a /{pair}/depth.
	DepthEndpoint = PublicAPIEndpoint + DepthPath
	// TransactionEndpoint is a /{pair}/transactions/{YYYYMMDD}.
	TransactionEndpoint = PublicAPIEndpoint + TransactionPath
	// CandlestickEndpoint is a /{pair}/candlestick/{candle-type}/{YYYYMMDD|YYYY}.
	CandlestickEndpoint = PublicAPIEndpoint + CandleStickPath
	// AssetEndpoint is a /v1/user/assets.
	AssetEndpoint = PrivateAPIEndpoint + AssetPath
	// OrderEndpoint is a /v1/user/spot/order.
	OrderEndpoint = PrivateAPIEndpoint + OrderPath
	// CancelOrderEndpoint is a /v1/user/spot/cancel_order.
	CancelOrderEndpoint = PrivateAPIEndpoint + CancelOrderPath
	// CancelOrdersEndpoint is a /v1/user/spot/cancel_orders.
	CancelOrdersEndpoint = PrivateAPIEndpoint + CancelOrdersPath
	// OrdersInfoEndpoint is a /v1/user/spot/orders_info.
	OrdersInfoEndpoint = PrivateAPIEndpoint + OrdersInfoPath
	// ActiveOrdersEndpoint is a /v1/user/spot/active_orders.
	ActiveOrdersEndpoint = PrivateAPIEndpoint + ActiveOrdersPath
	// TradeHistoryEndpoint is a /v1/user/spot/trade_history.
	TradeHistoryEndpoint = PrivateAPIEndpoint + TradeHistoryPath
	// WithdrawalAccountEndpoint is a /v1/user/withdrawal_account
	WithdrawalAccountEndpoint = PrivateAPIEndpoint + WithdrawalAccountPath
	// StatusEndpoint is a /v1/spot/status
	StatusEndpoint = PrivateAPIEndpoint + StatusPath
	// SettingEndpoint is a /v1/spot/pairs
	SettingEndpoint = PrivateAPIEndpoint + SettingPath

	// TickerPath is a path for ticker endpoint.
	TickerPath = "/%s/ticker"
	// DepthPath is a path for depth endpoint.
	DepthPath = "/%s/depth"
	// TransactionPath is a path for transactions endpoint.
	TransactionPath = "/%s/transactions/%s"
	// CandleStickPath is a path for candlestick endpoint.
	CandleStickPath = "/%s/candlestick/%s/%s"
	// AssetPath is a path for asset endpoint.
	AssetPath = "/v1/user/assets"
	// OrderPath is a path for order endpoint.
	OrderPath = "/v1/user/spot/order"
	// CancelOrderPath is a path for cancel order endpoint.
	CancelOrderPath = "/v1/user/spot/cancel_orders"
	// CancelOrdersPath is a path for cancel orders endpoint.
	CancelOrdersPath = "/v1/user/spot/cancel_orders"
	// OrdersInfoPath is a path for orders info endpoint.
	OrdersInfoPath = "/v1/user/spot/orders_info"
	// ActiveOrdersPath is a path for active orders endpoint.
	ActiveOrdersPath = "/v1/user/spot/active_orders"
	// TradeHistoryPath is a path for active orders endpoint.
	TradeHistoryPath = "/v1/user/spot/trade_history"
	// WithdrawalAccountPath is a path for withdrawal account.
	WithdrawalAccountPath = "/v1/user/withdrawal_account"
	// StatusPath is a path for status.
	StatusPath = "/v1/spot/status"
	// SettingPath is a path for setting.
	SettingPath = "/v1/spot/pairs"
)

// URL returns a url by endpoint and arguments.
func URL(endpoint string, query string, args ...interface{}) string {
	if query != "" {
		endpoint = endpoint + query
	}

	return fmt.Sprintf(endpoint, args...)
}
