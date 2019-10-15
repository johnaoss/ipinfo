// Package ipinfo provides a client implementation of ipinfo.io's public facing
// api to obtain general info about a specified IP.
package ipinfo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const apiURL = "https://ipinfo.io"

var ErrRateLimited = errors.New("rate limited")

// Client is an API client with authorization already set.
type Client struct {
	bearer string
	client *http.Client
}

// NewClient returns a new client instantiated with the given bearer token.
func NewClient(token string) *Client {
	return &Client{
		bearer: "Bearer " + token,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func getrequest(kind, ip string) (*http.Request, error) {
	var req *http.Request
	var err error

	if kind != "" {
		req, err = http.NewRequest(http.MethodGet, apiURL+"/"+ip+"/"+kind, nil)
	} else {
		req, err = http.NewRequest(http.MethodGet, apiURL+"/"+ip, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) getbody(kind, ip string) ([]byte, error) {
	req, err := getrequest(kind, ip)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.bearer)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("%v: %w", err, ErrRateLimited)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, err
}

// Info returns all the possible info obtained about a given IP.
func (c *Client) Info(ip string) (*FullResponse, error) {
	body, err := c.getbody("", ip)
	if err != nil {
		return nil, err
	}
	r := &FullResponse{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Location returns only the geolocation details about an IP.
func (c *Client) Location(ip string) (*Location, error) {
	body, err := c.getbody("geo", ip)
	if err != nil {
		return nil, err
	}

	l := &Location{}
	if err := json.Unmarshal(body, l); err != nil {
		return nil, err
	}
	return l, nil
}

// Info returns the results of the unauthorized API lookup. If the IP argument
// is the empty string, it will return the information about the calling IP.
func Info(ip string) (*Response, error) {
	req, err := getrequest("", ip)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("%v: %w", err, ErrRateLimited)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	r := &Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}
	return r, nil
}
