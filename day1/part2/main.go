package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	if err != nil {

	}

	scan := bufio.NewScanner(input)
	var a []int
	for scan.Scan() {
		line := scan.Text()
		i, err := strconv.Atoi(line)
		if err != nil {

		}
		a = append(a, i)
	}

	b := a[1 : len(a)-1]
	c := a[2:]
	a = a[0 : len(a)-2]

	fmt.Printf("b: %v\n", len(b))
	fmt.Printf("c: %v\n", len(c))
	fmt.Printf("a: %v\n", len(a))
	count := 0
	before := -1
	for i, v := range a {
		current := v + b[i] + c[i]
		if before == -1 {
			before = current
			continue
		}
		if current > before {
			count++
		}
		before = current

	}
	fmt.Println("count:", count)

}
