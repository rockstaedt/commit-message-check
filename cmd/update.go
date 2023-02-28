package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type UpdateConfig struct {
	Version      string
	TagUrl       string
	BinaryUrl    string
	DownloadPath string
}

func Update(config *UpdateConfig) int {
	latestTag := getLatestTag(config.TagUrl)
	if latestTag == "" {
		log.Println("Error at retrieving latest version.")
		return 1
	}

	if config.Version == latestTag {
		log.Println("Current version is latest version.")
		return 0
	}

	return downloadScript(config)
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

func downloadScript(config *UpdateConfig) int {
	_, err := os.Create(config.DownloadPath + "/commit-message-check")
	if err != nil {
		return 1
	}

	return 0
}
