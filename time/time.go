package time

import (
	"time"
)

// Time defined ext for time
type Time struct {
	time time.Time
}

// New return a new Time
func New(time time.Time) *Time {
	return &Time{
		time: time,
	}
}

// Now defined return timeNow
func Now() *Time {
	return &Time{
		time: time.Now(),
	}
}

// Value defined return time.Time
func (t *Time) Value() time.Time {
	return t.time
}

// StartOfMonth defined StartOfMonth
func (t *Time) StartOfMonth() *Time {
	y, m, _ := t.time.Date()
	t.time = time.Date(y, m, 1, 0, 0, 0, 0, t.time.Location())
	return t
}

// StartOfDay defined StartOfDay
func (t *Time) StartOfDay() *Time {
	y, m, d := t.time.Date()
	t.time = time.Date(y, m, d, 0, 0, 0, 0, t.time.Location())
	return t
}

// StartOf defined startOf Time
func (t *Time) StartOf(tp string) *Time {
	switch true {
	case tp == "month":
		t.time = t.StartOfMonth().Value()
	case tp == "day":
		t.time = t.StartOfDay().Value()
	}
	return t
}

// Range defined range by month
func (t *Time) Range(start time.Time, end time.Time, tp string) []*Time {
	times := []*Time{}
	switch true {
	case tp == "month":
		startUnix := New(start).StartOfMonth()
		endUnix := New(end).StartOfMonth()
		month := startUnix
		for {
			times = append(times, month)
			month = New(month.time.AddDate(0, 1, 0))
			month = month.StartOfMonth()
			if month.time.After(endUnix.Value()) {
				break
			}
		}
	case tp == "day":
		startUnix := New(start).StartOfDay()
		endUnix := New(end).StartOfDay()
		day := startUnix
		for {
			times = append(times, day)
			day = New(day.time.AddDate(0, 0, 1))
			day = day.StartOfDay()
			if day.time.After(endUnix.Value()) {
				break
			}
		}
	}
	return times
}
