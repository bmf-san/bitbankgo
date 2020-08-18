package trade

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/pkg/errors"
)

// GetTradeHistory gets trade history.
func (t *Trade) GetTradeHistory(params GetTradeHistoryParams) (*GetTradeHistoryResponse, error) {
	url := &url.URL{}
	query := url.Query()
	query.Add("pair", string(params.Pair))

	if params.OrderID != 0 {
		query.Add("order_id", strconv.Itoa(params.OrderID))
	}
	if params.EndID != 0 {
		query.Add("end_id", strconv.Itoa(params.EndID))
	}
	if params.Since != 0 {
		query.Add("since", strconv.Itoa(params.Since))
	}
	if params.End != 0 {
		query.Add("end", strconv.Itoa(params.End))
	}
	if params.Order != "" {
		query.Add("order", string(params.Order))
	}

	header, err := t.Client.CertificationHeader(api.TradeHistoryPath + "?" + query.Encode())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := t.Client.NewRequest(http.MethodGet, header, nil, api.URL(api.TradeHistoryEndpoint+"?", query.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := t.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tr := &GetTradeHistoryResponse{}
	if err = json.Unmarshal(body, tr); err != nil {
		return nil, errors.WithStack(err)
	}

	if tr.Success != 1 {
		return nil, errors.New(converter.ErrorText(tr.Data.Code))
	}

	return tr, nil
}
