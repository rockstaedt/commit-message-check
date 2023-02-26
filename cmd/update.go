package cmd

import (
	"encoding/json"
	"log"
	"net/http"
)

func Update(version, url string) int {
	latestTag := getLatestTag(url)
	if latestTag == "" {
		log.Println("Error at retrieving latest version.")
		return 1
	}

	if version != latestTag {
		return 1
	}

	log.Println("Current version is latest version.")
	return 0
}

type responseData struct {
	TagName string `json:"tag_name"`
}

func getLatestTag(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return ""
	}

	if res.StatusCode != 200 {
		return ""
	}

	var data responseData
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return ""
	}

	return data.TagName
}
