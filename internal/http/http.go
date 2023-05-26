package http

import (
	"fmt"
	"io"
	"net/http"
)

const apiURL = "https://www.toptal.com/developers/gitignore/api/"

// GetGitignoreFiles fetches the gitignore templates from https://toptal.com/developers/gitignore/api/
// repository.
func GetGitignoreFiles(opts string) (string, error) {
	URL := fmt.Sprintf("%s%s", apiURL, opts)
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
