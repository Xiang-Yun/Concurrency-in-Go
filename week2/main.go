// Write a program to sort an array of integers.
// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortPart(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting subarray", arr)
	sort.Ints(arr)
}

func readArr(arr *[]int) {
	fmt.Println("Enter integers separated by space:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	elements := strings.Split(input, " ")

	for _, element := range elements {
		num, err := strconv.Atoi(element)
		if err == nil {
			*arr = append(*arr, num)
		}
	}

	fmt.Println("The array is:", *arr)
}

func main() {
	var arr []int
	readArr(&arr)
	n := len(arr)

	partSize := n / 4
	var wg sync.WaitGroup

	wg.Add(4)
	go sortPart(arr[:partSize], &wg)
	go sortPart(arr[partSize:2*partSize], &wg)
	go sortPart(arr[2*partSize:3*partSize], &wg)
	go sortPart(arr[3*partSize:], &wg)
	wg.Wait()

	sortedArr := append(arr[:partSize], arr[partSize:2*partSize]...)
	sortedArr = append(sortedArr, arr[2*partSize:3*partSize]...)
	sortedArr = append(sortedArr, arr[3*partSize:]...)

	sort.Ints(sortedArr)

	fmt.Println("Sorted array:", sortedArr)
}
