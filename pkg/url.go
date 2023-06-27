package pkg

import (
	"strings"
)

func MakeUrl(rawURL string) string {

	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		return "https://" + rawURL
	}
	return rawURL

}
