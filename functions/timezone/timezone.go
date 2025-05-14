package timezone

import (
	"github.com/blutspende/bloodlab-common/enums/timezone"
	"time"
)

func GetLocation(timezone timezone.TimeZone) (*time.Location, error) {
	location, err := time.LoadLocation(string(timezone))
	if err != nil {
		return nil, err
	}
	return location, nil
}
