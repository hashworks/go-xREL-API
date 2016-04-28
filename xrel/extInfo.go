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
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "ext_info/info.json?id=" + id)
		defer response.Body.Close()
		if err == nil {
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
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "ext_info/media.json?id=" + id)
		defer response.Body.Close()
		if err == nil {
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

	return extInfoMediaItemsStruct, err
}

/*
RateExtInfo rates an ExtInfo.
Requires OAuth authentication.

	id		Ext Info ID.
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
		var client *http.Client
		client, err = getOAuth2Client()
		if err == nil {
			var parameters = url.Values{}
			parameters.Add("id", id)
			parameters.Add("rating", strconv.Itoa(rating))
			var response *http.Response
			response, err = client.PostForm(apiURL+"ext_info/rate.json", parameters)
			defer response.Body.Close()
			if err == nil {
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
