package data

import (
	"strings"
)

func makeSeriesSlug(parts []string) string {
	return strings.Join(parts, ":")
}

