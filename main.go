package main

import (
	"fmt"

)

func main() {
	calc := InitializeCalculator()

	fmt.Println(calc.Add(1,2))
}