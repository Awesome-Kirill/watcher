package middleware

import "sync/atomic"

type Stat struct {
	min uint64
	max uint64
	url uint64
}

func (s *Stat) IncMin() {
	atomic.AddUint64(&s.min, 1)
}
func (s *Stat) IncMax() {
	atomic.AddUint64(&s.max, 1)
}
func (s *Stat) IncUrl() {
	atomic.AddUint64(&s.url, 1)
}

func (s *Stat) Min() uint64 {
	return s.min
}

func (s *Stat) Max() uint64 {
	return s.max
}

func (s *Stat) Url() uint64 {
	return s.url
}
