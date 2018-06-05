// Package noip provides a convinent way to embed a no-ip.com client in your go program.
package noip

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseurl          = "https://dynupdate.no-ip.com/nic/update"
	defaultUseragent = "unprofession-al/noip/v1.0 maintainer@domain.com"
)

// Client holds all info and logic to talk to the dynupdate.no-ip.com API
type Client struct {
	user      string
	pass      string
	params    url.Values
	useragent string
	myip      string
}

// New creates a new no-ip.com client. If `myip` is an empty string the myip url param is omitted
// completely in the resulting API request url. If `useragent` is an empty string a default user agent
// will be written to the header in order to fulfill the API requirements.
func New(user string, pass string, hostname string, myip string, useragent string) *Client {
	params := url.Values{}
	params.Set("hostname", hostname)
	if myip != "" {
		params.Set("myip", myip)
	}

	c := &Client{
		user:      user,
		pass:      pass,
		useragent: useragent,
		params:    params,
	}

	if c.useragent == "" {
		c.useragent = defaultUseragent
	}

	return c
}

func (c *Client) update() (string, error) {
	url := fmt.Sprintf("%s?%s", baseurl, c.params.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Set("User-Agent", c.useragent)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	return body.String(), nil
}

// Run starts a goroutine which calls the API in a given interval (seconds). If `log` is set to
// `true`, each result of an API call will be printed to STDOUT.
func (c *Client) Run(interval int, log bool) {
	go func() {
		for {
			out, err := c.update()
			if log {
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("NO-IP is updated:", out)
				}
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}
