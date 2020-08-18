package order

import "github.com/bmf-san/bitbankgo/types"

// GetOrderParams is parameters for getting a order.
type GetOrderParams struct {
	Pair    types.TypePair `json:"pair"`
	OrderID int            `json:"order_id"`
}

// PostOrderParams is parameters for creating a order.
type PostOrderParams struct {
	Pair   types.TypePair `json:"pair"`
	Amount string         `json:"amount"`
	Price  string         `json:"price"`
	Side   string         `json:"side"`
	Type   string         `json:"type"`
}

// PostCancelOrderParams is parameters for canceling a order.
type PostCancelOrderParams struct {
	Pair    types.TypePair `json:"pair"`
	OrderID int            `json:"order_id"`
}

// PostCancelOrdersParams is parameters for canceling orders.
type PostCancelOrdersParams struct {
	Pair     types.TypePair `json:"pair"`
	OrderIDs []int          `json:"order_ids"`
}

// GetOrdersInfoParams is parameters for getting multiple orders.
type GetOrdersInfoParams struct {
	Pair     types.TypePair `json:"pair"`
	OrderIDs []int          `json:"order_ids"`
}

// GetActiveOrdersParams is parameters for getting multiple orders.
type GetActiveOrdersParams struct {
	Pair   types.TypePair `json:"pair"`
	Count  int            `json:"count"`
	FromID int            `json:"from_id"`
	EndID  int            `json:"end_id"`
	Since  int            `json:"since"`
	End    int            `json:"end"`
}
