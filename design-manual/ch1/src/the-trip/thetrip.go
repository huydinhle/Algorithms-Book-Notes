package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// Tripexp calculate trip expense
// This function just uses float64 number, but because float64
// is represented wrong sometimes in binary, we got weird result
// sometimes
func Tripexp(num int, expense ...float64) float64 {
	// calculat the average
	total := 0.00
	for _, v := range expense {
		total += v
		fmt.Printf("total = %+v\n", total)
	}
	average := total / float64(num)
	fmt.Printf("average = %+v\n", average)
	result := 0.00
	for _, v := range expense {
		if v > average {
			result = result + (v - average)
		}
	}
	return result
}

// Tripexp2 calculate trip expense
// We used the decimal package from github to get the precise
// answer but still can't
func Tripexp2(num int, expense ...float64) decimal.Decimal {
	// calculat the average
	total := decimal.NewFromFloat(0.00)
	for _, v := range expense {
		total = total.Add(decimal.NewFromFloat(v))
		fmt.Printf("total = %+v\n", total)
	}
	average := total.Div(decimal.NewFromFloat(float64(num)))
	fmt.Printf("average = %+v\n", average)
	result := decimal.NewFromFloat(0.00)
	for _, v := range expense {
		if decimal.NewFromFloat(v).Cmp(average) == -1 {
			// result = result + (v - average)
			result = result.Add(decimal.NewFromFloat(v).Sub(average))
		}
	}
	return result
}

// Tripexp3 calculate trip expense using a bunch of int
func Tripexp3(num int, expense ...int) int {
	// calculat the average
	total := 0
	for _, v := range expense {
		total += v
		fmt.Printf("total = %+v\n", total)
	}
	average := total / num
	// this is a fraction of a cent
	remainder := total % num
	fmt.Printf("average = %+v\n", average)
	fmt.Printf("remainder = %+v\n", remainder)
	positive := 0
	count := 0
	for _, v := range expense {
		if v > average {
			positive = positive + v
			fmt.Printf("positive = %+v\n", positive)
			count++
		}
	}
	fmt.Printf("count = %+v\n", count)
	result := positive - (average * count) - ((remainder / 4) * count)

	return result
}
