package main

import (
	"Module2/datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	var sampleCount = 0
	for _, number := range numbers {
		sum += number
		sampleCount++
	}
	fmt.Printf("Average: %0.2f\n", sum/float64(sampleCount))
}
