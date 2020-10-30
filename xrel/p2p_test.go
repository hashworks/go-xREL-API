package xrel

import "testing"

func TestGetP2PReleaseInfo(t *testing.T) {
	// https://www.xrel.to/p2p/11238-Killer-Elite-2011-German-AC3D-DL-720p-HDTV-x264-TwixX/nfo.html
	relByID, err := GetP2PReleaseInfo("6dbb52db2be6", true)
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
	if relByID.ID == "" {
		t.Error("Didn't expect the id to be empty.")
	}
	if relByID.ID != relByDirname.ID {
		t.Error("Expected to receive the same ids by id and by dirname but tjey are different.")
	}
	if relByID.Dirname == "" {
		t.Error("Didn't expect the dirname to be empty.")
	}
	if relByID.Dirname != relByDirname.Dirname {
		t.Error("Expected to receive the same dirname by id and by dirname but they are different.")
	}
}

// HD-1080p
const catID = "968f458a2"

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
				if cats[i].ID == catID {
					exists = true
					break
				}
			}
			if !exists {
				t.Error("Expected category does not exist.")
			}
		}
	}
}

func TestGetP2PReleases(t *testing.T) {
	const page = 1
	const perPage = 6

	// Test category request
	releases, err := GetP2PReleases(perPage, page, catID, "", "")
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
			if releases.List[0].ID == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}

	// Test group request
	// https://www.xrel.to/p2p/11238-Killer-Elite-2011-German-AC3D-DL-720p-HDTV-x264-TwixX/nfo.html
	const groupID = "13c6f87f6f"
	releases, err = GetP2PReleases(perPage, page, "", groupID, "")
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
			if releases.List[0].ID == "" {
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
			if releases.List[0].ID == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}
}
