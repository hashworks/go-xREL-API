package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
)

/*
GetUserInfo returns information about the currently active user.
Requires user OAuth authentication.

https://www.xrel.to/wiki/6441/api-user-info.html
*/
func GetUserInfo() (types.User, error) {
	var user types.User

	request, err := getOAuth2Request("GET", apiURL+"user/info.json", nil)
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
					err = json.Unmarshal(bytes, &user)
				}
			}
		}
	}

	return user, err
}
