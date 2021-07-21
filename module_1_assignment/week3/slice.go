package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var st []int = make([]int, 3)
	var in string
	fmt.Println("Do enter an integer or (X to exit):")
	for true {
		fmt.Scanln(&in)
		if in == "X" {
			break
		}
		ap, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Input is Wrong")
			continue
		}

		st = append(st, ap)
		sort.Ints(st[:])
		fmt.Println(st)
		fmt.Println("Do enter an integer again or (X to exit):")
	}
}
