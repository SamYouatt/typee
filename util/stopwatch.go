package util

import "time"

type Stopwatch struct {
	startTime time.Time
	started   bool
}

func (s *Stopwatch) Start() {
	s.startTime = time.Now()
	s.started = true
}

func (s *Stopwatch) ElapsedTime() time.Duration {
	if !s.started {
		panic("Can't get elapsed time from stopwatch that hasn't started")
	}

	return time.Since(s.startTime)
}

func (s *Stopwatch) IsRunning() bool {
	return s.started
}
