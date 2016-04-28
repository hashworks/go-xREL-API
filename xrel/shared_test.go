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

func test404Request(t *testing.T, err error) {
	if err == nil {
		t.Fatal("Expected an error.")
	}
	if eErr, ok := err.(*types.Error); ok {
		if eErr.Code != "id_not_found" {
			t.Errorf("Expected error code 'id_not_found', got code %s with message\n%s.", eErr.Code, eErr.Error())
		}
	} else {
		t.Errorf("Expected an extended error. Got:%s\n", err.Error())
	}
}

func TestGenerateGetParametersString(t *testing.T) {
	parameters := make(map[string]string)
	parameters["foo"] = "bar"
	parameters["y"] = "1"
	rString := generateGetParametersString(parameters)
	if rString != "?foo=bar&y=1" && rString != "?y=1&foo=bar" {
		t.Errorf("Expected '?foo=bar&y=1' or '?y=1&foo=bar', got '%s'.", rString)
	}
}
