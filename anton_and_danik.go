package main

import (
	"fmt"
)

func main() {
	var n int
	var h string
	fmt.Scan(&n, &h)
	a := 0
	d := 0
	for i := 0; i < n; i++ {
		if h[i] == 'A' {
			a++
		} else {
			d++
		}
	}
	if a > d {
		fmt.Println("Anton")
	} else if d > a {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
