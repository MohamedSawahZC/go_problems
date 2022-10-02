package main

import (
	"fmt"
)

func main() {
	var n, h, tp int
	fmt.Scan(&n, &h)
	w := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&tp)
		if tp > h {
			w += 2
		} else {
			w++
		}
	}
	fmt.Println(w)
}
