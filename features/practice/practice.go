package practice

import (
	"math"
	"math/rand"
	"strings"

	"github.com/SamYouatt/typee/corpus"
	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/util"
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
	stopwatch util.Stopwatch
}

func NewTest() *Test {
	stopwatch := util.Stopwatch{}

	testText := generateTest(15)

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

func generateTest(length int) string {
	wordsList := corpus.Enghlish1k()

	testText := wordsList[rand.Intn(len(wordsList))]
	for i := 0; i < length; i++ {
		index := rand.Intn(len(wordsList))
		word := wordsList[index]
		testText = testText + " " + word
	}

	return testText
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

	test.runningWpm = append(test.runningWpm, test.CalculateWpm())

	return
}

func (test *Test) playCharacter(char byte) (completed bool) {
	if test.complete {
		return true
	}

	if !test.stopwatch.IsRunning() {
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

func (test *Test) CalculateWpm() int {
	timeTaken := test.stopwatch.ElapsedTime()
	wpm := util.CalculateWpm(test.completedWords, timeTaken)

	return wpm
}

func (t *Test) CompleteTest() *domain.Result {
	timeTaken := t.stopwatch.ElapsedTime()
	wpm := t.CalculateWpm()

	return &domain.Result{
		NumWords:   t.numWords,
		TimeTaken:  timeTaken,
		Wpm:        wpm,
		RunningWpm: t.runningWpm,
	}
}
