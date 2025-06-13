package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) GetArea(areaName string) (areaResponse AreaResponse, err error) {
	if areaName == "" {
		return AreaResponse{}, errors.New("No area input...")
	}

	body, err := c.getRequest(fmt.Sprint(baseUrl, "/location-area/", areaName))
	if err != nil {
		return AreaResponse{}, fmt.Errorf("GET request error with url: %s got this error: %w", areaName, err)
	}

	if err := json.Unmarshal(body, &areaResponse); err != nil {
		return AreaResponse{}, fmt.Errorf("Failed to decode json, got err: %w\n", err)
	}

	return areaResponse, nil
}
