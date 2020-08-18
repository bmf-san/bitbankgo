package asset

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/bitbankgo/api"
	"github.com/bmf-san/bitbankgo/converter"
	"github.com/pkg/errors"
)

// Get gets a asset infomation
func (a *Asset) Get() (*AssetResponse, error) {
	header, err := a.Client.CertificationHeader(api.AssetPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := a.Client.NewRequest(http.MethodGet, header, nil, api.URL(api.AssetEndpoint, ""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := a.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ar := &AssetResponse{}
	if err = json.Unmarshal(body, ar); err != nil {
		return nil, errors.WithStack(err)
	}

	if ar.Success != 1 {
		return nil, errors.New(converter.ErrorText(ar.Data.Code))
	}

	return ar, nil
}
