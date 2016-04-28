package xrel

import "testing"

func TestGetReleaseInfo(t *testing.T) {
	_, err := GetReleaseInfo("notExistingID", true)
	testRateLimit(t, err)

	ri, err := GetReleaseInfo("Smokin.Aces.2.Assassins.Ball.UNRATED.German.2009.DVDRiP.XViD-AOE", false)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if ri.Id != "e14fc6ad27fda" {
			t.Error("Unexpected id")
		}
	}

	ri, err = GetReleaseInfo("e14fc6ad27fda", true)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if ri.Dirname != "Smokin.Aces.2.Assassins.Ball.UNRATED.German.2009.DVDRiP.XViD-AOE" {
			t.Error("Unexpected dirname")
		}
	}
}

func TestGetLatestReleases(t *testing.T) {
	_, err := GetLatestReleases(1, 6, "", "2999-12")
	testRateLimit(t, err)

	const page = 1
	const perPage = 6

	// Archive with category "Deutsche Filme und Serien"
	rl, err := GetLatestReleases(perPage, page, "7", "2016-03")
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if rl.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", rl.Pagination.CurrentPage)
		}
		if rl.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", rl.Pagination.PerPage)
		}
		if len(rl.List) == 0 {
			t.Error("Received unexpected release count of 0.")
		} else {
			if rl.List[0].Id == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}

	// Current
	rl, err = GetLatestReleases(perPage, page, "", "")
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if rl.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", rl.Pagination.CurrentPage)
		}
		if rl.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", rl.Pagination.PerPage)
		}
		if len(rl.List) == 0 {
			t.Error("Received unexpected release count of 0.")
		} else {
			if rl.List[0].Id == "" {
				t.Errorf("Didn't expect the first release not to have any id.")
			}
		}
	}
}

func TestGetReleaseFilters(t *testing.T) {
	filters, err := GetReleaseFilters()
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(filters) == 0 {
			t.Error("Expected some filters")
		}
	}
}

func TestBrowseReleaseCategory(t *testing.T) {
	const page = 1
	const perPage = 6

	rl, err := BrowseReleaseCategory("INTERNAL", "movie", perPage, page)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if rl.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", rl.Pagination.CurrentPage)
		}
		if rl.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", rl.Pagination.PerPage)
		}
		if len(rl.List) != 6 {
			t.Error("Received unexpected release count of %d.", len(rl.List))
		} else {
			for i := range rl.List {
				if rl.List[i].ExtInfo.Type != "movie" {
					t.Errorf("Didn't expect the type of release #%d not to be 'movie'. Was '%s'.", i, rl.List[i].ExtInfo.Type)
				}
			}
		}
	}

	rl, err = BrowseReleaseCategory("INTERNAL", "", perPage, page)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if rl.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", rl.Pagination.CurrentPage)
		}
		if rl.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", rl.Pagination.PerPage)
		}
		if len(rl.List) != 6 {
			t.Error("Received unexpected release count of %d.", len(rl.List))
		} else {
			if rl.List[0].Id == "" {
				t.Error("Didn't expect the id of the first release to be empty.")
			}
		}
	}
}

func TestGetReleaseCategories(t *testing.T) {
	cats, err := GetReleaseCategories()
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if len(cats) == 0 {
			t.Error("Expected some categories")
		}
	}
}

func TestGetReleaseByExtInfoID(t *testing.T) {
	_, err := GetReleaseByExtInfoID("notExistingId", 0, 0)
	testRateLimit(t, err)
	test404Request(t, err)

	const page = 1
	const perPage = 6

	rl, err := GetReleaseByExtInfoID("e5262b5349a", perPage, page)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if rl.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", rl.Pagination.CurrentPage)
		}
		if rl.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", rl.Pagination.PerPage)
		}
		if len(rl.List) != 6 {
			t.Error("Received unexpected release count of %d.", len(rl.List))
		} else {
			if rl.List[0].Id == "" {
				t.Error("Didn't expect the id of the first release to be empty.")
			}
		}
	}
}
