package models

import (
"fmt"
"time"
)
var TimeUnit time.Duration

type RatingSystem struct {
	values []int
}

func NewRatingSystem() *RatingSystem {
	return &RatingSystem{
		values: make([]int, 0),
	}
}

func (s *RatingSystem) Add(value int) {
	s.values = append(s.values, value)
}

func (s *RatingSystem) Return() float32 {
	var result float32 = 0
	for _, value := range s.values {
		result += float32(value)
	}

	return result / float32(len(s.values))
}

func Calculate(start int64, end int64, max float32) int {
	timeFrame := float32(end - start)
	multiplier := float32(TimeUnit) / float32(time.Second)

	maxWait := max * multiplier

	fmt.Printf("TimeFrame : %f : MaxWait : %f\n", timeFrame, maxWait)

	if timeFrame < maxWait {
		return 5
	}
	if timeFrame < maxWait*1.1 {
		return 4
	}
	if timeFrame < maxWait*1.2 {
		return 3
	}
	if timeFrame < maxWait*1.3 {
		return 2
	}
	if timeFrame < maxWait*1.4 {
		return 1
	}

	return 0
}
