package timezone

import (
	"time"
)

func GetLocation(timezone TimeZone) (*time.Location, error) {
	location, err := time.LoadLocation(string(timezone))
	if err != nil {
		return nil, err
	}
	return location, nil
}
