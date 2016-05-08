package xrel

import "testing"

func TestGetUpcomingTitles(t *testing.T) {
	t.Parallel()
	upcomingTitles, err := GetUpcomingTitles("us")
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(upcomingTitles) == 0 {
			t.Error("Expected it to return any upcoming titles, slice is empty.")
		} else {
			for i := range upcomingTitles {
				if upcomingTitles[i].Title == "" {
					t.Error("Expected any upcoming title to have a valid Title attribute.")
					t.Log(upcomingTitles[i])
				}
			}
		}
	}
}
