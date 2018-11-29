package main

import (
	"fmt"
)

func main() {
	//all := allPossibilities()

	var testa interface{}
	testa = madamNatsiou
	testb := testa.(hometown)

	fmt.Printf("test: %T %v", testb, testb)
}
