package types

import "time"

type P2PReleases struct {
	TotalCount int          `json:"total_count"`
	Pagination Pagination   `json:"pagination"`
	List       []P2PRelease `json:"list"`
}

const (
	P2P_RELEASE_MAIN_LANGUAGE_DE = "german"
	P2P_RELEASE_MAIN_LANGUAGE_EN = "english"
)

type P2PRelease struct {
	ID           string       `json:"id"`
	Dirname      string       `json:"dirname"`
	LinkURL      string       `json:"link_href"`
	MainLanguage string       `json:"main_lang"`
	PubTimeUnix  int64        `json:"pub_time"`
	PostTimeUnix int64        `json:"post_time"`
	SizeInMB     int          `json:"size_mb"`
	Group        Group        `json:"group"`
	NumRatings   int          `json:"num_ratings"`
	VideoRating  float32      `json:"video_rating"`
	AudioRating  float32      `json:"audio_rating"`
	ExtInfo      ShortExtInfo `json:"ext_info"`
	TVSeason     int          `json:"tv_season"`
	TVEpisode    int          `json:"tv_episode"`
	Comments     int          `json:"comments"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type P2PCategory struct {
	ID      string `json:"id"`
	MetaCat string `json:"meta_cat"`
	SubCat  string `json:"sub_cat"`
}

func (p2pRelease *P2PRelease) GetPubTime() time.Time {
	return time.Unix(p2pRelease.PubTimeUnix, 0)
}

func (p2pRelease *P2PRelease) GetPostTime() time.Time {
	return time.Unix(p2pRelease.PostTimeUnix, 0)
}
