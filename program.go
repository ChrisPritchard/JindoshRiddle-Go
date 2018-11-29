package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	result := solve()
	for _, v := range result {
		fmt.Printf("%s owns the %s\n", womanNames[v.woman], heirloomNames[v.owns])
	}

	elapsed := time.Now().Sub(start)

	fmt.Printf("\n")
	fmt.Printf("Time taken: %f seconds", elapsed.Seconds())
}
