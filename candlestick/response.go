package candlestick

// CandleStickResponse is a response for CandleStickResponse.
type CandleStickResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code        int `json:"code"`
		CandleStick []struct {
			Type  string          `json:"type"`
			Ohlcv [][]interface{} `json:"ohlcv"`
		} `json:"candlestick"`
	} `json:"data"`
}
