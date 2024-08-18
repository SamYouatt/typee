package main

import (
	"slices"
	"testing"
)

func TestStopwatch_OnlyStartsOnFirstInput(t *testing.T) {
	test := NewTest("Test")

	if test.stopwatch.started {
		t.Errorf("Expected stopwatch to be stopped until first input")
	}

	test.playCharacter('T')

	if !test.stopwatch.started {
		t.Errorf("Expected stopwatch to be running after first input")
	}
}

func TestPlayCharacter_WithCorrectCharacter(t *testing.T) {
	test := NewTest("Test text")

	test.playCharacter('T')

	if test.currentIndex != 1 {
		t.Errorf("Expected correct character to advance index")
	}
}

func TestPlayCharacter_WithIncorrectCharacter(t *testing.T) {
	test := NewTest("Test text")

	test.playCharacter('x')

	if test.currentIndex != 0 {
		t.Errorf("Should not have advanced current index for incorrect letter")
	}

	if !slices.Contains(test.errorIndices, 0) {
		t.Errorf("Expected index to be added to error indices")
	}

	if !test.currentlyInvalid {
		t.Errorf("Expected to be marked as currently invalid after bad character entry")
	}

	if len(test.errorTimes) != 1 {
		t.Errorf("Expected occurence time of error to be added to error times")
	}
}

func TestPlaySpace_WhenCorrectAdvancesWordCount(t *testing.T) {
	test := NewTest("Test text")
	test.playCharacter('T')
	test.playCharacter('e')
	test.playCharacter('s')
	test.playCharacter('t')

	test.playSpace()

	if test.currentIndex != 5 {
		t.Errorf("Expected correct space to advance index")
	}

	if test.completedWords != 1 {
		t.Errorf("Expected the number of completed words to increase")
	}
}

func TestPlaySpace_WhenCorrect_TracksRunningWpm(t *testing.T) {
	test := NewTest("Test text")
	test.playCharacter('T')
	test.playCharacter('e')
	test.playCharacter('s')
	test.playCharacter('t')

	test.playSpace()

	if len(test.runningWpm) != 1 {
		t.Errorf("Expected running wpm to be calculated when word was completed")
	}
}

func TestPlaySpace_WhenIncorrect(t *testing.T) {
	test := NewTest("Test text")
	test.currentIndex = 3

	test.playSpace()

	if test.currentIndex != 3 {
		t.Errorf("Should not have increased the current index or incorrect entry")
	}

	if test.completedWords != 0 {
		t.Errorf("Should not have marked word as complete when incorrect")
	}

	if !slices.Contains(test.errorIndices, 3) {
		t.Errorf("Expected incorrect index to be added to array")
	}

	if !test.currentlyInvalid {
		t.Errorf("Expected to be currently invalid after bad space entry")
	}
}

func TestPlayBackspace_WhenYetToTypeLetter(t *testing.T) {
	test := NewTest("Test text")
	test.currentIndex = 2

	test.playBackspace()

	if test.currentIndex != 2 {
		t.Errorf("Should not be able to go back when pending letter")
	}
}

func TestPlayBackspace_WhenIncorrectLetterEntered(t *testing.T) {
	test := NewTest("Test text")
	test.currentIndex = 2

	test.playCharacter('x')

	test.playBackspace()

	if test.currentIndex != 2 {
		t.Errorf("Should not be able to go back when pending letter")
	}
}

func TestPlayCharacter_OnLastCharacter_MarksTestAsComplete(t *testing.T) {
	test := NewTest("Test")
	test.currentIndex = 3

	test.playCharacter('t')

	if !test.complete {
		t.Errorf("Should have been marked as complete after last character correct")
	}
}

func TestPlayCharacter_OnNonLastCharacter_ReturnsFalse(t *testing.T) {
	test := NewTest("Test")
	test.currentIndex = 2

	if test.playCharacter('s') {
		t.Fatalf("Expected to return false when test not complete")
	}
}

func TestPlayCharacter_OnLastCharacter_ReturnsTrue(t *testing.T) {
	test := NewTest("Test")
	test.currentIndex = 3

	if !test.playCharacter('t') {
		t.Fatalf("Expected completed test to return true")
	}
}
