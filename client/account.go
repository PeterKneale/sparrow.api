package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ReadAccountPath computes a request path to the read action of Account.
func ReadAccountPath() string {
	return fmt.Sprintf("/account")
}

// read current account
func (c *Client) ReadAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReadAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadAccountRequest create the request corresponding to the read action endpoint of the Account resource.
func (c *Client) NewReadAccountRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateAccountPath computes a request path to the update action of Account.
func UpdateAccountPath() string {
	return fmt.Sprintf("/account")
}

// Update account
func (c *Client) UpdateAccount(ctx context.Context, path string, payload *UpdateAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAccountRequest create the request corresponding to the update action endpoint of the Account resource.
func (c *Client) NewUpdateAccountRequest(ctx context.Context, path string, payload *UpdateAccountPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
