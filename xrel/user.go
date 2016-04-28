package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
)

/*
GetUserInfo returns information about the currently active user.
Requires OAuth authentication.

https://www.xrel.to/wiki/6441/api-user-info.html
*/
func GetUserInfo() (types.User, error) {
	var user types.User

	client, err := getOAuth2Client()
	if err == nil {
		var response *http.Response
		response, err = client.Get(apiURL + "user/info.json")
		defer response.Body.Close()
		if err == nil {
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
