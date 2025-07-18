package utils

import (
	"github.com/blutspende/bloodlab-common/timezone"
	"time"
)

func FormatTimeStringToBerlinTime(timeString, format string) time.Time {
	location, err := timezone.EuropeBerlin.GetLocation()
	if err != nil {
		return time.Time{}
	}
	result, err := time.ParseInLocation(format, timeString, location)
	if err != nil {
		return time.Time{}
	}
	return result
}
