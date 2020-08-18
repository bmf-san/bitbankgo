package setting

// GetSettingResponse is a response for getting setting.
type GetSettingResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code  int `json:"code"`
		Pairs []struct {
			Name                string `json:"name"`
			BaseAsset           string `json:"base_asset"`
			MakerFeeRateBase    string `json:"maker_fee_rate_base"`
			TakerFeeRateBase    string `json:"taker_fee_rate_base"`
			MakerFeeRateQuote   string `json:"maker_fee_rate_quote"`
			TakerFeeRateQuote   string `json:"taker_fee_rate_quote"`
			UnitAmount          string `json:"unit_amount"`
			LimitMaxAmount      string `json:"limit_max_amount"`
			MarketMaxAmount     string `json:"market_max_amount"`
			MarketAllowanceRate string `json:"market_allowance_rate"`
			PriceDigits         int    `json:"price_digits"`
			AmountDigits        int    `json:"amount_digits"`
			IsEnabled           bool   `json:"is_enabled"`
			StopOrder           bool   `json:"stop_order"`
			StopOrderAndCancel  bool   `json:"stop_order_and_cancel"`
		} `json:"pairs"`
	} `json:"data"`
}
