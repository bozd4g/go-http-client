package gohttpclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type (
	// Client is a struct who has BaseUrl property
	Client struct {
		baseUrl    string
		httpClient *http.Client

		headers map[string]Header
		query   map[string]string
		body    []byte
		timeout time.Duration
	}

	// Clienter is a interface who calls the methods
	Clienter interface{}

	Header struct {
		Value     string
		IsDefault bool
	}
)

const (
	DEFAULT_TIMEOUT = 10 * time.Second
)

// New func returns a Client struct
func New(baseUrl string, opts ...ClientOption) *Client {
	httpClient := &http.Client{Timeout: DEFAULT_TIMEOUT}
	client := &Client{httpClient: httpClient, baseUrl: baseUrl, timeout: DEFAULT_TIMEOUT}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Get func returns a request
func (c *Client) Get(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Post func returns a request
func (c *Client) Post(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)

}

// Put func returns a request
func (c *Client) Put(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseUrl+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Patch func returns a request
func (c *Client) Patch(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.baseUrl+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Delete func returns a request
func (c *Client) Delete(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseUrl+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

func (c *Client) Connect(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodConnect, c.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

func (c *Client) Options(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, c.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

func (c *Client) Trace(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, http.MethodTrace, c.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// PrepareRequest func returns a request
func (c *Client) PrepareRequest(ctx context.Context, method, endpoint string, opts ...Option) (*http.Request, error) {
	clear := c.initOpts(opts...)
	defer clear()

	req, err := http.NewRequestWithContext(ctx, method, c.baseUrl+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}

	return c.prepareReq(req), nil
}

func (c *Client) initOpts(opts ...Option) func() {
	for _, opt := range opts {
		opt(c)
	}

	return func() {
		for key, header := range c.headers {
			if !header.IsDefault {
				delete(c.headers, key)
			}
		}

		c.query = make(map[string]string)
		c.body = nil
	}
}

func (c *Client) prepareReq(req *http.Request) *http.Request {
	// set headers
	for key, header := range c.headers {
		req.Header.Set(key, header.Value)
	}

	// set query
	q := req.URL.Query()
	for key, value := range c.query {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
	return req
}

func (c *Client) sendReq(ctx context.Context, req *http.Request) (*Response, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res, err := c.httpClient.Do(req.WithContext(reqCtx))
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return &Response{res, body}, nil
}
