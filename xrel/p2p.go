package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
GetP2PReleaseInfo returns information about a single release, specified by the complete dirname or an API release id.

http://www.xrel.to/wiki/3697/api-p2p-rls-info.html
*/
func GetP2PReleaseInfo(query string, isID bool) (types.P2PRelease, error) {
	var (
		p2pReleaseStruct types.P2PRelease
		err              error
	)

	if query == "" {
		err = types.NewError("client", "argument_missing", "query", "")
	} else {
		if isID {
			query = "?id=" + query
		} else {
			query = "?dirname=" + query
		}
		req, err := getRequest("GET", apiURL+"p2p/rls_info.json"+query, nil)
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
						err = json.Unmarshal(bytes, &p2pReleaseStruct)
					}
				}
			}
		}
	}

	return p2pReleaseStruct, err
}

/*
GetP2PReleases allows to browse P2P/non-scene releases.

	perPage		:= 25	Number of releases per page. Min. 5, max. 100.
	page     	:= 1	Page number (1 to N).

	Set only one of the following:
	categoryID	:= ""	P2P category ID from GetP2PCategories()
	groupID		:= ""	P2P release group ID
	extInfoID	:= ""	Ext Info ID


http://www.xrel.to/wiki/3699/api-p2p-releases.html
*/
func GetP2PReleases(perPage, page int, categoryID, groupID, extInfoID string) (types.P2PReleases, error) {
	var p2pReleasesStruct types.P2PReleases

	form := url.Values{}

	if perPage != 0 {
		if perPage < 5 {
			perPage = 5
		}
		if perPage > 100 {
			perPage = 100
		}
		form.Add("per_page", strconv.Itoa(perPage))
	}
	if page > 0 {
		form.Add("page", strconv.Itoa(page))
	}
	if categoryID != "" {
		form.Add("category_id", categoryID)
	} else if groupID != "" {
		form.Add("group_id", groupID)
	} else if extInfoID != "" {
		form.Add("ext_info_id", extInfoID)
	}

	req, err := getRequest("GET", apiURL+"p2p/releases.json?"+form.Encode(), nil)
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
					err = json.Unmarshal(bytes, &p2pReleasesStruct)
				}
			}
		}
	}

	return p2pReleasesStruct, err
}

/*
GetP2PCategories returns a list of available P2P release categories and their IDs. You can use the category ID in GetP2PReleases().

http://www.xrel.to/wiki/3698/api-p2p-categories.html
*/
func GetP2PCategories() ([]types.P2PCategory, error) {
	var err error

	// According to xREL we should cache the results for 24h
	currentUnix := time.Now().Unix()
	if types.Config.LastP2PCategoryRequest == 0 || currentUnix-types.Config.LastP2PCategoryRequest > 86400 || len(types.Config.P2PCategories) == 0 {
		var req *http.Request
		req, err = getRequest("GET", apiURL+"p2p/categories.json", nil)
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
						err = json.Unmarshal(bytes, &types.Config.P2PCategories)
						if err == nil {
							types.Config.LastP2PCategoryRequest = currentUnix
						}
					}
				}
			}
		}
	}

	return types.Config.P2PCategories, err
}
