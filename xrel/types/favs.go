package types

import (
	"crypto/sha1"
	"encoding/hex"
)

type FavList struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	IsPublic       bool   `json:"public"`
	DoNotify       bool   `json:"notify"`
	DoAutoRead     bool   `json:"auto_read"`
	IncludesP2P    bool   `json:"include_p2p"`
	Description    string `json:"description"`
	PasswordHash   string `json:"passwort_hash"`
	EntryCount     int    `json:"entry_count"`
	UnreadReleases int    `json:"unread_releases"`
}

/*
	Test if a password for a list is correct.

	https://www.xrel.to/wiki/1754/api-favs-lists.html
*/
func (f *FavList) TestPassword(password string) bool {
	if f.PasswordHash == "" {
		return true
	}

	h := sha1.New()
	h.Write([]byte(password))
	passwordHashString := hex.EncodeToString(h.Sum(nil))

	for i := 0; i < h.Size(); i++ {
		if passwordHashString[i] != f.PasswordHash[i] {
			return false
		}
	}

	return true
}

type ShortFavList struct {
	Id             int    `json:"id"`
	Name           string `json:"string"`
	EntryCount     int    `json:"entry_count"`
	UnreadReleases int    `json:"unread_releases"`
}

type FavListEntryModificationResult struct {
	FavList ShortFavList `json:"fav_list"`
	ExtInfo ShortExtInfo `json:"ext_info"`
}
