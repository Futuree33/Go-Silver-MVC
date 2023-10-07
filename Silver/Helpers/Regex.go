package Helpers

import (
	"regexp"
)

func Regex(route string) *regexp.Regexp {
	pattern := regexp.QuoteMeta(route)
	pattern = regexp.MustCompile(`\\{([a-zA-Z]+)\\}`).ReplaceAllString(pattern, `(?P<$1>[a-zA-Z0-9-\s]+)`)
	pattern = "^" + pattern + "$"

	return regexp.MustCompile(pattern)
}
