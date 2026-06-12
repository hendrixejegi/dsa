package main

import (
	"dsa/utils"
	"fmt"
	"math/rand"
)

// after every iteration the end of the array will always be sorted
// the outer loop tracks the length of the list removing the sorted part after each iteration
// the inner loop handle the sorting logic for the unsorted part of the list
func bubbleSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		swapped := false

		for j := 0; j < i-1; j++ {
			currentElem := arr[j]
			nextElem := arr[j+1]

			if currentElem > nextElem {
				arr[j], arr[j+1] = nextElem, currentElem
				swapped = true
			}
		}

		if !swapped {
			fmt.Println("Sorted array")
			break
		}
	}

}

func main() {
	arr := make([]int, 300)

	for i := range arr {
		arr[i] = rand.Intn(1000)
	}

	utils.Duration(func() {
		bubbleSort(arr)
		fmt.Println(arr)
	})
}
