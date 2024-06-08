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

func merge_2(a, b []int, channel chan []int) {
	result := []int{}
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}
	result = append(result, a...)
	result = append(result, b...)
	channel <- result
}

func merge_4(a, b, c, d []int) []int {
	channel := make(chan []int, 2)

	go merge_2(a, b, channel)
	go merge_2(c, d, channel)

	ab := <-channel
	cd := <-channel

	merge_2(ab, cd, channel)

	return <-channel
}

func sortWait(a []int, wg *sync.WaitGroup) {
	sort.Ints(a)
	fmt.Printf("Sorted subarray: %v\n", a)
	wg.Done()
}

func main() {
	var integers []int
	fmt.Println("Please insert at least 4 integeres space separated:")
	scanner := bufio.NewScanner(os.Stdin)
	for !scanner.Scan() {
		fmt.Println("Please insert at least 4 integeres space separated:")
		scanner = bufio.NewScanner(os.Stdin)
	}
	str_integers := strings.Split(scanner.Text(), " ")
	for _, str_value := range str_integers {
		value, err := strconv.Atoi(str_value)
		if err != nil {
			fmt.Printf("The following error occurred: %v\n", err)
			os.Exit(1)
		}
		integers = append(integers, value)
	}

	n := len(integers)
	if n < 4 {
		sort.Ints(integers)
		fmt.Printf("%v\n", integers)
		os.Exit(0)
	}

	a := integers[:n/4]
	b := integers[n/4 : n/2]
	c := integers[n/2 : 3*n/4]
	d := integers[3*n/4:]

	var wg sync.WaitGroup
	wg.Add(4)

	go sortWait(a, &wg)
	go sortWait(b, &wg)
	go sortWait(c, &wg)
	go sortWait(d, &wg)

	wg.Wait()

	fmt.Println(merge_4(a, b, c, d))
}
