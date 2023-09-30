package main

import "fmt"

type OperationStrategy interface {
	Execute(int, int) int
}
type AddOperation struct{}

func (AddOperation) Execute(a, b int) int {
	return a + b
}

type SubstractOperation struct{}

func (s *SubstractOperation) Execute(a, b int) int {
	return a - b
}

type Calculator struct {
	strategy OperationStrategy
}

func (k *Calculator) SetStrategy(strategy OperationStrategy) {
	k.strategy = strategy
}
func (j *Calculator) Calculate(a, b int) int {
	return j.strategy.Execute(a, b)
}

func main() {
	calculator := new(Calculator)
	calculator.SetStrategy(&AddOperation{})
	result := calculator.Calculate(7, 8)
	fmt.Println(result)

	calculator.SetStrategy(&SubstractOperation{})
	result = calculator.Calculate(7, 8)
	fmt.Println(result)
}
