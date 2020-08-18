package withdrawal

import "github.com/bmf-san/bitbankgo/types"

// GetWithdrawalAccountParams is a parameters for getting a withdrawal account.
type GetWithdrawalAccountParams struct {
	Asset types.TypeAsset `json:"asset"`
}
