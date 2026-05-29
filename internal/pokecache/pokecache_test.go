package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 10 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "1",
			value: []byte("1"),
		},
		{
			key:   "2",
			value: []byte("2"),
		},
	}

	for i, case_ := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(case_.key, case_.value)
			if val, ok := cache.Get(case_.key); !ok || string(val) != string(case_.value) {
				t.Errorf("expected %s to be in cache \n got %s, ok: %v", case_.key, val, ok)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const BaseTime = 5 * time.Millisecond
	const WaitTime = BaseTime + 5*time.Millisecond
	cache := NewCache(BaseTime)
	cache.Add("example.com", []byte("testdata"))
	val, ok := cache.Get("example.com")
	if !ok || string(val) != "testdata" {
		t.Errorf("expected testdata to be in cache, got %s, ok: %v", val, ok)
	}

	time.Sleep(WaitTime)
	_, ok = cache.Get("example.com")
	if ok {
		t.Errorf("expected testdata to be removed from cache")
	}
}
