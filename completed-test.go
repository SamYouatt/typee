package main

import "time"

type CompletedTest struct {
	numWords  int
	timeTaken time.Duration
	wpm       int
}

func (t *Test) CompleteTest() *CompletedTest {
	timeTaken := t.stopwatch.ElapsedTime()
	wpm := CalculateWpm(t.completedWords, timeTaken)

	return &CompletedTest{
		numWords:  t.numWords,
		timeTaken: timeTaken,
		wpm:       wpm,
	}
}
