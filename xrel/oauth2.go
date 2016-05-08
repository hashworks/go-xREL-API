package xrel

import (
	"encoding/json"
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	oAuth2ClientKey    string
	oAuth2ClientSecret string
	oAuth2RedirectURL  string
	oAuth2Scopes       []string
)

/*
ConfigureOAuth2 will set needed oAuth2 variables.
If you set no redirectURL xREL will display the code to the user.
Get your client key & secret here: http://www.xrel.to/api-apps.html
*/
func ConfigureOAuth2(clientKey, clientSecret, redirectURL string, scopes []string) {
	if redirectURL == "" {
		redirectURL = "urn:ietf:wg:oauth:2.0:oob"
	}
	oAuth2ClientKey = clientKey
	oAuth2ClientSecret = clientSecret
	oAuth2RedirectURL = redirectURL
	oAuth2Scopes = scopes
}

/*
GetOAuth2AuthURL returns an URL where the user can login and get a verification code from.
State can be any string. You may set this value to any value, and it will be returned after the authentication. It might also be useful to prevent CSRF attacks.
Use this after ConfigureOAuth2.
*/
func GetOAuth2AuthURL(state string) string {
	return apiURL + "oauth2/auth?response_type=code&client_id=" + url.QueryEscape(oAuth2ClientKey) + "&redirect_uri=" + url.QueryEscape(oAuth2RedirectURL) + "&state=" + url.QueryEscape(state) + "&scope=" + url.QueryEscape(strings.Join(oAuth2Scopes, " "))
}

/*
PerformOAuth2UserAuthentication will perform an user authentication with the code you received from GetOAuth2URL.
*/
func PerformOAuth2UserAuthentication(code string) error {
	client := http.DefaultClient
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("client_id", oAuth2ClientKey)
	values.Add("client_secret", oAuth2ClientSecret)
	values.Add("code", code)
	values.Add("redirect_uri", oAuth2RedirectURL)
	values.Add("scope", strings.Join(oAuth2Scopes, ""))
	response, err := client.PostForm(apiURL+"oauth2/token", values)
	if err == nil {
		err = handleOAuth2TokenResponse(response)
	}
	return err
}

/*
PerformOAuth2UserAuthentication will perform an application authentication.
Use this after ConfigureOAuth2.
*/
func PerformOAuth2ApplicationAuthentication() error {
	client := http.DefaultClient
	values := url.Values{}
	values.Add("grant_type", "client_credentials")
	values.Add("client_id", oAuth2ClientKey)
	values.Add("client_secret", oAuth2ClientSecret)
	values.Add("redirect_uri", oAuth2RedirectURL)
	values.Add("scope", strings.Join(oAuth2Scopes, ""))
	response, err := client.PostForm(apiURL+"oauth2/token", values)
	if err == nil {
		err = handleOAuth2TokenResponse(response)
	}
	return err
}

func refreshAccessToken() error {
	client := http.DefaultClient
	values := url.Values{}
	values.Add("grant_type", "refresh_token")
	values.Add("client_id", oAuth2ClientKey)
	values.Add("client_secret", oAuth2ClientSecret)
	values.Add("refresh_token", types.Config.OAuth2Token.RefreshToken)
	response, err := client.PostForm(apiURL+"oauth2/token", values)
	if err == nil {
		err = handleOAuth2TokenResponse(response)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func handleOAuth2TokenResponse(response *http.Response) error {
	unixNow := time.Now().Unix()
	defer response.Body.Close()
	err := checkResponse(response)
	if err == nil {
		var bytes []byte
		bytes, err = ioutil.ReadAll(response.Body)
		if err == nil {
			err = json.Unmarshal(bytes, &types.Config.OAuth2Token)
			if err == nil {
				types.Config.OAuth2Token.Expires = time.Unix(unixNow+types.Config.OAuth2Token.ExpiresIn, 0)
			}
		}
	}
	return err
}

/*
Reader example:
	form := url.Values{}
	[...]
	strings.NewReader(form.Encode())
Or:
	strings.NewReader("z=post&both=y&prio=2&empty=")
*/
func getOAuth2Request(method, url string, body io.Reader) (*http.Request, error) {
	var req *http.Request
	var err error
	if types.Config.OAuth2Token.AccessToken != "" {
		if time.Now().After(types.Config.OAuth2Token.Expires) {
			err = refreshAccessToken()
		}
		if err == nil {
			req, err = http.NewRequest(method, url, body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", types.Config.OAuth2Token.TokenType+" "+types.Config.OAuth2Token.AccessToken)
		}
	} else {
		err = types.NewError("client", "not_authenticated", "", "")
	}
	return req, err
}
