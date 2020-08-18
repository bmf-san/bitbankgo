package order

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/pkg/errors"
)

// GetOrder gets an order infomation.
func (o *Order) GetOrder(params GetOrderParams) (*GetOrderResponse, error) {
	url := &url.URL{}
	query := url.Query()
	query.Add("pair", string(params.Pair))
	query.Add("order_id", strconv.Itoa(params.OrderID))

	header, err := o.Client.CertificationHeader(api.OrderPath + "?" + query.Encode())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := o.Client.NewRequest(http.MethodGet, header, nil, api.URL(api.OrderEndpoint+"?", query.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := o.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	or := &GetOrderResponse{}
	if err = json.Unmarshal(body, or); err != nil {
		return nil, errors.WithStack(err)
	}

	if or.Success != 1 {
		return nil, errors.New(converter.ErrorText(or.Data.Code))
	}

	return or, nil
}

// PostOrder creates a new order.
func (o *Order) PostOrder(params PostOrderParams) (*PostOrderResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	header, err := o.Client.CertificationHeader(string(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := o.Client.NewRequest(http.MethodPost, header, body, api.URL(api.OrderEndpoint, ""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err = o.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	or := &PostOrderResponse{}
	if err = json.Unmarshal(body, or); err != nil {
		return nil, errors.WithStack(err)
	}

	if or.Success != 1 {
		return nil, errors.New(converter.ErrorText(or.Data.Code))
	}

	return or, nil
}

// PostCancelOrder cancels order.
func (o *Order) PostCancelOrder(params PostCancelOrderParams) (*PostCancelOrderResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	header, err := o.Client.CertificationHeader(string(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := o.Client.NewRequest(http.MethodPost, header, body, api.URL(api.CancelOrderEndpoint, ""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err = o.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	or := &PostCancelOrderResponse{}
	if err = json.Unmarshal(body, or); err != nil {
		return nil, errors.WithStack(err)
	}

	if or.Success != 1 {
		return nil, errors.New(converter.ErrorText(or.Data.Code))
	}

	return or, nil
}

// PostCancelOrders cancels orders.
func (o *Order) PostCancelOrders(params PostCancelOrdersParams) (*PostCancelOrdersResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	header, err := o.Client.CertificationHeader(string(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := o.Client.NewRequest(http.MethodPost, header, body, api.URL(api.CancelOrdersEndpoint, ""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err = o.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	or := &PostCancelOrdersResponse{}
	if err = json.Unmarshal(body, or); err != nil {
		return nil, errors.WithStack(err)
	}

	if or.Success != 1 {
		return nil, errors.New(converter.ErrorText(or.Data.Code))
	}

	return or, nil
}

// GetOrdersInfo gets multiple orders.
func (o *Order) GetOrdersInfo(params GetOrdersInfoParams) (*PostOrdersInfoResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	header, err := o.Client.CertificationHeader(string(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := o.Client.NewRequest(http.MethodPost, header, body, api.URL(api.OrdersInfoEndpoint, ""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err = o.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	or := &PostOrdersInfoResponse{}
	if err = json.Unmarshal(body, or); err != nil {
		return nil, errors.WithStack(err)
	}

	if or.Success != 1 {
		return nil, errors.New(converter.ErrorText(or.Data.Code))
	}

	return or, nil
}

// GetActiveOrders gets multiple orders.
func (o *Order) GetActiveOrders(params GetActiveOrdersParams) (*GetActiveOrdersResponse, error) {
	url := &url.URL{}
	query := url.Query()
	query.Add("pair", string(params.Pair))
	if params.Count != 0 {
		query.Add("count", strconv.Itoa(params.Count))
	}
	if params.FromID != 0 {
		query.Add("from_id", strconv.Itoa(params.FromID))
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

	header, err := o.Client.CertificationHeader(api.ActiveOrdersPath + "?" + query.Encode())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := o.Client.NewRequest(http.MethodGet, header, nil, api.URL(api.ActiveOrdersEndpoint+"?", query.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := o.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	or := &GetActiveOrdersResponse{}
	if err = json.Unmarshal(body, or); err != nil {
		return nil, errors.WithStack(err)
	}

	if or.Success != 1 {
		return nil, errors.New(converter.ErrorText(or.Data.Code))
	}

	return or, nil
}
