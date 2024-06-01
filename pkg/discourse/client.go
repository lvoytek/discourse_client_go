package discourse

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client   *http.Client
	host     string
	apiKey   string
	username string
}

func NewClient(DiscourseURL string, APIKey string, Username string) *Client {
	client := &http.Client{}

	return &Client{
		client:   client,
		host:     strings.TrimRight(DiscourseURL, "/"),
		apiKey:   APIKey,
		username: Username,
	}
}

func NewAnonymousClient(DiscourseURL string) *Client {
	return NewClient(DiscourseURL, "", "")
}

func (client *Client) Post(endpoint string, data []byte) error {
	return client.sendWithNoReturn("POST", endpoint, data)
}

func (client *Client) Put(endpoint string, data []byte) error {
	return client.sendWithNoReturn("PUT", endpoint, data)
}

func (client *Client) Delete(endpoint string, data []byte) error {
	return client.sendWithNoReturn("DELETE", endpoint, data)
}

func (client *Client) Get(endpoint string, data []byte) ([]byte, error) {
	return client.sendWithBodyJSONReturn("GET", endpoint, data)
}

func (client *Client) sendWithNoReturn(method string, endpoint string, data []byte) error {
	res, err := client.send(method, endpoint, data)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return createRequestError(res.StatusCode, res.Request.RequestURI)
	}

	return nil

}

func (client *Client) sendWithBodyJSONReturn(method string, endpoint string, data []byte) ([]byte, error) {
	res, err := client.send(method, endpoint, data)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, createRequestError(res.StatusCode, res.Request.RequestURI)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (client *Client) send(method string, endpoint string, data []byte) (*http.Response, error) {
	urlString := fmt.Sprintf("%s/%s.json", client.host, endpoint)

	if client.apiKey != "" && client.username != "" {
		auth := url.Values{}
		auth.Set("Api-Key", client.apiKey)
		auth.Set("Api-Username", client.username)
		urlString = fmt.Sprintf("%s?%s", urlString, auth.Encode())
	}

	req, err := http.NewRequest(method, urlString, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createRequestError(status int, requestURL string) error {
	return fmt.Errorf("HTTP Status Error: %d from URL %s", status, requestURL)
}
