package stat

import "sync/atomic"

type Stat struct {
	Min uint64
	Max uint64
	Url uint64
}

func (s *Stat) IncMin() {
	atomic.AddUint64(&s.Min, 1)
}
func (s *Stat) IncMax() {
	atomic.AddUint64(&s.Max, 1)
}
func (s *Stat) IncUrl() {
	atomic.AddUint64(&s.Url, 1)
}

func (s *Stat) Stat() Stat {
	return *s
}
