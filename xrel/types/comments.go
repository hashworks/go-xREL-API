package types

import "time"

type Comments struct {
	TotalCount int        `json:"total_count"`
	Pagination Pagination `json:"pagination"`
	List       []Comment  `json:"list"`
}

type Comment struct {
	Id       string `json:"id"`
	TimeUnix int64  `json:"time"`
	Author   Author `json:"author"`
	Text     string `json:"text"`
	LinkHref string `json:"link_href"`
	Rating   Rating `json:"rating"`
	Votes    Votes  `json:"votes"`
	Edits    Edits  `json:"edits"`
}

func (c *Comment) GetTime() time.Time {
	return time.Unix(c.TimeUnix, 0)
}

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Rating struct {
	Video int `json:"video"`
	Audio int `json:"audio"`
}

type Votes struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

type Edits struct {
	Count    int   `json:"count"`
	LastUnix int64 `json:"last"`
}

func (edits *Edits) GetLastEditTime() time.Time {
	return time.Unix(edits.LastUnix, 0)
}
