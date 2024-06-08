package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a series of integers separated by space:")

	scanner.Scan()
	input := scanner.Text()

	inputArray := StringToIntArray(input)
	arrayLen := len(inputArray)
	subLen := arrayLen / 4

	subArrays := make([][]int, 4)
	for i := range subArrays {
		start := i * subLen
		end := start + subLen
		if i == 3 {
			end = arrayLen
		}
		subArrays[i] = inputArray[start:end]
		fmt.Printf("Subarray %d: %v\n", i+1, subArrays[i])
	}

	var wg sync.WaitGroup
	for i := range subArrays {
		wg.Add(1)
		go func(i int) {
			sort.Ints(subArrays[i])
			fmt.Printf("Sorted subarray %d: %v\n", i+1, subArrays[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	sortedArray := MergeSortedArrays(subArrays)
	fmt.Println("Sorted array:", sortedArray)
}

// MergeSortedArrays takes an array of sorted subarrays and merges them into a single sorted array
func MergeSortedArrays(subArrays [][]int) []int {
	var sortedArray []int
	for _, arr := range subArrays {
		sortedArray = Merge(sortedArray, arr)
	}
	return sortedArray
}

// Merge is a helper function that merges two sorted arrays
func Merge(a, b []int) (c []int) {
	c = make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}

	for i < len(a) {
		c[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		c[k] = b[j]
		j++
		k++
	}

	return
}

// StringToIntArray takes a string of numbers separated by spaces and returns a slice of integers
func StringToIntArray(input string) []int {
	strs := strings.Fields(input)
	var ints []int
	for _, s := range strs {
		var x int
		fmt.Sscanf(s, "%d", &x)
		ints = append(ints, x)
	}
	return ints
}
