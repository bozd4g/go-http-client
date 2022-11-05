package gohttpclient

import "time"

type (
	Option       func(c *Client)
	ClientOption Option
)

func WithDefaultHeaders() ClientOption {
	return func(c *Client) {
		if c.headers == nil {
			c.headers = make(map[string]Header)
		}

		c.headers["Content-Type"] = Header{Value: "application/json", IsDefault: true}
		c.headers["Accept"] = Header{Value: "application/json", IsDefault: true}
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
		c.httpClient.Timeout = timeout
	}
}

func WithHeader(key, value string) Option {
	return func(c *Client) {
		if c.headers == nil {
			c.headers = make(map[string]Header)
		}

		c.headers[key] = Header{Value: value, IsDefault: false}
	}
}

func WithQuery(key, value string) Option {
	return func(c *Client) {
		if c.query == nil {
			c.query = make(map[string]string)
		}

		c.query[key] = value
	}
}

func WithBody(body []byte) Option {
	return func(c *Client) {
		c.body = body
	}
}
