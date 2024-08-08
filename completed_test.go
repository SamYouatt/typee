package main

import "time"

type CompletedTest struct {
	numWords int
    timeTaken time.Duration
}

func (t *Test) CompleteTest() CompletedTest {
	timeTaken := t.stopwatch.ElapsedTime()

	return CompletedTest{
		numWords: t.numWords,
        timeTaken: timeTaken,
	}
}
