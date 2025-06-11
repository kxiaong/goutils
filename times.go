package util

import (
	"time"

	"github.com/pkg/errors"
)

const (
	DateLayout   = "2006-01-02"
	HourLayout   = "2006-01-02 15"
	MinuteLayout = "2006-01-02 15:04"
	SecondLayout = "2006-01-02 15:04:05"
	Y4M2D2       = "20060102"
	Y4M2D2H2m2   = "200601021504"
	CNY4M2D2     = "2006年01月02日"
	CNY2M2D2H2   = "2006年01月02日15点"
)

func GetStartOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetEndOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

func GetStartOfQuarter(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute()/15*15, 0, 0, d.Location())
}

func GetStartOfTomorrow(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day()+1, 0, 0, 0, 0, d.Location())
}

func GetStartOfYesterday(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day()-1, 0, 0, 0, 0, d.Location())
}

func GetStartOfHour(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), 0, 0, 0, d.Location())
}

func GetStartOfLastHour(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour()-1, 0, 0, 0, d.Location())
}

func ParseDateStr(date string, layout string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.Local
	}
	t, err := time.ParseInLocation(layout, date, loc)
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "ParseDateStr, date: %v, layout: %v", date, layout)
	}
	return t, nil
}

func FormatDate(t time.Time, layout string) string {
	return t.Format(layout)
}
