package test_test

import (
	"testing"

	"github.com/albatarnik/httpcache"
	"github.com/albatarnik/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
