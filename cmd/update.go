package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"runtime"
)

type responseData struct {
	TagName string `json:"tag_name"`
}

func Update(config *model.Config) int {
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

func downloadScript(config *model.Config) int {
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

	log.Printf(
		"[SUCCESS]\t Updated commit-message-check successfully to %s",
		config.LatestVersion,
	)

	return 0
}

func getBinaryUrl(config *model.Config) string {
	return fmt.Sprintf(
		"%s/commit-message-check-%s-%s-%s",
		config.BinaryBaseUrl,
		config.LatestVersion,
		runtime.GOOS,
		runtime.GOARCH,
	)
}
