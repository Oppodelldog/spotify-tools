package timer

import "time"

type Timer struct {
	Hours   int
	Minutes int
	IsSet   bool
	SetAt   time.Time
}

func (t Timer) AsDue() Due {
	return Due{
		Start:    t.SetAt,
		Duration: time.Hour*time.Duration(t.Hours) + time.Minute*time.Duration(t.Minutes),
	}
}
