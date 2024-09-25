package main

import (
	"fmt"

	"github.com/sunsunskibiz/wire/calculator"
)

func main() {
	calc := InitializeCalculator()

	fmt.Println(calc.Add(1,2))
}

func InitializeCalculator() calculator.Calculator {
	adder := calculator.NewAdder()
	calc := calculator.NewCalculator(adder)
	return calc
}