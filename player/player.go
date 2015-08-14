package player

import (
	"fmt"
	"math/rand"
)

type Player interface {
	GetSymbol() string
	GetInput() int
	ValidInput() error
}
type ComputerInput struct {
	Symbol string
}

func (c *ComputerInput) GetSymbol() string {
	return c.Symbol
}

func (c *ComputerInput) GetInput() int {
	return rand.Intn(9)
}
func (c *ComputerInput) ValidInput() error {
	return nil
}

type UserInput struct {
	input  int
	Symbol string
}

func (u *UserInput) GetSymbol() string {
	return u.Symbol
}
func (u *UserInput) GetInput() int {
	fmt.Println("What is your input?")
	fmt.Scanf("%d\n", &u.input)
	return u.input
}
func (u *UserInput) ValidInput() error {
	if u.input < 0 || u.input > 8 {
		return fmt.Errorf("Input must be between 0 and 9")
	}
	return nil
}
