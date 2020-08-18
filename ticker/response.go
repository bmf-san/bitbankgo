package ticker

// TickerResponse is a response for Ticker.
type TickerResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code      int    `json:"code"`
		Sell      string `json:"sell"`
		Buy       string `json:"buy"`
		High      string `json:"high"`
		Low       string `json:"low"`
		Last      string `json:"last"`
		Vol       string `json:"vol"`
		Timestamp int    `json:"timestamp"`
	} `json:"data"`
}
