package types

import "time"

type UpcomingTitle struct {
	ID          string                    `json:"id"`
	Type        string                    `json:"type"`
	Title       string                    `json:"title"`
	LinkHref    string                    `json:"link_href"`
	Genre       string                    `json:"genre"`
	AltTitle    string                    `json:"alt_title"`
	CoverURL    string                    `json:"cover_url"`
	Releases    []UpcomingTitleRelease    `json:"releases"`
	P2PReleases []UpcomingTitleP2PRelease `json:"p2p_releases"`
}

type UpcomingTitleRelease struct {
	ID       string `json:"id"`
	Dirname  string `json:"dirname"`
	LinkURL  string `json:"link_href"`
	TimeUnix int64  `json:"time"`
	Flags    Flags  `json:"flags"`
}

func (upcomingTitleRelease *UpcomingTitleRelease) GetTime() time.Time {
	return time.Unix(upcomingTitleRelease.TimeUnix, 0)
}

type UpcomingTitleP2PRelease struct {
	ID          string `json:"id"`
	Dirname     string `json:"dirname"`
	LinkURL     string `json:"link_href"`
	PubTimeUnix int64  `json:"pub_time"`
}

func (upcomingTitleP2PRelease *UpcomingTitleP2PRelease) GetPubTime() time.Time {
	return time.Unix(upcomingTitleP2PRelease.PubTimeUnix, 0)
}
