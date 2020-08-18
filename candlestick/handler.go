package candlestick

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/bmf-san/bitbankgo/types"
	"github.com/pkg/errors"
)

// Get gets candlestick infomation.
// The formats that time supports are YYYYMMDDD and YYYY.
func (c *CandleStick) Get(pair types.TypePair, candle types.TypeCandle, time time.Time) (*CandleStickResponse, error) {
	req, err := c.Client.NewRequest(http.MethodGet, nil, nil, api.URL(api.CandlestickEndpoint, "", pair, candle, time.Format("20060102")))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cr := &CandleStickResponse{}
	if err = json.Unmarshal(body, cr); err != nil {
		return nil, errors.WithStack(err)
	}

	if cr.Success != 1 {
		return nil, errors.New(converter.ErrorText(cr.Data.Code))
	}

	return cr, nil
}
