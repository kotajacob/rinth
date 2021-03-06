// rinth is a library for reading mod data from Modrinth.
package rinth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BaseURL is the default base URL for the Modrinth JSON API.
const BaseURL = "https://api.modrinth.com/api/v1/"

// Client used for metservice-go.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient constructs a client using http.DefaultClient and the default
// base URL. The returned client is ready for use.
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    BaseURL,
	}
}

// StatusError is returned when a bad responce code is received from the API.
type StatusError struct {
	Code int
}

var _ error = StatusError{}

func (e StatusError) Error() string {
	return fmt.Sprintf("bad responce status code: %d", e.Code)
}

// Do sends an API request and returns the API responce. The API responce is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occured. If v is nil, and no error happens, the
// responce is returned as is.
func (c *Client) Do(ctx context.Context, path string, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("rinth failed to build request: %v", err)
	}
	req = req.WithContext(ctx)

	rsp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("rinth failed to do request: %v", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, StatusError{Code: rsp.StatusCode}
	}

	switch v := v.(type) {
	case nil:
	default:
		decErr := json.NewDecoder(rsp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty responce body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return rsp, err
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
