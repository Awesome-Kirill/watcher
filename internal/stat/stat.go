package stat

import "sync/atomic"

type Stat struct {
	Min uint64
	Max uint64
	URL uint64
}

func (s *Stat) IncMin() {
	atomic.AddUint64(&s.Min, 1)
}
func (s *Stat) IncMax() {
	atomic.AddUint64(&s.Max, 1)
}
func (s *Stat) IncURL() {
	atomic.AddUint64(&s.URL, 1)
}

func (s *Stat) Stat() Stat {
	return *s
}
