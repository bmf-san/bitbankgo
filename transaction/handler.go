package transaction

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/bmf-san/bitbankgo/types"
	"github.com/pkg/errors"
)

// Get gets transaction infomation.
func (t *Transaction) Get(pair types.TypePair, time time.Time) (*TransactionResponse, error) {
	req, err := t.Client.NewRequest(http.MethodGet, nil, nil, api.URL(api.TransactionEndpoint, "", pair, time.Format("20060102")))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := t.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tr := &TransactionResponse{}
	if err = json.Unmarshal(body, tr); err != nil {
		return nil, errors.WithStack(err)
	}

	if tr.Success != 1 {
		return nil, errors.New(converter.ErrorText(tr.Data.Code))
	}

	return tr, nil
}
