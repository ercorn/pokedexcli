package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + (5 * time.Millisecond)
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area/", []byte("testdata"))

	_, exists := cache.Get("https://pokeapi.co/api/v2/location-area/")
	if !exists {
		t.Errorf("ERROR, cache miss: did not find key")
		return
	}

	time.Sleep(waitTime)

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area/")
	if ok {
		t.Errorf("ERROR, cache hit: found key")
	}
}
