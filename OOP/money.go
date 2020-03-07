package main

import "fmt"

type money float64

func (m money) price_game()string{
	return fmt.Sprintf("$%.2f", m)
}
