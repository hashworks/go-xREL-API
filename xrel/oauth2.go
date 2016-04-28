package xrel

import (
	"github.com/hashworks/go-xREL-API/xrel/types"
	"golang.org/x/oauth2"
	"net/http"
)

var oauth2Config = &oauth2.Config{
	Endpoint: oauth2.Endpoint{
		AuthURL:  apiURL + "oauth2/auth",
		TokenURL: apiURL + "oauth2/token",
	},
}

/*
SetOAuthConsumerKeyAndSecret sets the OAuth consumer key and secret you received from xREL.
If you set no redirectURL xREL will display the code to the user.
Get them here: http://www.xrel.to/api-apps.html
*/
func ConfigureOAuth2(clientKey, clientSecret string, redirectURL string, scopes []string) {
	oauth2Config.ClientID = clientKey
	oauth2Config.ClientSecret = clientSecret
	oauth2Config.Scopes = scopes
	if redirectURL == "" {
		oauth2Config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
	}
}

/*
GetOAuthRequestURL returns an URL where the user can login and get a verification code from.
*/
func GetOAuth2RequestURL() string {
	return oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOnline)
}

/*
Initiate the verification code exchange. On success, xREL will return a token we are gonna save in the Config variable.
*/
func InitiateOAuth2CodeExchange(code string) error {
	token, err := oauth2Config.Exchange(oauth2.NoContext, code)
	if err == nil {
		types.Config.OAuth2Token = token
	}
	return err
}

func getOAuth2Client() (*http.Client, error) {
	var client *http.Client
	var err error
	if oauth2Config.ClientID != "" && oauth2Config.ClientSecret != "" && types.Config.OAuth2Token != nil {
		client = oauth2Config.Client(oauth2.NoContext, types.Config.OAuth2Token)
	} else {
		err = types.NewError("client", "not_authenticated", "", "")
	}
	return client, err
}
