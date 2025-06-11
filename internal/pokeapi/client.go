package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) getRequest(url string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Error while requesting PokeAPI: %w", err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("Request failed with status code: %d and body:\n %s\n", res.StatusCode, body)
	}
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to read data: %w and body:\n %s", err, body)
	}
	return body, nil
}
