package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetCachedLocationArea(t *testing.T) {
	const timing = 20 * time.Second
	const baseAreaUrl = baseUrl + "/location-area"
	const randMin = 20
	const randMax = 1088
	randAreaUrl := baseAreaUrl + fmt.Sprintf("?offset=%d&limit=20", rand.Intn(randMax-randMin)+randMin)
	client := NewClient(timing, timing)

	nilAreaResponse, err := client.GetLocationArea(nil)
	if err != nil {
		t.Errorf("unexpected error, nil request didn't succeeded")
		return
	}
	randAreaResponse, err := client.GetLocationArea(&randAreaUrl)
	if err != nil {
		t.Errorf("unexpected error, rand request didn't succeeded")
		return
	}

	cases := []struct {
		key string
		val LocationAreaResponse
	}{
		{
			key: baseAreaUrl,
			val: nilAreaResponse,
		},
		{
			key: randAreaUrl,
			val: randAreaResponse,
		},
	}
	for _, expected := range cases {
		var cachedResp LocationAreaResponse
		cachedBytes, ok := client.Cache.Get(expected.key)

		if !ok {
			t.Errorf("expected to find key")
			return
		}
		json.Unmarshal(cachedBytes, &cachedResp)
		expectedLast := expected.val.Results[len(expected.val.Results)-1]
		cachedLast := cachedResp.Results[len(cachedResp.Results)-1]

		if cachedResp.Count != expected.val.Count ||
			cachedResp.Next != expected.val.Next ||
			cachedResp.Previous != expected.val.Previous ||
			cachedLast.Name != expectedLast.Name || cachedLast.URL != expectedLast.URL {
			t.Errorf("expected to be the same value")
			return
		}
	}
}
