package ksort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for i := 0; i < 100; i++{
		arr := make([]int,20)
		for j := 0; j < 20; j++{
			arr[j] = rand.Intn(15)
		}
		fmt.Println(arr)
		QuickSort(arr)
		fmt.Println(arr)
		fmt.Println()
	}
}