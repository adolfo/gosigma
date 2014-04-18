// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package https

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Client represents HTTPS client connection with optional basic authentication
type Client struct {
	protocol         *http.Client
	username         string
	password         string
	connectTimeout   time.Duration
	readWriteTimeout time.Duration
	transport        *http.Transport
}

// NewClient returns new Client object with transport configured for https.
// Parameter tlsConfig is optional and can be nil, the default TLSClientConfig of
// http.Transport will be used in this case.
func NewClient(tlsConfig *tls.Config) *Client {
	if tlsConfig == nil {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	redirectChecker := func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		lastReq := via[len(via)-1]
		if auth := lastReq.Header.Get("Authorization"); len(auth) > 0 {
			req.Header.Add("Authorization", auth)
		}
		return nil
	}

	https := &Client{
		protocol: &http.Client{
			Transport:     tr,
			CheckRedirect: redirectChecker,
		},
		transport: tr,
	}

	tr.Dial = https.dialer

	return https
}

// NewAuthClient returns new Client object with configured https transport
// and attached authentication. Parameter tlsConfig is optional and can be nil, the
// default TLSClientConfig of http.Transport will be used in this case.
func NewAuthClient(username, password string, tlsConfig *tls.Config) *Client {
	https := NewClient(tlsConfig)
	https.username = username
	https.password = password
	return https
}

// ConnectTimeout sets connection timeout
func (c *Client) ConnectTimeout(timeout time.Duration) {
	c.connectTimeout = timeout
	c.transport.CloseIdleConnections()
}

// ReadWriteTimeout sets read-write timeout
func (c *Client) ReadWriteTimeout(timeout time.Duration) {
	c.readWriteTimeout = timeout
}

// Get performs get request to the url.
func (c Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(c.username) != 0 {
		req.SetBasicAuth(c.username, c.password)
	}

	return c.do(req)
}

// GetQuery performs get request to the url with the list of query arguments.
func (c Client) GetQuery(url string, values url.Values) (*http.Response, error) {
	if len(values) == 0 {
		return c.Get(url)
	}

	url += "?" + values.Encode()

	return c.Get(url)
}

func (c Client) do(r *http.Request) (*http.Response, error) {
	var chStop chan int

	readWriteTimeout := c.readWriteTimeout
	if readWriteTimeout > 0 {
		chStop = make(chan int)
		go func() {
			select {
			case <-time.After(readWriteTimeout):
				c.transport.CancelRequest(r)
			case <-chStop:
				return
			}
		}()
	}

	resp, err := c.protocol.Do(r)

	if chStop != nil {
		close(chStop)
	}

	return resp, err
}

func (c *Client) dialer(netw, addr string) (net.Conn, error) {
	return net.DialTimeout(netw, addr, c.connectTimeout)
}
