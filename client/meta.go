package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AliveMetaPath computes a request path to the alive action of meta.
func AliveMetaPath() string {
	return fmt.Sprintf("/api/health/alive")
}

// Perform aliveness check.
func (c *Client) AliveMeta(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAliveMetaRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAliveMetaRequest create the request corresponding to the alive action endpoint of the meta resource.
func (c *Client) NewAliveMetaRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ReadyMetaPath computes a request path to the ready action of meta.
func ReadyMetaPath() string {
	return fmt.Sprintf("/api/health/ready")
}

// Perform readiness check.
func (c *Client) ReadyMeta(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReadyMetaRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadyMetaRequest create the request corresponding to the ready action endpoint of the meta resource.
func (c *Client) NewReadyMetaRequest(ctx context.Context, path string) (*http.Request, error) {
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

// RootMetaPath computes a request path to the root action of meta.
func RootMetaPath() string {
	return fmt.Sprintf("/api")
}

// Perform root check.
func (c *Client) RootMeta(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRootMetaRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRootMetaRequest create the request corresponding to the root action endpoint of the meta resource.
func (c *Client) NewRootMetaRequest(ctx context.Context, path string) (*http.Request, error) {
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
