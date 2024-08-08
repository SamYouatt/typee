package main

import (
	"strings"
)

type Test struct {
	text             string
	numWords         int
	completedWords   int
	currentIndex     int
	currentlyInvalid bool
	errorIndices     []int
	complete         bool
	stopwatch        Stopwatch
}

func NewTest(testText string) *Test {
	stopwatch := Stopwatch{}
	stopwatch.Start()
	numWords := len(strings.Fields(testText))

	return &Test{
		text:             testText,
		numWords:         numWords,
		completedWords:   0,
		currentIndex:     0,
		currentlyInvalid: false,
		errorIndices:     []int{},
		complete:         false,
		stopwatch:        stopwatch,
	}
}

func (test *Test) PlayBackspace() {
	if test.complete || !test.currentlyInvalid {
		return
	}

	test.currentlyInvalid = false
}

func (test *Test) PlaySpace() {
	if test.complete {
		return
	}

	if test.text[test.currentIndex] == ' ' {
		test.completedWords++
		test.currentIndex++
	} else {
		test.errorIndices = append(test.errorIndices, test.currentIndex)
		test.currentlyInvalid = true
	}
}

// Returns true when the test has been completed
func (test *Test) PlayCharacter(char byte) bool {
	if test.complete {
		return false
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
		test.errorIndices = append(test.errorIndices, test.currentIndex)
		test.currentlyInvalid = true
	}

	return false
}
