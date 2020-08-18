package ticker

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/bmf-san/bitbankgo/types"
	"github.com/pkg/errors"
)

// Get gets ticker infomation.
func (t *Ticker) Get(pair types.TypePair) (*TickerResponse, error) {
	req, err := t.Client.NewRequest(http.MethodGet, nil, nil, api.URL(api.TickerEndpoint, "", pair))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := t.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tr := &TickerResponse{}
	if err = json.Unmarshal(body, &tr); err != nil {
		return nil, errors.WithStack(err)
	}

	if tr.Success != 1 {
		return nil, errors.New(converter.ErrorText(tr.Data.Code))
	}

	return tr, nil
}
