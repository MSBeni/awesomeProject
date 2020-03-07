package main

import "fmt"

type money float64

func (m money) print()string{
	return fmt.Sprintf("$%.2f", m)
}
