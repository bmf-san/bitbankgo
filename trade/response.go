package trade

import "github.com/bmf-san/bitbankgo/types"

// GetTradeHistoryResponse is a response for getting a order.
type GetTradeHistoryResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code   int `json:"code"`
		Trades []struct {
			TradeID        int            `json:"trade_id"`
			Pair           types.TypePair `json:"Pair"`
			OrderID        int            `json:"order_id"`
			Side           string         `json:"side"`
			Type           string         `json:"type"`
			Amount         string         `json:"amount"`
			Price          string         `json:"price"`
			MakerTaker     string         `json:"maker_taker"`
			FeeAmountBase  string         `json:"fee_amount_base"`
			FeeAmountQuote string         `json:"fe10:12 PMe_amount_quote"`
			ExecutedAt     int            `json:"executed_at"`
		} `json:"trades"`
	} `json:"data"`
}
