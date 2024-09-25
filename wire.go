// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sunsunskibiz/wire/calculator"
)

func InitializeCalculator() calculator.Calculator {
	wire.Build(calculator.NewCalculator, calculator.NewAdder)

	return calculator.Calculator{}
}
