package xtime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func IntDateOfDay(t time.Time) (int, error) {
	return IntDateOfDayWithFormat(t, "20060102")
}

//本地时区
func DayOfIntDateInLocal(i int) (time.Time, error) {
	return DayOfIntDateInLocationWithFormat(i, "20060102", time.Local)
}

func IntDateOfDayWithFormat(t time.Time, layout string) (int, error) {
	date, err := strconv.Atoi(t.Format(layout))
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return date, nil
}

func DayOfIntDateWithFormat(i int, layout string) (time.Time, error) {
	day, err := time.Parse(layout, fmt.Sprint(i))
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return day, err
}

func DayOfIntDateInLocationWithFormat(i int, layout string, loc *time.Location) (time.Time, error) {
	day, err := time.ParseInLocation(layout, fmt.Sprint(layout), loc)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return day, err
}

func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func BeginingOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	h := t.Hour()
	return time.Date(y, m, d, h, 0, 0, 0, t.Location())
}

func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	h := t.Hour()
	return time.Date(y, m, d, h, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}
