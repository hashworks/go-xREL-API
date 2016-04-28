package xrel

import "testing"

func TestGetComments(t *testing.T) {
	const page = 1
	const perPage = 6

	_, err := GetComments("notExistingId", false, perPage, page)
	testRateLimit(t, err)
	test404Request(t, err)

	// >10 comments
	// https://www.xrel.to/comments/1124872/London-Has-Fallen-German-AC3-Dubbed-720p-WebHD-h264-PsO.html
	comments, err := GetComments("ab842dd3112a08", false, perPage, page)
	testRateLimit(t, err)
	if err != nil {
		t.Error(err.Error())
	} else {
		if comments.Pagination.CurrentPage != page {
			t.Errorf("Received unexpected page %d.", comments.Pagination.CurrentPage)
		}
		if comments.Pagination.PerPage != perPage {
			t.Errorf("Received unexpected per page count of %d.", comments.Pagination.PerPage)
		}
		if len(comments.List) == 0 {
			t.Error("Received unexpected comment count of 0.")
		} else {
			if comments.List[0].ID == "" {
				t.Errorf("Didn't expect the first comment not to have any id.")
			}
		}
	}
}
