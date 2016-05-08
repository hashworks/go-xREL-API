package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

/*
SearchReleases searches for Scene and P2P releases. Please note that additional search rate limiting applies.
See http://www.xrel.to/wiki/2727/api-rate-limiting.html

	query					Search keyword.	(required)
	includeScene		:= false	If true, Scene releases will be included in the search results.
	includeP2P		:= false	If true, P2P releases will be included in the search results.
	limit			:= 25		Number of returned search results. Maximum and default 25.

http://www.xrel.to/wiki/6320/api-search-releases.html
*/
func SearchReleases(query string, includeScene, includeP2P bool, limit int) (types.ReleaseSearchResult, error) {
	var (
		searchResult types.ReleaseSearchResult
		err          error
	)

	if query == "" {
		err = types.NewError("client", "argument_missing", "query", "")
	} else {
		form := url.Values{}
		form.Add("q", url.QueryEscape(query))
		if includeScene {
			form.Add("scene", "1")
		} else {
			form.Add("scene", "0")
		}
		if includeP2P {
			form.Add("p2p", "1")
		} else {
			form.Add("p2p", "0")
		}
		if limit != 0 {
			if limit < 1 {
				limit = 1
			}
			if limit > 25 {
				limit = 25
			}
			form.Add("limit", strconv.Itoa(limit))
		}
		var req *http.Request
		req, err = getRequest("GET", apiURL+"search/releases.json?"+form.Encode(), nil)
		if err == nil {
			var response *http.Response
			client := http.DefaultClient
			response, err = client.Do(req)
			if err == nil {
				defer response.Body.Close()
				err = checkResponse(response)
				if err == nil {
					var bytes []byte
					bytes, err = ioutil.ReadAll(response.Body)
					if err == nil {
						err = json.Unmarshal(bytes, &searchResult)
					}
				}
			}
		}
	}

	return searchResult, err
}

/*
SearchExtInfos searches for ExtInfos. Please note that additional search rate limiting applies.
See http://www.xrel.to/wiki/2727/api-rate-limiting.html

	query			Search keyword.
	extInfoType	:= ""	One of: movie|tv|game|console|software|xxx - or leave empty to search Ext Infos of all types.
	limit		:= 25	Number of returned search results. Maximum and default 25.

http://www.xrel.to/wiki/6319/api-search-ext-info.html
*/
func SearchExtInfos(query, extInfoType string, limit int) (types.ExtInfoSearchResult, error) {
	var (
		searchResult types.ExtInfoSearchResult
		err          error
	)

	if query == "" {
		err = types.NewError("client", "argument_missing", "query", "")
	} else {
		form := url.Values{}
		form.Add("q", query)
		if limit != 0 {
			if limit < 1 {
				limit = 1
			}
			if limit > 25 {
				limit = 25
			}
			form.Add("limit", strconv.Itoa(limit))
		}
		switch extInfoType {
		case "":
		case "movie", "tv", "game", "console", "software", "xxx":
			form.Add("type", extInfoType)
		default:
			err = types.NewError("client", "invalid_argument", "extInfoType", "")
		}
		if err == nil {
			var req *http.Request
			req, err = getRequest("GET", apiURL+"search/ext_info.json?"+form.Encode(), nil)
			if err == nil {
				var response *http.Response
				client := http.DefaultClient
				response, err = client.Do(req)
				if err == nil {
					defer response.Body.Close()
					err = checkResponse(response)
					if err == nil {
						var bytes []byte
						bytes, err = ioutil.ReadAll(response.Body)
						if err == nil {
							err = json.Unmarshal(bytes, &searchResult)
						}
					}
				}
			}
		}
	}

	return searchResult, err
}
