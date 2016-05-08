package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
GetFavsLists returns a list of all the current user's favorite lists.
Requires user OAuth2 authentication.

http://www.xrel.to/wiki/1754/api-favs-lists.html
*/
func GetFavsLists() ([]types.FavList, error) {
	var favLists []types.FavList

	client := http.DefaultClient
	request, err := getOAuth2Request("POST", apiURL+"favs/lists.json", nil)
	if err == nil {
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
			err = checkResponse(response)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(bytes, &favLists)
				}
			}
		}
	}

	return favLists, err
}

/*
GetFavsListEntries returns entries of a favorite list.
Requires user OAuth2 authentication.

	favsListID			The favorite list ID, as obtained through GetFavsLists().
	getReleases	:= false	If true, an inline list of unread(!) releases will be returned with each ext_info entry.

http://www.xrel.to/wiki/1823/api-favs-list-entries.html
*/
func GetFavsListEntries(favsListID string, getReleases bool) ([]types.ExtendedExtInfo, error) {
	var extendedExtInfos []types.ExtendedExtInfo

	if favsListID == "" {
		return extendedExtInfos, types.NewError("client", "argument_missing", "favsListId", "")
	}

	form := url.Values{}
	form.Add("id", favsListID)
	if getReleases {
		form.Add("get_releases", "true")
	}
	request, err := getOAuth2Request("POST", apiURL+"favs/list_entries.json", strings.NewReader(form.Encode()))
	if err == nil {
		client := http.DefaultClient
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
			err = checkResponse(response)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(bytes, &extendedExtInfos)
				}
			}
		}
	}

	return extendedExtInfos, err
}

/*
AddFavsListEntry adds an ExtInfo to a favorite list.
Requires user OAuth2 authentication.

	favsListID	The favorite list ID, as obtained through GetFavsLists().
	extInfoID	The Ext Info ID, as obtained through other API calls.

http://www.xrel.to/wiki/6316/api-favs-list-addentry.html
*/
func AddFavsListEntry(favsListID, extInfoID string) (types.FavListEntryModificationResult, error) {
	var (
		favListAddEntryResult types.FavListEntryModificationResult
		err                   error
	)

	if favsListID == "" {
		return favListAddEntryResult, types.NewError("client", "argument_missing", "favsListId", "")
	}
	if extInfoID == "" {
		return favListAddEntryResult, types.NewError("client", "argument_missing", "extInfoId", "")
	}

	form := url.Values{}
	form.Add("id", favsListID)
	form.Add("ext_info_id", extInfoID)
	request, err := getOAuth2Request("POST", apiURL+"favs/list_addentry.json", strings.NewReader(form.Encode()))
	if err == nil {
		client := http.DefaultClient
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
			err = checkResponse(response)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(bytes, &favListAddEntryResult)
				}
			}
		}
	}

	return favListAddEntryResult, err
}

/*
RemoveFavsListEntry removes an ExtInfo from a favorite list.
Requires user OAuth2 authentication.

	favsListID	The favorite list ID, as obtained through GetFavsLists().
	extInfoID	The ExtInfo ID, as obtained through other API calls.

http://www.xrel.to/wiki/6317/api-favs-list-delentry.html
*/
func RemoveFavsListEntry(favsListID, extInfoID string) (types.FavListEntryModificationResult, error) {
	var (
		favListRemoveEntryResult types.FavListEntryModificationResult
		err                      error
	)

	if favsListID == "" {
		return favListRemoveEntryResult, types.NewError("client", "argument_missing", "favsListId", "")
	}
	if extInfoID == "" {
		return favListRemoveEntryResult, types.NewError("client", "argument_missing", "extInfoId", "")
	}

	form := url.Values{}
	form.Add("id", favsListID)
	form.Add("ext_info_id", extInfoID)
	request, err := getOAuth2Request("POST", apiURL+"favs/list_delentry.json", strings.NewReader(form.Encode()))
	if err == nil {
		client := http.DefaultClient
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
			err = checkResponse(response)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(bytes, &favListRemoveEntryResult)
				}
			}
		}
	}

	return favListRemoveEntryResult, err
}

/*
MarkFavsListEntryAsRead marks a release on a favorite list as read.
Requires user OAuth2 authentication.

	favsListID	The favorite list ID, as obtained through GetFavsLists().
	releaseID	The API release ID, as obtained through other API calls.
	isP2P		Is the release ID a P2P release?

https://www.xrel.to/wiki/6344/api-favs-list-markread.html
*/
func MarkFavsListEntryAsRead(favsListID, releaseID string, isP2P bool) (types.ShortFavList, error) {
	var (
		shortFavList types.ShortFavList
		err          error
	)

	if favsListID == "" {
		return shortFavList, types.NewError("client", "argument_missing", "favsListId", "")
	}
	if releaseID == "" {
		return shortFavList, types.NewError("client", "argument_missing", "releaseId", "")
	}

	form := url.Values{}
	form.Add("id", favsListID)
	form.Add("release_id", releaseID)
	if isP2P {
		form.Add("type", "p2p_rls")
	} else {
		form.Add("type", "release")
	}
	request, err := getOAuth2Request("POST", apiURL+"favs/list_markread.json", strings.NewReader(form.Encode()))
	if err == nil {
		client := http.DefaultClient
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
			err = checkResponse(response)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(bytes, &shortFavList)
				}
			}
		}
	}

	return shortFavList, err
}
