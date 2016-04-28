// Contains structs and constants for the xREL package, reflecting the xREL API JSON returns.
package types

const (
	TYPE_MOVIE    = "movie"
	TYPE_TV       = "tv"
	TYPE_GAME     = "game"
	TYPE_CONSOLE  = "console"
	TYPE_SOFTWARE = "software"
	TYPE_XXX      = "xxx"

	// 2006-01-02 15:04:05.999999999 -0700 MST
	TIME_FORMAT_COMMENT = "02. Jan 2006, 03:04 pm"
	TIME_FORMAT_RELEASE = "02.01.2006 03:04 pm"
)

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_pages"`
}
