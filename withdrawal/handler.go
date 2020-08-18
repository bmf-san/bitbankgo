package withdrawal

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/client"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/pkg/errors"
)

// Withdrawal is a withdrawal.
type Withdrawal struct {
	Client *client.Client
}

// GetWithdrawalAccount gets Withdrawal history.
func (w *Withdrawal) GetWithdrawalAccount(params GetWithdrawalAccountParams) (*GetWithdrawalAccountResponse, error) {
	url := &url.URL{}
	query := url.Query()
	query.Add("asset", string(params.Asset))

	header, err := w.Client.CertificationHeader(api.WithdrawalAccountPath + "?" + query.Encode())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := w.Client.NewRequest(http.MethodGet, header, nil, api.URL(api.WithdrawalAccountEndpoint+"?", query.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := w.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	wr := &GetWithdrawalAccountResponse{}
	if err = json.Unmarshal(body, wr); err != nil {
		return nil, errors.WithStack(err)
	}

	if wr.Success != 1 {
		return nil, errors.New(converter.ErrorText(wr.Data.Code))
	}

	return wr, nil
}
