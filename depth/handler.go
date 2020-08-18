package depth

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/bmf-san/bitbankgo/types"
	"github.com/pkg/errors"
)

// Get gets depth infomation.
func (d *Depth) Get(pair types.TypePair) (*DepthResponse, error) {
	req, err := d.Client.NewRequest(http.MethodGet, nil, nil, api.URL(api.DepthEndpoint, "", pair))
	if err != nil {
		return nil, err
	}

	body, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}

	dr := &DepthResponse{}
	if err = json.Unmarshal(body, &dr); err != nil {
		return nil, err
	}

	if dr.Success != 1 {
		return nil, errors.New(converter.ErrorText(dr.Data.Code))
	}

	return dr, nil
}
