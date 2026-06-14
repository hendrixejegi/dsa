package utils

import (
	"math/rand"
)

func CreateArr(length int) []int {
	arr := make([]int, length)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
