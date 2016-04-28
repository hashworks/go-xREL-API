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
		} else {
			extra = string(bytes)
		}
	}

	if response.StatusCode == 404 {
		return types.NewError("client", "function_not_found", "", "")
	}

	return types.NewError("client", "parsing_failed", err.Error()+"\n"+extra, "")
}

func generateGetParametersString(parameters map[string]string) string {
	var query string

	for k, v := range parameters {
		if query == "" {
			query = "?"
		} else {
			query += "&"
		}
		query += k + "=" + v
	}

	return query
}

/*
getClient returns an OAuth2 client if authenticated and a normal client otherwise.
*/
func getClient() *http.Client {
	var client *http.Client
	client, err := getOAuth2Client()
	if err != nil {
		client = http.DefaultClient
	}
	return client
}
