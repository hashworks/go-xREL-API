package xrel

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
)

/*
GetUpcomingTitles returns a list upcoming movies and their releases.
The `country` parameter can be either 'de' (default) for upcoming movies in germany or `us` for upcoming movies in the US/international.

http://www.xrel.to/wiki/1827/api-calendar-upcoming.html
*/
func GetUpcomingTitles(country string) ([]types.UpcomingTitle, error) {
	var upcomingTitles []types.UpcomingTitle

	var requestURL = apiURL + "calendar/upcoming.json"
	if country != "" {
		requestURL = requestURL + "?country=" + country
	}

	client := getClient()
	response, err := client.Get(requestURL)
	defer response.Body.Close()
	if err == nil {
		err = checkResponse(response)
		if err == nil {
			bytes, err := ioutil.ReadAll(response.Body)
			if err == nil {
				err = json.Unmarshal(bytes, &upcomingTitles)
			}
		}
	}

	return upcomingTitles, err
}
