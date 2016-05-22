package crate

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var Scheme = "crate"

type Client struct {
	url *url.URL
}

func NewCrateClient(downloadUrl *url.URL) (*Client, error) {
	return &Client{url: downloadUrl}, nil
}

func (c *Client) Get(w io.Writer) error {
	resp, err := http.Get(c.url.String())
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to get %s: %s", c.url.String(), err.Error()))
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to read response from crate: %s", err.Error()))
	}
	_, err = w.Write(data)
	return err
}

func (c *Client) Url() *url.URL {
	return c.url
}

func (c *Client) String() string {
	return fmt.Sprintf("<CrateClient url=%s>", c.url.String())
}
