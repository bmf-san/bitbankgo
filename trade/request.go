package trade

import "github.com/bmf-san/bitbankgo/types"

// GetTradeHistoryParams is parameters for getting trade history.
type GetTradeHistoryParams struct {
	Pair    types.TypePair `json:"pair"`
	Count   int            `json:"count"`
	OrderID int            `json:"order_id"`
	EndID   int            `json:"end_id"`
	Since   int            `json:"since"`
	End     int            `json:"end"`
	Order   string         `json:"order"`
}
