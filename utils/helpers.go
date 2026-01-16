package utils

import (
	"time"

	"github.com/blutspende/bloodlab-common/timezone"
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

func ParseBerlinTimeStringToUTCTime(timeString string) time.Time {
	location, err := timezone.EuropeBerlin.GetLocation()
	if err != nil {
		return time.Time{}
	}
	berlinTime, err := time.ParseInLocation("20060102150405", timeString, location)
	if err != nil {
		return time.Time{}
	}
	return berlinTime.UTC()
}
