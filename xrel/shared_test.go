package xrel

import (
	"github.com/hashworks/go-xREL-API/xrel/types"
	"testing"
)

func testRateLimit(t *testing.T, err error) {
	if eErr, ok := err.(*types.Error); ok {
		if eErr.Code == "ratelimit_reached" {
			t.Fatal("Rate limit reached, stopping.")
		}
	}
}

func TestGetRequest(t *testing.T) {
	_, err := getRequest("POST", "https://hash.wórks", nil)
	if err != nil {
		t.Error("Didn't expected an error: " + err.Error())
	}
}
