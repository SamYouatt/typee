package main

import (
	"math"
	"strings"
)

type Test struct {
	text             string
	numWords         int
	completedWords   int
	currentIndex     int
	currentlyInvalid bool
	errorIndices     []int

	// time in second an error occured
	errorTimes []int

	// wpm calculated after each word completion
	runningWpm []int

	complete  bool
	stopwatch Stopwatch
}

func NewTest(testText string) *Test {
	stopwatch := Stopwatch{}
	numWords := len(strings.Fields(testText))

	return &Test{
		text:             testText,
		numWords:         numWords,
		completedWords:   0,
		currentIndex:     0,
		currentlyInvalid: false,
		errorIndices:     []int{},
		errorTimes:       []int{},
		complete:         false,
		stopwatch:        stopwatch,
	}
}

func (test *Test) playBackspace() {
	if test.complete || !test.currentlyInvalid {
		return
	}

	test.currentlyInvalid = false
}

func (test *Test) playSpace() {
	if test.complete {
		return
	}

	if test.text[test.currentIndex] != ' ' {
		test.errorIndices = append(test.errorIndices, test.currentIndex)
		test.currentlyInvalid = true
		return
	}

	test.completedWords++
	test.currentIndex++

	timeTaken := test.stopwatch.ElapsedTime()
	wpm := CalculateWpm(test.completedWords, timeTaken)
	test.runningWpm = append(test.runningWpm, wpm)

	return
}

func (test *Test) playCharacter(char byte) (completed bool) {
	if test.complete {
		return true
	}

	if !test.stopwatch.started {
		test.stopwatch.Start()
	}

	if test.currentlyInvalid {
		// Have to hit backspace to fix error
		return false
	}

	if test.text[test.currentIndex] == char {
		if test.currentIndex == len(test.text)-1 {
			test.complete = true
			return true
		} else {
			test.currentIndex++
		}
	} else {
		currentElapsedSecs := test.stopwatch.ElapsedTime().Seconds()
		errorOccurredAt := int(math.Round(currentElapsedSecs))

		test.errorIndices = append(test.errorIndices, test.currentIndex)
		test.errorTimes = append(test.errorTimes, errorOccurredAt)
		test.currentlyInvalid = true
	}

	return false
}

func (test *Test) PlayInput(input string) (completed bool) {
	switch input {
	case " ":
		test.playSpace()
		return false
	case "backspace":
		test.playBackspace()
		return false
	default:
		return test.playCharacter(input[0])
	}
}
