package main

import "strings"

// Represents a single run of a typing test
// `test` is the contents for the test
// `currentIndex` is the index of the character expected to be typed next
// `currentlyInvalid` indicates that the current location was entered incorrectly
// `errorIndices` is an array of indices into the text where errors occurred. There can be mulitple occurrences of any index
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

func (test *Test) PlaySpace() {
	if test.complete {
		return
	}

	if test.text[test.currentIndex] == ' ' {
		test.completedWords++
		test.currentIndex++
	} else {
		test.errorIndices = append(test.errorIndices, test.currentIndex)
	}
}

// Plays the passed in character on the test
func (test *Test) PlayCharacter(char byte) {
	if test.complete {
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
	}
}
