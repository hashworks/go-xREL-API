package xrel

import "testing"

func TestGetP2PReleaseInfo(t *testing.T) {
	_, err := GetP2PReleaseInfo("notExistingId", true)
	testRateLimit(t, err)
	test404Request(t, err)

	// Scene release
	// https://www.xrel.to/movie-nfo/60557/Keinohrhasen-German-AC3-DVDRip-XviD-CRUCiAL.html
	_, err = GetP2PReleaseInfo("f638d1cfec8d", true)
	testRateLimit(t, err)
	test404Request(t, err)

	// https://www.xrel.to/p2p/11238-Killer-Elite-2011-German-AC3D-DL-720p-HDTV-x264-TwixX/nfo.html
	relById, err := GetP2PReleaseInfo("6dbb52db2be6", true)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	relByDirname, err := GetP2PReleaseInfo("Killer.Elite.2011.German.AC3D.DL.720p.HDTV.x264-TwixX", false)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	if relById.Id == "" {
		t.Error("Didn't expect the id to be empty.")
	}
	if relById.Id != relByDirname.Id {
		t.Error("Expected to receive the same ids by id and by dirname but tjey are different.")
	}
	if relById.Dirname == "" {
		t.Error("Didn't expect the dirname to be empty.")
	}
	if relById.Dirname != relByDirname.Dirname {
		t.Error("Expected to receive the same dirname by id and by dirname but they are different.")
	}
}

// HD-1080p
const catId = "968f458a2"

func TestGetP2PCategories(t *testing.T) {
	cats, err := GetP2PCategories()
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(cats) == 0 {
			t.Error("Didn't expect the category list to be empty.")
		} else {
			exists := false
			for i := range cats {
				if cats[i].Id == catId {
					exists = true
					break
				}
			}
			if !exists {
				t.Error("Expected category %s to exist.")
			}
		}
	}
}

func TestGetP2PReleases(t *testing.T) {
	const page = 1
	const perPage = 6
	//	_, err := GetP2PReleases(perPage, page, "", "", "notExistingId")
	//	testRateLimit(t, err)
	//	test404Request(t, err)
	//	_, err = GetP2PReleases(perPage, page, "", "notExistingId", "")
	//	testRateLimit(t, err)
	//	test404Request(t, err)
	//	_, err = GetP2PReleases(perPage, page, "notExistingId", "", "")
	//	testRateLimit(t, err)
	//	test404Request(t, err)

	// Test category request
	releases, err := GetP2PReleases(perPage, page, catId, "", "")
	testRateLimit(t, err)

	if err != nil {
		t.Error(err.Error())
	} else {
		if releases.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", releases.Pagination.CurrentPage)
		}
		if releases.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", releases.Pagination.PerPage)
		}
		if len(releases.List) == 0 {
			t.Error("Received unexpected release count of 0.")
		} else {
			if releases.List[0].Id == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}

	// Test group request
	// https://www.xrel.to/p2p/11238-Killer-Elite-2011-German-AC3D-DL-720p-HDTV-x264-TwixX/nfo.html
	const groupId = "13c6f87f6f"
	releases, err = GetP2PReleases(perPage, page, "", groupId, "")
	testRateLimit(t, err)

	if err != nil {
		t.Error(err.Error())
	} else {
		if releases.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", releases.Pagination.CurrentPage)
		}
		if releases.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", releases.Pagination.PerPage)
		}
		if len(releases.List) == 0 {
			t.Error("Received unexpected release count of 0.")
		} else {
			if releases.List[0].Id == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}

	// Test ExtInfo request
	// https://www.xrel.to/movie/36977/Easy-Virtue-Eine-unmoralische-Ehefrau.html
	releases, err = GetP2PReleases(perPage, page, "", "", "f22936fb9071")
	testRateLimit(t, err)

	if err != nil {
		t.Error(err.Error())
	} else {
		if releases.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", releases.Pagination.CurrentPage)
		}
		if releases.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", releases.Pagination.PerPage)
		}
		if len(releases.List) == 0 {
			t.Error("Received unexpected release count of 0.")
		} else {
			if releases.List[0].Id == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}
}
