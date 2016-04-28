package xrel

import "testing"

func TestGetExtInfo(t *testing.T) {
	_, err := GetExtInfo("notExistingId")
	testRateLimit(t, err)
	test404Request(t, err)

	// https://www.xrel.to/game/78663/L-A-Noire.html
	const id = "f25c556d13347"
	eInfo, err := GetExtInfo(id)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if eInfo.Id != id {
			t.Errorf("Expected id '%s', received '%s'.", id, eInfo.Id)
		}
		if eInfo.Type != "game" {
			t.Errorf("Expected type 'game', received '%s'.", eInfo.Type)
		}
		if eInfo.Title != "L.A. Noire" {
			t.Errorf("Expected title 'L.A. Noire', received '%s'.", eInfo.Title)
		}
	}
}

func TestGetExtInfoMedia(t *testing.T) {
	_, err := GetExtInfoMedia("notExistingId")
	testRateLimit(t, err)
	test404Request(t, err)

	// https://www.xrel.to/movie/132163/Straight-Outta-Compton.html
	const id = "08f068b020443"
	eInfoMedia, err := GetExtInfoMedia(id)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	}
	if len(eInfoMedia) == 0 {
		t.Error("Unexpectingly received no ExtInfoMedias.")
	} else {
		if eInfoMedia[0].Type == "" {
			t.Error("Expected a proper type, got nothing.")
		}
	}
}
