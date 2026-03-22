package urlrequest

import (
	"net/url"
)

func resolveRedirectUrl(current *url.URL, location string) string {
	next, err := url.Parse(location)

	if err != nil {
		return location
	}

	if next.Host == "" {
		next.Host = current.Host
		next.Scheme = current.Scheme
	}

	return next.String()

}
