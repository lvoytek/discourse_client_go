package discourse

import "net/http"

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
		host:     DiscourseURL,
		apiKey:   APIKey,
		username: Username,
	}
}

func NewAnonymousClient(DiscourseURL string) *Client {
	client := &http.Client{}

	return &Client{
		client:   client,
		host:     DiscourseURL,
		apiKey:   "",
		username: "",
	}
}
