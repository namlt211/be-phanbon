package mytime

import (
	constant "green/constants"
	"time"
)

var loc *time.Location

func SetTimezone(timezone string) error {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	loc = location
	return nil
}

func GetTime(t time.Time) time.Time {
	return t.In(loc)
}

func Now() time.Time {
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}

func NowUnix() int64 {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc).Unix()
}

func CreateNewDate(dateString string, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	clientTime, err := time.ParseInLocation(constant.DateTimeFormatDisplay, dateString, location)
	if err != nil {
		return time.Time{}, err
	}

	return clientTime.UTC(), nil
}

func GetTimeNowFormat(timezone string) string {
	location, _ := time.LoadLocation(timezone)

	return time.Now().In(location).Format(constant.DateFormat)
}
