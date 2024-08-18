package util

import (
	"math"
	"time"
)

func CalculateWpm(words int, timeTaken time.Duration) int {
	minutesTaken := timeTaken.Minutes()
	wpm := math.Round(float64(words) / minutesTaken)
	return int(wpm)
}
