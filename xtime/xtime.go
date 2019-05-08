package xtime

import (
	"strconv"
	"time"
)

func IntDateOfDay(t time.Time) int {
	date, _ := strconv.Atoi(t.Format("20060102"))
	return date
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
