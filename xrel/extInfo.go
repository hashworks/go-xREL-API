package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/*
GetExtInfo returns information about an ExtInfo.

http://www.xrel.to/wiki/2725/api-ext-info-info.html
*/
func GetExtInfo(id string) (types.ExtendedExtInfo, error) {
	var (
		extInfoStruct types.ExtendedExtInfo
		err           error
	)

	if id == "" {
		err = types.NewError("client", "argument_missing", "id", "")
	} else {
		client := http.DefaultClient
		var request *http.Request
		request, err = getRequest("GET", apiURL+"ext_info/info.json?id="+id, nil)
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
						err = json.Unmarshal(bytes, &extInfoStruct)
					}
				}
			}
		}
	}

	return extInfoStruct, err
}

/*
GetExtInfoMedia returns media associated with an Ext Info.

http://www.xrel.to/wiki/6314/api-ext-info-media.html
*/
func GetExtInfoMedia(id string) ([]types.ExtInfoMediaItem, error) {
	var (
		extInfoMediaItemsStruct []types.ExtInfoMediaItem
		err                     error
	)

	if id == "" {
		err = types.NewError("client", "argument_missing", "id", "")
	} else {
		client := http.DefaultClient
		var request *http.Request
		request, err = getRequest("GET", apiURL+"ext_info/media.json?id="+id, nil)
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
						err = json.Unmarshal(bytes, &extInfoMediaItemsStruct)
					}
				}
			}
		}
	}

	return extInfoMediaItemsStruct, err
}

/*
RateExtInfo rates an ExtInfo.
Requires user OAuth authentication.

	id	Ext Info ID.
	rating	Rating between 1 (bad) to 10 (good). You may only vote once, and may not change your vote.

http://www.xrel.to/wiki/6315/api-ext-info-rate.html
*/
func RateExtInfo(id string, rating int) (types.ExtendedExtInfo, error) {
	var (
		extInfoStruct types.ExtendedExtInfo
		err           error
	)

	if id == "" {
		err = types.NewError("client", "argument_missing", "id", "")
	} else if rating < 1 || rating > 10 {
		err = types.NewError("client", "argument_missing", "rating", "")
	} else {
		client := http.DefaultClient
		var request *http.Request
		form := url.Values{}
		form.Add("id", id)
		form.Add("rating", strconv.Itoa(rating))
		request, err = getOAuth2Request("POST", apiURL+"ext_info/rate.json", strings.NewReader(form.Encode()))
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
						err = json.Unmarshal(bytes, &extInfoStruct)
					}
				}
			}
		}
	}

	return extInfoStruct, err
}
