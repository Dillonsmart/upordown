package pkg

import (
	"net/url"
	"strings"
)

func GetHost(urlStr string) (string, error) {
	if !strings.Contains(urlStr, "://") {
		urlStr = "https://" + urlStr
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	hostname := parsedURL.Hostname()
	if hostname == "" {
		return "", err
	}

	// Remove www. prefix if present
	if strings.HasPrefix(hostname, "www.") {
		hostname = strings.TrimPrefix(hostname, "www.")
	}

	return hostname, nil
}
