package crate

import (
	"fmt"
	"net/url"
)

var Scheme = "crate"

type Client struct {
	url *url.URL
}

func NewCrateClient(downloadUrl *url.URL) (*Client, error) {
	return &Client{url: downloadUrl}, nil
}

func (c *Client) Get() (string, error) {
	return "", nil
}

func (c *Client) String() string {
	return fmt.Sprintf("<CrateClient url=%s>", c.url.String())
}
