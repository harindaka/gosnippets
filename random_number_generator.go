package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers() {
	randomNumber := randInt(10, 20)
	fmt.Println(randomNumber)

	minValue := -10
	maxValue := 10

	minOccurances := 0
	maxOccurances := 0
	outOfRangeOccurances := 0
	negativeOccurances := 0
	positiveOccurances := 0

	for i := 0; i < 1000; i++ {
		randomNumber = randInt(minValue, maxValue)
		if randomNumber == minValue {
			minOccurances++
		} else if randomNumber == maxValue {
			maxOccurances++
		} else if randomNumber < minValue || randomNumber > maxValue {
			outOfRangeOccurances++
		} else if randomNumber < 0 {
			negativeOccurances++
		} else {
			positiveOccurances++
		}
	}

	fmt.Println("Min Occurances:", minOccurances)
	fmt.Println("Max Occurances:", maxOccurances)
	fmt.Println("Positive Occurances within Range:", positiveOccurances)
	fmt.Println("Negative Occurances within Range:", negativeOccurances)
	fmt.Println("Out of Range Occurances:", outOfRangeOccurances) //Should always be zero

}

func randInt(min int, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn((max+1)-min) + min
}
