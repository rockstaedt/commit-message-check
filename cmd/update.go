package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"io"
	"net/http"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"runtime"
)

type responseData struct {
	TagName string `json:"tag_name"`
}

func (h *Handler) update() int {
	h.Config.LatestVersion = getLatestTag(h.Config.TagUrl)
	if h.Config.LatestVersion == "" {
		h.notify("Error at retrieving latest version.", "red")
		return 1
	}

	if h.Config.Version == h.Config.LatestVersion {
		h.notify("Current version is latest version.")
		return 0
	}

	statusMsg := downloadScript(&h.Config)
	if statusMsg == "" {
		h.notify("Error while downloading binary.", "red")
		return 1
	}

	h.notify(color.Green + statusMsg)

	return 0
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

func downloadScript(config *model.Config) string {
	file, err := os.Create(config.DownloadPath + "/commit-message-check")
	if err != nil {
		return ""
	}

	res, err := http.Get(getBinaryUrl(config))
	if err != nil {
		return ""
	}

	if res.StatusCode != 200 {
		return ""
	}

	_, _ = io.Copy(file, res.Body)

	return "Updated commit-message-check successfully to " + config.LatestVersion
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
