package main

import (
	"slices"
	"testing"
)

func TestPlayCharacter_WithCorrectCharacter(t *testing.T) {
	test := NewTest("Test text")

    test.PlayCharacter('T')

	if test.currentIndex != 1 {
		t.Errorf("Expected correct character to advance index")
	}
}

func TestPlayCharacter_WithIncorrectCharacter(t *testing.T) {
	test := NewTest("Test text")

    test.PlayCharacter('x')

	if test.currentIndex != 0 {
		t.Errorf("Should not have advanced current index for incorrect letter")
	}

	if !slices.Contains(test.errorIndices, 0) {
		t.Errorf("Expected index to be added to error indices")
	}

	if !test.currentlyInvalid {
		t.Errorf("Expected to be marked as currently invalid after bad character entry")
	}
}

func TestPlaySpace_WhenCorrectAdvancesWordCount(t *testing.T) {
	test := NewTest("Test text")
	test.currentIndex = 4

	test.PlaySpace()

	if test.currentIndex != 5 {
		t.Errorf("Expected correct space to advance index")
	}

	if test.completedWords != 1 {
		t.Errorf("Expected the number of completed words to increase")
	}
}

func TestPlaySpace_WhenIncorrect(t *testing.T) {
	test := NewTest("Test text")
	test.currentIndex = 3

	test.PlaySpace()

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

	test.PlayBackspace()

	if test.currentIndex != 2 {
		t.Errorf("Should not be able to go back when pending letter")
	}
}

func TestPlayBackspace_WhenIncorrectLetterEntered(t *testing.T) {
	test := NewTest("Test text")
	test.currentIndex = 2

	test.PlayCharacter('x')

	test.PlayBackspace()

	if test.currentIndex != 2 {
		t.Errorf("Should not be able to go back when pending letter")
	}
}

func TestPlayCharacter_OnLastCharacter_MarksTestAsComplete(t *testing.T) {
	test := NewTest("Test")
	test.currentIndex = 3

	test.PlayCharacter('t')

	if !test.complete {
		t.Errorf("Should have been marked as complete after last character correct")
	}
}
