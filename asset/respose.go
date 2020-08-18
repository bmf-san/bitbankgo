package asset

// AssetResponse is a response for Asset.
type AssetResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code   int `json:"code"`
		Assets []struct {
			Asset           string      `json:"asset"`
			FreeAmount      string      `json:"free_amount"`
			AmountPrecision int         `json:"amount_precision"`
			OnhandAmount    string      `json:"onhand_amount"`
			LockedAmound    string      `json:"locked_amount"`
			WithdrawalFee   interface{} `json:"withdrawal_fee"`
			StopDeposit     bool        `json:"stop_deposit"`
			StopWithdrawal  bool        `json:"stop_withdrawal"`
		} `json:"assets"`
	} `json:"data"`
	Asset string
}
