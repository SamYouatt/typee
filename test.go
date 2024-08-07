package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

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

func (test *Test) PlayCharacter(char byte) tea.Cmd {
	if test.complete {
		return nil
	}

	if test.currentlyInvalid {
		// Have to hit backspace to fix error
		return nil
	}

	if test.text[test.currentIndex] == char {
		if test.currentIndex == len(test.text)-1 {
			test.complete = true
			return func() tea.Msg { return TestCompleteMsg{} }
		} else {
			test.currentIndex++
		}
	} else {
		test.errorIndices = append(test.errorIndices, test.currentIndex)
		test.currentlyInvalid = true
	}

	return nil
}
