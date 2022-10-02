package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	s := ""

	fmt.Scanln(&n)
	fmt.Scanln(&s)

	a := strings.Count(s, "A")
	d := strings.Count(s, "D")

	switch {
	case a > d:
		fmt.Println("Anton")
	case a < d:
		fmt.Println("Danik")
	case a == d:
		fmt.Println("Friendship")

	}
}
