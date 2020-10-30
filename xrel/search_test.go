package xrel

import "testing"

// Important: Limit this to a maximum of 2 search tests to avoid the rate limiting of max. 2 searches / 10 seconds!

func TestSearchReleases(t *testing.T) {
	const limit = 6

	rl, err := SearchReleases("Limitless", true, false, limit)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if rl.Total != limit {
			t.Errorf("Expected the total release count to be %d, was %d.", limit, rl.Total)
		}
		if len(rl.SceneResults) != limit {
			t.Errorf("Expected the scene release count to be %d, was %d.", limit, len(rl.SceneResults))
		} else {
			if rl.SceneResults[0].ID == "" {
				t.Error("Didn't expect the id of the first release to be empty.")
			}
		}
		if len(rl.P2PResults) != 0 {
			t.Errorf("Expected the p2p release count to be 0, was %d.", len(rl.P2PResults))
		}
	}
}

func TestSearchExtInfos(t *testing.T) {
	const limit = 1

	ei, err := SearchExtInfos("Limitless", "movie", limit)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if ei.Total != limit {
			t.Errorf("Expected the total result count to be %d, was %d.", limit, ei.Total)
		}
		if len(ei.Results) != limit {
			t.Errorf("Expected the result count to be %d, was %d.", limit, len(ei.Results))
		} else {
			if ei.Results[0].ID == "" {
				t.Error("Didn't expect the id of the first result to be empty.")
			}
		}
	}
}
