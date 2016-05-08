package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/*
GetComments returns comments for a given API release id or API P2P release id.
Only the id is required.

	isP2P   := false
	perPage := 25 		// min. 5, max 100
	page    := 1

http://www.xrel.to/wiki/6313/api-comments-get.html
*/
func GetComments(id string, isP2P bool, perPage int, page int) (types.Comments, error) {
	var comments types.Comments

	form := url.Values{}

	form.Add("id", id)
	if isP2P {
		form.Add("type", "p2p_rls")
	} else {
		form.Add("type", "release")
	}
	if perPage != 0 {
		if perPage < 5 {
			perPage = 5
		}
		if perPage > 100 {
			perPage = 100
		}
		form.Add("per_page", strconv.Itoa(perPage))
	}
	if page > 0 {
		form.Add("page", strconv.Itoa(page))
	}
	req, err := getRequest("GET", apiURL+"comments/get.json?"+form.Encode(), nil)
	if err == nil {
		var response *http.Response
		client := http.DefaultClient
		response, err = client.Do(req)
		if err == nil {
			defer response.Body.Close()
			err = checkResponse(response)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					err = json.Unmarshal(bytes, &comments)
				}
			}
		}
	}

	return comments, err
}

/*
AddComment adds a comment to a given API release id or API P2P release id.
Requires user OAuth2 authentication.

	id			API release id or API P2P release id.
	isP2P			If the provided id is a P2P release id.
	text		:= ""	The comment. You may use BBCode to format the text.
				Can be empty if both video_rating and audio_rating are set.
	videoRating	:= 0
	audioRating	:= 0	Video and audio rating between 1 (bad) to 10 (good). 0 means no rating.
						You must always rate both or none. You may only vote once, and may not change your vote.

http://www.xrel.to/wiki/6312/api-comments-add.html
*/
func AddComment(id string, isP2P bool, text string, videoRating, audioRating int) (types.Comment, error) {
	var (
		comment types.Comment
		err     error
	)

	if id == "" {
		err = types.NewError("client", "argument_missing", "id", "")
	} else if (videoRating > 0 && audioRating < 1) || (videoRating < 1 && audioRating > 0) ||
		videoRating > 10 || audioRating > 10 {
		err = types.NewError("client", "invalid_argument", "video or audio rating`", "")
	} else if videoRating < 1 && text == "" {
		err = types.NewError("client", "argument_missing", "text or rating", "")
	} else {
		form := url.Values{}
		form.Add("id", id)
		if isP2P {
			form.Add("type", "p2p_rls")
		} else {
			form.Add("type", "release")
		}
		if text != "" {
			form.Add("text", text)
		}
		if videoRating > 0 {
			form.Add("video_rating", strconv.Itoa(videoRating))
			form.Add("audio_rating", strconv.Itoa(audioRating))
		}
		var request *http.Request
		request, err = getOAuth2Request("POST", apiURL+"comments/add.json", strings.NewReader(form.Encode()))
		if err == nil {
			client := http.DefaultClient
			var response *http.Response
			response, err = client.Do(request)
			if err == nil {
				defer response.Body.Close()
				err = checkResponse(response)
				if err == nil {
					var bytes []byte
					bytes, err = ioutil.ReadAll(response.Body)
					if err == nil {
						err = json.Unmarshal(bytes, &comment)
					}
				}
			}
		}
	}

	return comment, err
}
