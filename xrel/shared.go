/*
Package xrel contains functions to authorize with and access the complete xREL.to API.

Here is an example on how to use the OAuth2 authentication:

	xrel.ConfigureOAuth2("OAUTH2_CLIENT_KEY", "OAUTH2_CLIENT_SECRET", "", []string{"viewnfo", "addproof"})
	fmt.Println("(1) Go to: " + xrel.GetOAuth2RequestURL())
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Print("(3) Enter that verification code here: ")
	verificationCode := ""
	fmt.Scanln(&verificationCode)
	err := xrel.InitiateOAuth2CodeExchange(verificationCode)
	ok(err)

*/
package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const apiURL = "https://api.xrel.to/v2/"

func checkResponse(response *http.Response) error {
	if response.Header.Get("x-ratelimit-limit") != "" {
		types.Config.RateLimitMax, _ = strconv.Atoi(response.Header.Get("x-ratelimit-limit"))
	}
	if response.Header.Get("x-ratelimit-remaining") != "" {
		types.Config.RateLimitRemaining, _ = strconv.Atoi(response.Header.Get("x-ratelimit-remaining"))
	}
	if response.Header.Get("x-ratelimit-reset") != "" {
		types.Config.RateLimitResetUnix, _ = strconv.ParseInt(response.Header.Get("x-ratelimit-reset"), 10, 32)
	}

	switch response.StatusCode {
	case 200:
		var err error
		return err
	case 429:
		return types.NewError("api", "ratelimit_reached", "", "")
	case 500:
		return types.NewError("api", "internal_error", "", "")
	}

	var extra string

	bytes, err := ioutil.ReadAll(response.Body)
	if err == nil {
		var xErr *types.Error
		err = json.Unmarshal(bytes, &xErr)
		if err == nil {
			return xErr
		}
		extra = string(bytes)
	}
	if response.StatusCode == 404 {
		return types.NewError("client", "function_not_found", "", "")
	}

	return types.NewError("client", "parsing_failed", err.Error()+"\n"+extra, "")
}

/*
getRequest returns an OAuth2 request if authenticated and a normal request otherwise.

Reader example:
	form := url.Values{}
	[...]
	strings.NewReader(form.Encode())
Or:
	strings.NewReader("z=post&both=y&prio=2&empty=")
*/
func getRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := getOAuth2Request(method, url, body)
	if err != nil {
		req, err = http.NewRequest(method, url, body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, err
}
