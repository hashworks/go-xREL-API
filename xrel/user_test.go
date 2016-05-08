package xrel

import "testing"

func TestGetUserInfo(t *testing.T) {
	configureOAuth2ForTests(t)
	info, err := GetUserInfo()
	if err != nil {
		t.Error("Unexpected error: " + err.Error())
	} else {
		if info.ID == "" {
			t.Error("Expected an user id.")
		}
	}
}
