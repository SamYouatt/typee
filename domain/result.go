package domain

import "time"

type Result struct {
	NumWords  int
	TimeTaken time.Duration
	Wpm       int

	RunningWpm []int
}
