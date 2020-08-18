package status

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/pkg/errors"
)

// GetStatus gets status.
func (s *Status) GetStatus() (*GetStatusResponse, error) {
	req, err := s.Client.NewRequest(http.MethodGet, nil, nil, api.URL(api.StatusEndpoint, ""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := s.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gr := &GetStatusResponse{}
	if err = json.Unmarshal(body, gr); err != nil {
		return nil, errors.WithStack(err)
	}

	if gr.Success != 1 {
		return nil, errors.New(converter.ErrorText(gr.Data.Code))
	}

	return gr, nil
}
