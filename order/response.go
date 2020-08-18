package order

import "github.com/bmf-san/bitbankgo/types"

// GetOrderResponse is a response for getting a order.
type GetOrderResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code            int            `json:"code"`
		OrderID         int            `json:"order_id"`
		Pair            types.TypePair `json:"pair"`
		Side            string         `json:"side"`
		Type            string         `json:"type"`
		StartAmount     string         `json:"start_amount"`
		RemainingAmount string         `json:"remaining_amount"`
		ExecutedAmount  string         `json:"executed_amount"`
		Price           string         `json:"price"`
		AveragePrice    string         `json:"average_price"`
		OrderedAt       int            `json:"ordered_at"`
		Status          string         `json:"status"`
	} `json:"data"`
}

// PostOrderResponse is a response for creating a order.
type PostOrderResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code            int            `json:"code"`
		OrderID         int            `json:"order_id"`
		Pair            types.TypePair `json:"pair"`
		Side            string         `json:"side"`
		Type            string         `json:"type"`
		StartAmount     string         `json:"start_amount"`
		RemainingAmount string         `json:"remaining_amount"`
		ExecutedAmount  string         `json:"executed_amount"`
		Price           string         `json:"price"`
		AveragePrice    string         `json:"average_price"`
		OrderedAt       int            `json:"ordered_at"`
		Status          string         `json:"status"`
	} `json:"data"`
}

// PostCancelOrderResponse is a response for canceling a order.
type PostCancelOrderResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code            int            `json:"code"`
		OrderID         int            `json:"order_id"`
		Pair            types.TypePair `json:"pair"`
		Side            string         `json:"side"`
		Type            string         `json:"type"`
		StartAmount     string         `json:"start_amount"`
		RemainingAmount string         `json:"remaining_amount"`
		ExecutedAmount  string         `json:"executed_amount"`
		Price           string         `json:"price"`
		AveragePrice    string         `json:"average_price"`
		OrderedAt       int            `json:"ordered_at"`
		CanceledAt      int            `json:"canceled_at"`
		Status          string         `json:"status"`
	} `json:"data"`
}

// PostCancelOrdersResponse is a response for canceling orders.
type PostCancelOrdersResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code   int `json:"code"`
		Orders []struct {
			OrderID         int            `json:"order_id"`
			Pair            types.TypePair `json:"pair"`
			Side            string         `json:"side"`
			Type            string         `json:"type"`
			StartAmount     string         `json:"start_amount"`
			RemainingAmount string         `json:"remaining_amount"`
			ExecutedAmount  string         `json:"executed_amount"`
			Price           string         `json:"price"`
			AveragePrice    string         `json:"average_price"`
			OrderedAt       int            `json:"ordered_at"`
			CanceledAt      int            `json:"canceled_at"`
			Status          string         `json:"status"`
		} `json:"orders"`
	}
}

// PostOrdersInfoResponse is a response for getting multiple orders.
type PostOrdersInfoResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code   int `json:"code"`
		Orders []struct {
			OrderID         int            `json:"order_id"`
			Pair            types.TypePair `json:"pair"`
			Side            string         `json:"side"`
			Type            string         `json:"type"`
			StartAmount     string         `json:"start_amount"`
			RemainingAmount string         `json:"remaining_amount"`
			ExecutedAmount  string         `json:"executed_amount"`
			Price           string         `json:"price"`
			AveragePrice    string         `json:"average_price"`
			OrderedAt       int            `json:"ordered_at"`
			CanceledAt      int            `json:"canceled_at"`
			Status          string         `json:"status"`
		} `json:"orders"`
	} `json:"data"`
}

// GetActiveOrdersResponse is a response for getting multiple orders.
type GetActiveOrdersResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code   int `json:"code"`
		Orders []struct {
			Pair            types.TypePair `json:"pair"`
			Side            string         `json:"side"`
			Type            string         `json:"type"`
			StartAmount     string         `json:"start_amount"`
			RemainingAmount string         `json:"remaining_amount"`
			ExecutedAmount  string         `json:"executed_amount"`
			Price           string         `json:"price"`
			AveragePrice    string         `json:"average_price"`
			OrderedAt       int            `json:"ordered_at"`
			Status          string         `json:"status"`
		} `json:"orders"`
	} `json:"data"`
}
