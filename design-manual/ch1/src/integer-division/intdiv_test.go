package main

import "testing"

func TestIntdivPositive(t *testing.T) {
	divisor := 30
	divident := 5
	quotent, err := Intdiv(divisor, divident)
	if quotent != 6 || err != nil {
		t.Error("Expect 6, got ", quotent)
	}
}

func TestIntdivOneNegative1(t *testing.T) {
	divisor := -30
	divident := 5
	quotent, err := Intdiv(divisor, divident)
	if quotent != -6 || err != nil {
		t.Error("Expect -6, got ", quotent)
	}
}

func TestIntdivTwoNegative(t *testing.T) {
	divisor := -30
	divident := -5
	quotent, err := Intdiv(divisor, divident)
	if quotent != 6 || err != nil {
		t.Error("Expect 6, got ", quotent)
	}
}

func TestIntdivOneNegative2(t *testing.T) {
	divisor := 30
	divident := -5
	quotent, err := Intdiv(divisor, divident)
	if quotent != -6 || err != nil {
		t.Error("Expect -6, got ", quotent)
	}
}

func TestIntdivWithRemainder(t *testing.T) {
	divisor := 32
	divident := -5
	quotent, err := Intdiv(divisor, divident)
	if quotent != -6 || err != nil {
		t.Error("Expect -6, got ", quotent)
	}
}
