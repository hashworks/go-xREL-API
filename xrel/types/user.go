package types

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Secret         string `json:"secret"`
	Locale         string `json:"locale"`
	AvatarURL      string `json:"avatar_url"`
	AvatarThumbURL string `json:"avatar_thumb_url"`
}
