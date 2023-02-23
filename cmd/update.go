package cmd

import (
	"encoding/json"
	"net/http"
)

type responseData struct {
	TagName string `json:"tag_name"`
}

func Update(version, url string) int {
	res, _ := http.Get(url)

	if res.StatusCode != 200 {
		return 1
	}

	var data responseData
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return 2
	}

	if version != data.TagName {
		return 3
	}

	return 0
}
