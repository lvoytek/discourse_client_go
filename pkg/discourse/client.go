package discourse

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

func (client *Client) PostWithReturn(endpoint string, data []byte) ([]byte, error) {
	return client.sendWithBodyJSONReturn("POST", endpoint, data)
}

func (client *Client) Put(endpoint string, data []byte) error {
	return client.sendWithNoReturn("PUT", endpoint, data)
}

func (client *Client) PutWithReturn(endpoint string, data []byte) ([]byte, error) {
	return client.sendWithBodyJSONReturn("PUT", endpoint, data)
}

func (client *Client) Delete(endpoint string, data []byte) error {
	return client.sendWithNoReturn("DELETE", endpoint, data)
}

func (client *Client) Get(endpoint string) ([]byte, error) {
	return client.sendWithBodyJSONReturn("GET", endpoint, []byte{})
}

func (client *Client) GetWithBodyJSONInput(endpoint string, data []byte) ([]byte, error) {
	return client.sendWithBodyJSONReturn("GET", endpoint, data)
}

func (client *Client) sendWithNoReturn(method string, endpoint string, data []byte) error {
	res, err := client.send(method, endpoint, data)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		if err != nil {
			return createRequestError(res.StatusCode, []byte{})
		} else {
			return createRequestError(res.StatusCode, body)
		}
	}

	return nil

}

func (client *Client) sendWithBodyJSONReturn(method string, endpoint string, data []byte) ([]byte, error) {
	res, err := client.send(method, endpoint, data)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, createRequestError(res.StatusCode, body)
	}

	return body, nil
}

func (client *Client) send(method string, endpoint string, data []byte) (*http.Response, error) {
	urlString := fmt.Sprintf("%s/%s.json", client.host, endpoint)

	req, err := http.NewRequest(method, urlString, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if client.apiKey != "" && client.username != "" {
		req.Header.Add("Api-Key", client.apiKey)
		req.Header.Add("Api-Username", client.username)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createRequestError(status int, responseData []byte) error {
	return fmt.Errorf("HTTP Status Error: %d\n%s", status, string(responseData))
}
