package types

import "golang.org/x/oauth2"

/*
Config contains the OAuth2 Token and cached results. Save this somewhere and restore it on every run.

The rate limit information gets set on any request.
*/
var Config = struct {
	OAuth2Token *oauth2.Token

	// 24h caching http://www.xrel.to/wiki/6318/api-release-categories.html
	LastCategoryRequest int64
	Categories          []Category

	// 24h caching http://www.xrel.to/wiki/2996/api-release-filters.html
	LastFilterRequest int64
	Filters           []Filter

	// 24h caching http://www.xrel.to/wiki/3698/api-p2p-categories.html
	LastP2PCategoryRequest int64
	P2PCategories          []P2PCategory

	RateLimitMax       int
	RateLimitRemaining int
	RateLimitResetUnix int64
}{}
