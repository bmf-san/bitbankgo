package depth

// DepthResponse is a response for Depth.
type DepthResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code int        `json:"code"`
		Asks [][]string `json:"asks"`
		Bids [][]string `json:"bids"`
	} `json:"data"`
}
