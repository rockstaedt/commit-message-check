package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

type UpdateConfig struct {
	Version       string
	LatestVersion string
	TagUrl        string
	BinaryBaseUrl string
	DownloadPath  string
}

func Update(config *UpdateConfig) int {
	config.LatestVersion = getLatestTag(config.TagUrl)
	if config.LatestVersion == "" {
		log.Println("Error at retrieving latest version.")
		return 1
	}

	if config.Version == config.LatestVersion {
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
	file, err := os.Create(config.DownloadPath + "/commit-message-check")
	if err != nil {
		return 1
	}

	res, err := http.Get(getBinaryUrl(config))
	if err != nil {
		return 2
	}

	if res.StatusCode != 200 {
		return 3
	}

	_, _ = io.Copy(file, res.Body)

	return 0
}

func getBinaryUrl(config *UpdateConfig) string {
	return fmt.Sprintf(
		"%s/commit-message-check-%s-%s-%s",
		config.BinaryBaseUrl,
		config.LatestVersion,
		runtime.GOOS,
		runtime.GOARCH,
	)
}
