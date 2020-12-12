package timer

import "time"

type Due struct {
	Start    time.Time
	Duration time.Duration
}

func (d Due) Due() time.Duration {
	t0 := d.Start.Add(d.Duration)

	return time.Since(t0)
}

func (d Due) IsOverdue() bool {
	return d.Due() > 0
}
