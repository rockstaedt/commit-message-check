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

	var data responseData
	_ = json.NewDecoder(res.Body).Decode(&data)

	if version != data.TagName {
		return 1
	}

	return 0
}
