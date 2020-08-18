package withdrawal

// GetWithdrawalAccountResponse is a response for getting a withdrawal account.
type GetWithdrawalAccountResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code     int `json:"code"`
		Accounts []struct {
			UUID    string `json:"uuid"`
			Label   string `json:"label"`
			Address string `json:"address"`
		} `json:"accounts"`
	} `json:"data"`
}
