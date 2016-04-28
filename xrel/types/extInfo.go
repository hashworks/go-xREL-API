package types

import "time"

type ShortExtInfo struct {
	ID         string   `json:"id"`
	Type       string   `json:"type"`
	Title      string   `json:"title"`
	LinkURL    string   `json:"link_href"`
	Rating     float32  `json:"rating"`
	NumRatings int      `json:"num_ratings"`
	URIs       []string `json:"uris"`
}

type ExtendedExtInfo struct {
	ID           string               `json:"id"`
	Type         string               `json:"type"`
	Title        string               `json:"title"`
	LinkURL      string               `json:"link_href"`
	Genre        string               `json:"genre"`
	AltTitle     string               `json:"alt_title"`
	CoverURL     string               `json:"cover_url"`
	URIs         []string             `json:"uris"`
	Rating       float32              `json:"rating"`
	OwnRating    string               `json:"own_rating"`
	NumRatings   int                  `json:"num_ratings"`
	ReleaseDates []ExtInfoReleaseDate `json:"release_dates"`
	Externals    []ExtInfoExternal    `json:"externals"`
	Releases     []Release            `json:"releases"`
	P2PReleases  []P2PRelease         `json:"p2p_releases"`
}

type ExtInfoReleaseDate struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type ExtInfoExternal struct {
	Source  ExtInfoExternalSource `json:"source"`
	LinkURL string                `json:"link_url"`
	Plot    string                `json:"plot"`
}

type ExtInfoExternalSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	MEDIA_TYPE_IMAGE = "image"
	MEDIA_TYPE_VIDEO = "video"
)

type ExtInfoMediaItem struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	TimeUnix    int64  `json:"time"`
	URLFull     string `json:"url_full"` // IsImage()
	URLThumb    string `json:"url_thumb"`
	YoutubeID   string `json:"youtube_id"` // IsVideo()
	VideoURL    string `json:"video_url"`  // IsVideo()
}

func (extInfoMediaItem *ExtInfoMediaItem) GetTime() time.Time {
	return time.Unix(extInfoMediaItem.TimeUnix, 0)
}
