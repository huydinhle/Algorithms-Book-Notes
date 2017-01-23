package main

import (
	"errors"
	"fmt"
)

// Intdiv without / or *
func Intdiv(divisor int, divident int) (int, error) {
	sign := 1
	if divisor < 0 {
		divisor *= -1
		sign *= -1
	}
	if divident < 0 {
		divident *= -1
		sign *= -1
	}
	var quotent int
	if divident == 0 {
		return 0, errors.New("crap myself sorry")
	}
	for divisor >= divident {
		fmt.Printf("divisor = %+v\n", divisor)
		fmt.Printf("divident = %+v\n", divident)
		divisor -= divident
		quotent++
	}
	return quotent * sign, nil
}
