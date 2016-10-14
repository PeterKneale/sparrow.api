package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AliveHealthPath computes a request path to the alive action of health.
func AliveHealthPath() string {
	return fmt.Sprintf("/alive")
}

// Perform health check.
func (c *Client) AliveHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAliveHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAliveHealthRequest create the request corresponding to the alive action endpoint of the health resource.
func (c *Client) NewAliveHealthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
