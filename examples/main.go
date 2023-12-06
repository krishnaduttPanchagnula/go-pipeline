package main

import (
	"errors"
	"fmt"

	"github.com/krishnaduttPanchagnula/go-pipeline"
)

func addTwo(x interface{}) (interface{}, error) {
	switch val := x.(type) {
	case int:
		return val + 2, nil
	case []int:
		// Add 2 to each element in the slice
		for i := range val {
			val[i] += 2
		}
		return val, nil
	default:
		return nil, errors.New("unsupported input type")
	}
}

func multiplyByThree(x interface{}) (interface{}, error) {
	switch val := x.(type) {
	case int:
		return val * 3, nil
	case []int:
		// Multiply each element in the slice by 3
		for i := range val {
			val[i] *= 3
		}
		return val, nil
	default:
		return nil, errors.New("unsupported input type")
	}
}

func main() {
	// Create a new pipeline Instance
	p := pipeline.NewPipeline()

	// Add functions to the pipeline
	p.AddFunction(addTwo)
	p.AddFunction(multiplyByThree)

	// Execute the pipeline with different types of initial values
	initialValue := 5
	finalResult, err := p.Execute(initialValue)
	printResult(initialValue, finalResult, err)

	initialList := []int{1, 2, 3}
	finalList, err := p.Execute(initialList)
	printResult(initialList, finalList, err)
}

func printResult(initial, final interface{}, err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Initial Value:", initial)
		fmt.Println("Final Result:", final)
	}
}
