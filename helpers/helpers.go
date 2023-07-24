package helpers

import (
	"os"
	"strings"
)

func EnforceHTTPS(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("APP_DOMAIN") {
		return false
	}

	// Remove https://, http:// and www. from the URL
	newURL := strings.Replace(url, "https://", "", 1) 
	newURL = strings.Replace(newURL, "http://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1) 
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("APP_DOMAIN") {
		return false
	}

	return true

}
