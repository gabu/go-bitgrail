// Package bitgrail is the unofficial client to access to BitGrail API
package bitgrail

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	// APIEndpoint is the BitGrail API endpoint.
	APIEndpoint = "https://bitgrail.com/api/v1/"
)

// Client manages all the communication with the BitGrail API.
type Client struct {
	// Base URL for Public API requests.
	BaseURL *url.URL

	// Auth data
	APIKey    string
	APISecret string
}

// NewClient creates new BitGrail API client.
func NewClient() *Client {
	baseURL, _ := url.Parse(APIEndpoint)
	return &Client{BaseURL: baseURL}
}

// newRequest create new API request. Relative url can be provided in refURL.
func (c *Client) newRequest(ctx context.Context, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}
	if params != nil {
		rel.RawQuery = params.Encode()
	}

	var req *http.Request
	u := c.BaseURL.ResolveReference(rel)
	req, err = http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

// newAuthenticatedRequest creates new http request for authenticated routes.
func (c *Client) newAuthenticatedRequest(ctx context.Context, method string, values url.Values) (*http.Request, error) {
	if values == nil {
		values = url.Values{}
	}

	values.Set("nonce", strconv.FormatInt(time.Now().UnixNano(), 10))
	body := values.Encode()

	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, method)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Sign request
	h := hmac.New(sha512.New, []byte(c.APISecret))
	h.Write([]byte(body))
	sign := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("KEY", c.APIKey)
	req.Header.Add("SIGNATURE", sign)

	req = req.WithContext(ctx)

	return req, nil
}

// Auth sets api key and secret for usage is requests that requires authentication.
func (c *Client) Auth(key string, secret string) *Client {
	c.APIKey = key
	c.APISecret = secret

	return c
}

var httpDo = func(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

// Do executes API request created by NewRequest method or custom *http.Request.
func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := httpDo(req)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = checkResponse(response)
	if err != nil {
		// Return response in case caller need to debug it.
		return response, errors.Wrap(err, "")
	}

	if v == nil {
		return response, nil
	}

	apiRes := APIResponse{}
	err = json.Unmarshal(response.Body, &apiRes)
	if err != nil {
		return response, errors.Wrap(err, "Faild to unmarshal response body to json")
	}
	if apiRes.Success != 1 {
		ev := map[string]string{}
		err = json.Unmarshal(apiRes.Response, &ev)
		if err != nil {
			return response, errors.Wrap(err, "Faild to unmarshal response.error to json")
		}
		return response, &ErrorResponse{Response: response, Message: "error: " + ev["error"]}
	}

	err = json.Unmarshal(apiRes.Response, v)
	if err != nil {
		return response, errors.Wrap(err, "")
	}

	return response, nil
}

// APIResponse is the API's response base json format.
type APIResponse struct {
	Success  int // 1 is successed
	Response json.RawMessage
}

// Response is wrapper for standard http.Response and provides
// more methods.
type Response struct {
	Response *http.Response
	Body     []byte
}

// newResponse creates new wrapper.
func newResponse(r *http.Response) *Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body = []byte(`Error reading body:` + err.Error())
	}

	return &Response{r, body}
}

// String converts response body to string.
// An empty string will be returned if error.
func (r *Response) String() string {
	return string(r.Body)
}

// ErrorResponse is the custom error type that is returned if the API returns an
// error.
type ErrorResponse struct {
	Response *Response
	Message  string `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Response.Request.Method,
		r.Response.Response.Request.URL,
		r.Response.Response.StatusCode,
		r.Message,
	)
}

// checkResponse checks response status code and response
// for errors.
func checkResponse(r *Response) error {
	if c := r.Response.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	// Try to decode error message
	errorResponse := &ErrorResponse{Response: r}
	err := json.Unmarshal(r.Body, errorResponse)
	if err != nil {
		errorResponse.Message = "Error decoding response error message. " +
			"Please see response body for more information."
	}

	return errorResponse
}
