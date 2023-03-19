package helpers

import (
	"net"
	"net/url"
	"os"
	"strings"
)

// check if the domain and host are valid
func IsValidURL(link string) bool {
	// Check it's an Absolute URL or absolute path
	uri, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	// Check it's an acceptable scheme
	switch uri.Scheme {
	case "http":
	case "https":
	default:
		return false
	}

	// Check it's a valid domain name
	_, err = net.LookupHost(uri.Host)
	return err == nil
}

// ensure user doesn't shorten our own domain
func RemoveDomainError(link string) bool {
	domain := os.Getenv("DOMAIN")
	if link == domain {
		return false
	}

	fixedURL := strings.Replace(link, "http://", "", 1)
	fixedURL = strings.Replace(fixedURL, "https://", "", 1)
	fixedURL = strings.Replace(fixedURL, "www.", "", 1)
	fixedURL = strings.Split(fixedURL, "/")[0]

	return fixedURL == domain
}

// enforce http
func EnforceHTTP(url string) string {
	if len(url) < 4 {
		return ""
	}

	if url[:4] != "http" {
		return "http://" + url
	}
	return url

}
