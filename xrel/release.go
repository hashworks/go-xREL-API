package xrel

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
GetReleaseInfo returns information about a single release, specified by the complete dirname or an API release id.

http://www.xrel.to/wiki/1680/api-release-info.html
*/
func GetReleaseInfo(query string, isID bool) (types.Release, error) {
	var release types.Release

	if isID {
		query = "?id=" + query
	} else {
		query = "?dirname=" + query
	}

	req, err := getRequest("GET", apiURL+"release/info.json"+query, nil)
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
					err = json.Unmarshal(bytes, &release)
				}
			}
		}
	}

	return release, err
}

func getReleases(url string, form url.Values) (types.Releases, error) {
	var releases types.Releases

	req, err := getRequest("GET", url+"?"+form.Encode(), nil)
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
					err = json.Unmarshal(bytes, &releases)
				}
			}
		}
	}

	return releases, err
}

/*
GetLatestReleases returns the latest releases. Also allows to browse the archive by month.

	perPage	:= 25	Number of releases per page. Min. 5, max. 100.
	page	:= 1	Page number (1 to N).
	filter	:= ""	Filter ID (from GetReleaseFilters()) or "overview" to use the currently logged in user's overview filter.
	archive	:= ""	Empty = current releases, YYYY-MM for archive.

http://www.xrel.to/wiki/2994/api-release-latest.html
*/
func GetLatestReleases(perPage, page int, filter, archive string) (types.Releases, error) {
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
	if filter != "" {
		form.Add("filter", filter)
	}
	if archive != "" {
		form.Add("archive", archive)
	}

	return getReleases(apiURL+"release/latest.json", form)
}

/*
GetReleaseFilters returns a list of public, predefined release filters. You can use the filter ID in GetLatestReleases().

http://www.xrel.to/wiki/2996/api-release-filters.html
*/
func GetReleaseFilters() ([]types.Filter, error) {
	var err error

	currentUnix := time.Now().Unix()
	if types.Config.LastFilterRequest == 0 || currentUnix-types.Config.LastFilterRequest > 86400 || len(types.Config.Filters) == 0 {
		var req *http.Request
		req, err = getRequest("GET", apiURL+"release/filters.json", nil)
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
						err = json.Unmarshal(bytes, &types.Config.Filters)
						if err == nil {
							types.Config.LastFilterRequest = currentUnix
						}
					}
				}
			}
		}
	}

	return types.Config.Filters, err
}

/*
BrowseReleaseCategory returns scene releases from the given category.

	categoryName		Category name from GetReleaseCategories()
	extInfoType := ""	Use one of: movie|tv|game|console|software|xxx - or leave empty to browse releases of all types

http://www.xrel.to/wiki/3751/api-release-browse-category.html
*/
func BrowseReleaseCategory(categoryName, extInfoType string, perPage, page int) (types.Releases, error) {
	var (
		releasesStruct types.Releases
		err            error
	)

	if categoryName == "" {
		err = types.NewError("client", "argument_missing", "category name", "")
	} else {
		form := url.Values{}

		form.Add("category_name", categoryName)
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
		switch extInfoType {
		case "":
		case "movie", "tv", "game", "console", "software", "xxx":
			form.Add("ext_info_type", extInfoType)
		default:
			err = types.NewError("client", "invalid_argument", "extInfoType", "")
		}
		if err == nil {
			releasesStruct, err = getReleases(apiURL+"release/browse_category.json", form)
		}
	}

	return releasesStruct, err
}

/*
GetReleaseCategories returns a list of available release categories. You can use the category name in BrowseReleaseCategory().

http://www.xrel.to/wiki/6318/api-release-categories.html
*/
func GetReleaseCategories() ([]types.Category, error) {
	var err error

	currentUnix := time.Now().Unix()
	if types.Config.LastCategoryRequest == 0 || currentUnix-types.Config.LastCategoryRequest > 86400 || len(types.Config.Categories) == 0 {
		var req *http.Request
		req, err := getRequest("GET", apiURL+"release/categories.json", nil)
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
						err = json.Unmarshal(bytes, &types.Config.Categories)
						if err == nil {
							types.Config.LastCategoryRequest = currentUnix
						}
					}
				}
			}
		}
	}

	return types.Config.Categories, err
}

/*
GetReleaseByExtInfoID returns all releases associated with a given ExtInfo.

	id				ExtInfoID.
	perPage	:= 25	Number of releases per page. Min. 5, max. 100.
	page	:= 1	Page number (1 to N).

http://www.xrel.to/wiki/2822/api-release-ext-info.html
*/
func GetReleaseByExtInfoID(id string, perPage, page int) (types.Releases, error) {
	form := url.Values{}
	form.Add("id", id)
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

	return getReleases(apiURL+"release/ext_info.json", form)
}

/*
AddReleaseProofImage adds the base64 of a proof picture to a given API release id.
More info on proof pictures can be found here: https://www.xrel.to/wiki/6305/Proofs.html

Please read the rules before posting proofs: https://www.xrel.to/wiki/6308/Regeln-Proofs.html

https://www.xrel.to/wiki/6444/api-release-addproof.html
*/
func AddReleaseProofImage(ids []string, imageBase64 string) (types.AddProofResult, error) {
	var proofResult types.AddProofResult

	var form = url.Values{}
	for i := range ids {
		form.Add("id", ids[i])
	}
	form.Add("image", imageBase64)
	request, err := getOAuth2Request("POST", apiURL+"release/addproof.json", strings.NewReader(form.Encode()))
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
					err = json.Unmarshal(bytes, &proofResult)
				}
			}
		}
	}
	return proofResult, err
}

/*
AddReleaseProofImageByPath adds a proof picture to a given API release id by filepath.
More info on proof pictures can be found here: https://www.xrel.to/wiki/6305/Proofs.html

Please read the rules before posting proofs: https://www.xrel.to/wiki/6308/Regeln-Proofs.html

https://www.xrel.to/wiki/6444/api-release-addproof.html
*/
func AddReleaseProofImageByPath(ids []string, fp string) (types.AddProofResult, error) {
	var proofResult types.AddProofResult

	if len(ids) == 0 {
		return proofResult, types.NewError("client", "argument_missing", "ids", "")
	}
	if fp == "" {
		return proofResult, types.NewError("client", "argument_missing", "fp", "")
	}
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return proofResult, types.NewError("client", "file_not_found", fp, "")
	}

	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return proofResult, err
	}

	imageBase64 := base64.StdEncoding.EncodeToString(bytes)
	return AddReleaseProofImage(ids, imageBase64)
}
