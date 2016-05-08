package types

import "time"

type OAuth2Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Expires      time.Time
	RefreshToken string `json:"refresh_token"`
}
