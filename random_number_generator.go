package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers() {
	randomNumber := randInt(10, 20)
	fmt.Println(randomNumber)

	minOccurances := 0
	maxOccurances := 0
	withinRangeOccurances := 0
	outOfRangeOccurances := 0

	for i := 0; i < 1000; i++ {
		randomNumber = randInt(10, 20)
		if randomNumber == 10 {
			minOccurances++
		} else if randomNumber == 20 {
			maxOccurances++
		} else if randomNumber < 10 || randomNumber > 20 {
			outOfRangeOccurances++
		} else {
			withinRangeOccurances++
		}
	}

	fmt.Println("Min Occurances:", minOccurances)
	fmt.Println("Max Occurances:", maxOccurances)
	fmt.Println("Out of Range Occurances:", outOfRangeOccurances)
	fmt.Println("Within Range Occurances:", withinRangeOccurances)

}

func randInt(min int, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn((max+1)-min) + min
}
