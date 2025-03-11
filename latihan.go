package main

import (
	"fmt"
)

func faktorTerdekat(n int) {
	i := 1
	hasil := 1
	prevFact := 1 

	for hasil < n {
		prevFact = hasil
		i++
		hasil *= i
	}

    if n-prevFact <= hasil-n {
		fmt.Println(prevFact)
	} else {
		fmt.Println(hasil)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	faktorTerdekat(n)
}
