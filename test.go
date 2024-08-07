package main

import "strings"

type Test struct {
	text             string
	numWords         int
	completedWords   int
	currentIndex     int
	currentlyInvalid bool
	errorIndices     []int
	complete         bool
}

func NewTest(testText string) *Test {
	numWords := len(strings.Fields(testText))

	return &Test{
		text:             testText,
		numWords:         numWords,
		completedWords:   0,
		currentIndex:     0,
		currentlyInvalid: false,
		errorIndices:     []int{},
		complete:         false,
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

func (test *Test) PlayCharacter(char byte) {
	if test.complete {
		return
	}

	if test.currentlyInvalid {
		// Have to hit backspace to fix error
		return
	}

	if test.text[test.currentIndex] == char {
		if test.currentIndex == len(test.text)-1 {
			test.complete = true
		} else {
			test.currentIndex++
		}
	} else {
		test.errorIndices = append(test.errorIndices, test.currentIndex)
		test.currentlyInvalid = true
	}
}
