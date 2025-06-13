package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Lyra-poing-serre/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(interval),
	}
}

func (c *Client) getRequest(url string) (body []byte, err error) {
	body, ok := c.cache.Get(url)
	if ok {
		//fmt.Print("Retrived cache entry for : ")
		//fmt.Println(url)
		return body, nil
	}
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
	c.cache.Add(url, body)
	return body, nil
}
