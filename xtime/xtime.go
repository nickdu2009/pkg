package xtime

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func IntDateOfDay(t time.Time) (int, error) {
	return IntDateOfDayWithFormat(t, "20060102")
}

func DayOfIntDate(i int) (time.Time, error) {
	return DayOfIntDateWithFormat(i, "20060102")
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
