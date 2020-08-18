package status

import "github.com/bmf-san/bitbankgo/types"

// GetStatusResponse is a response for getting status.
type GetStatusResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code     int `json:"code"`
		Statuses []struct {
			Pair      types.TypePair   `json:"Pair"`
			Status    types.TypeStatus `json:"status"`
			MinAmount string           `json:"min_amount"`
		}
	}
}
