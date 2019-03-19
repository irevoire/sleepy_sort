package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func swap(n int, val chan int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
	val<-n
}

func sort(l []int) []int {
	var sorted []int
	val := make(chan int)

	for _, n := range(l) {
		go swap(n, val)
	}

	for range(l) {
		sorted = append(sorted, <-val)
	}

	return sorted
}

func main() {
	list := os.Args[2:]

	var list2 []int
	for _, x := range list {
		v, _ := strconv.Atoi(x)
		list2 = append(list2, v)
	}
	fmt.Println(list2)
	fmt.Println(sort(list2))
}

