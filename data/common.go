package data

import (
	"strings"
	"time"
)

func makeSeriesSlug(parts []string) string {
	return strings.Join(parts, ":")
}

func updateIfEarlier(tracker *time.Time, datum time.Time) {
	if datum.Before(*tracker) || tracker.IsZero() {
		*tracker = datum
	}
}

func updateIfLater(tracker *time.Time, datum time.Time) {
	if datum.After(*tracker) || tracker.IsZero() {
		*tracker = datum
	}
}
