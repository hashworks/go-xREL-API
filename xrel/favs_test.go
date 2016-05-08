package xrel

import (
	"strconv"
	"testing"
)

func TestFavsLists(t *testing.T) {
	configureOAuth2ForTests(t)
	favLists, err := GetFavsLists()
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}
	for i := range favLists {
		if favLists[i].ID == 0 {
			t.Fatal("Didn't expected any ID to be empty.")
		}
	}
	favEntries, err := GetFavsListEntries(strconv.Itoa(favLists[0].ID), true)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}
	for i := range favEntries {
		if favEntries[i].Title == "" {
			t.Fatal("Didn't expected any title to be empty.")
		}
	}
}
