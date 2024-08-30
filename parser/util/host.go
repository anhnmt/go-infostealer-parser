package util

import (
	"net/url"
	"strings"

	"github.com/samber/lo"
)

// GetHostFromUrl returns host from url
func GetHostFromUrl(rawUrl string) string {
	rawUrl = strings.TrimSpace(rawUrl)
	u, err := url.Parse(rawUrl)
	if err == nil {
		if u.Host == "" {
			return rawUrl
		}

		if strings.EqualFold(u.Scheme, "android") {
			return strings.Join(lo.Reverse(strings.Split(u.Host, ".")), ".")
		}

		return u.Host
	}

	return rawUrl
}
