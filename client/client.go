package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Client is a api client.
type Client struct {
	HTTPClient *http.Client
	APIKey     string
	APISecret  string
}

// New creates a new api client.
func New(apiKey string, apiSecret string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

// NewRequest returns a new Request given a method and url.
func (c *Client) NewRequest(method string, header http.Header, body []byte, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = header
	}

	return req, nil
}

// CertificationHeader creates a header.
func (c *Client) CertificationHeader(path string) (http.Header, error) {
	accessNonce := int(time.Now().UnixNano() / int64(time.Millisecond))
	mac := hmac.New(sha256.New, []byte(c.APISecret))
	_, err := mac.Write([]byte(fmt.Sprintf("%d%s%s", accessNonce, path, "")))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	accessSignature := hex.EncodeToString(mac.Sum(nil))
	header := make(http.Header)
	header["ACCESS-KEY"] = []string{c.APIKey}
	header["ACCESS-NONCE"] = []string{strconv.Itoa(accessNonce)}
	header["ACCESS-SIGNATURE"] = []string{accessSignature}
	return header, nil
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
