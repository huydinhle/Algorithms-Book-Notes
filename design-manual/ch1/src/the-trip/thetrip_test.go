package main

import "testing"

func TestTripexp1(t *testing.T) {
	if v, _ := Tripexp2(3, 10.00, 20.00, 30.00).Float64(); v != 10.00 {
		t.Error("Expecting 10.00 but got ", v)
	}
}

func TestTripexp2(t *testing.T) {
	if v, _ := Tripexp2(4, 15.00, 15.01, 3.00, 3.01).Float64(); v != 11.99 {
		t.Error("Expecting 11.99 but got ", v)
	}
}

func TestTripexp3(t *testing.T) {
	if v := Tripexp(3, 10.00, 20.00, 30.00); v != 10.00 {
		t.Error("Expecting 10.00 but got ", v)
	}
}

func TestTripexp4(t *testing.T) {
	if v := Tripexp(4, 15.00, 15.01, 3.00, 3.01); v != 11.99 {
		t.Error("Expecting 11.99 but got ", v)
	}
}
