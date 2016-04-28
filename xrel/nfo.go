package xrel

import (
	"io/ioutil"
	"net/http"
)

/**
GetNFOById returns PNG image data of a NFO file for a given release id.
Requires OAuth2 authentication.

Please cache the file on your device. You may not modify the image in any way or hide the footer.

https://www.xrel.to/wiki/6438/api-nfo-release.html
https://www.xrel.to/wiki/6437/api-nfo-p2p-rls.html
*/
func GetNFOById(id string, isP2P bool) ([]byte, error) {
	var (
		base64Image []byte
		err         error
	)

	client, err := getOAuth2Client()
	if err == nil {
		var response *http.Response
		requestURL := apiURL + "nfo/release.json"
		if isP2P {
			requestURL = apiURL + "nfo/p2p_rls.json"
		}
		response, err = client.Get(requestURL + "?id=" + id)
		defer response.Body.Close()
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
