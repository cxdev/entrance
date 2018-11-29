package testsupport

import (
	"time"
)

func parseToTimeWithoutError(layout, dateString string) *time.Time {
	timeObj, err := time.Parse(layout, dateString)
	if err != nil {
		return nil
	}
	return &timeObj
}
