package cmd

import (
	"encoding/json"
	"net/http"
)

type responseData struct {
	TagName string `json:"tag_name"`
}

func Update(version, url string) int {
	res, err := http.Get(url)
	if err != nil {
		return 1
	}

	if res.StatusCode != 200 {
		return 2
	}

	var data responseData
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return 3
	}

	if version != data.TagName {
		return 4
	}

	return 0
}
