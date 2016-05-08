package xrel

import (
	"io/ioutil"
	"net/http"
)

/*
GetNFOByID returns PNG image data of a NFO file for a given release id.
Requires user or application OAuth2 authentication.

	id		The release id as obtained trough over methods.
	isP2P	Is the release a P2P release?

Please cache the file on your device. You may not modify the image in any way or hide the footer.

https://www.xrel.to/wiki/6438/api-nfo-release.html
https://www.xrel.to/wiki/6437/api-nfo-p2p-rls.html
*/
func GetNFOByID(id string, isP2P bool) ([]byte, error) {
	var (
		base64Image []byte
		err         error
	)

	requestURL := apiURL + "nfo/release.json?id="
	if isP2P {
		requestURL = apiURL + "nfo/p2p_rls.json?id="
	}
	requestURL = requestURL + id
	client := http.DefaultClient
	var request *http.Request
	request, err = getOAuth2Request("GET", requestURL, nil)
	if err == nil {
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			err = checkResponse(response)
			if err == nil {
				base64Image, err = ioutil.ReadAll(response.Body)
				if err == nil {
					return base64Image, err
				}
			}
		}
	}

	return base64Image, err
}
