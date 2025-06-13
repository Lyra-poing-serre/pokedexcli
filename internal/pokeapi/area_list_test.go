package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetCachedLocationArea(t *testing.T) {
	const interval = 5 * time.Second
	const timeout = 5 * time.Second
	const baseAreaUrl = BaseUrl + "/location-area"
	const randMin = 20
	const randMax = 1088
	randAreaUrl := baseAreaUrl + fmt.Sprintf("?offset=%d&limit=20", rand.Intn(randMax-randMin)+randMin)
	client := NewClient(timeout, interval)

	nil_area_response, err := client.GetLocationArea(nil)
	if err != nil {
		t.Errorf("unexpected error, nil request didn't succeeded")
		return
	}
	rand_area_response, err := client.GetLocationArea(&randAreaUrl)
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
			val: nil_area_response,
		},
		{
			key: randAreaUrl,
			val: rand_area_response,
		},
	}
	for _, expected_cache := range cases {
		var cached_unmar_resp LocationAreaResponse
		fmt.Println(client.cache)
		fmt.Print("Testing : ")
		fmt.Println(expected_cache)
		cached_resp, ok := client.cache.Get(expected_cache.key)

		if !ok {
			t.Errorf("expected to find key")
			return
		}
		json.Unmarshal(cached_resp, &cached_unmar_resp)

		if cached_unmar_resp.Count != expected_cache.val.Count ||
			cached_unmar_resp.Next != expected_cache.val.Next ||
			cached_unmar_resp.Previous != expected_cache.val.Previous ||
			len(cached_unmar_resp.Results) != len(expected_cache.val.Results) {
			t.Errorf("expected to be the same value")
			return
		}
	}
}
