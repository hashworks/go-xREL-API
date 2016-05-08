package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"testing"
)

var (
	// Set the following three with -ldflags "-X xrel.oAuth2ClientKey=foo [...]"
	testOAuth2ClientKey    string
	testOAuth2ClientSecret string
	testOAuth2ConfigPath   string
)

func configureOAuth2ForTests(t *testing.T) {
	if types.Config.OAuth2Token.AccessToken == "" {
		if testOAuth2ClientKey != "" {
			ConfigureOAuth2(testOAuth2ClientKey, testOAuth2ClientSecret, "", []string{"viewnfo", "addproof"})
			separator := string(filepath.Separator)
			if testOAuth2ConfigPath == "" {
				usr, err := user.Current()
				if err != nil {
					testOAuth2ConfigPath = "."
				} else {
					testOAuth2ConfigPath = usr.HomeDir + separator + ".config" + separator + "xREL"
				}
				testOAuth2ConfigPath += separator + "config.json"
			}
			configData, err := ioutil.ReadFile(testOAuth2ConfigPath)
			if err == nil {
				err = json.Unmarshal(configData, &types.Config)
			} else {
				t.Log(err.Error())
				t.Skip("Failed to read oAuth2 config file, skipping.")
			}
		} else {
			t.Skip("No oAuth2 client key set, skipping.")
		}
	}
}
