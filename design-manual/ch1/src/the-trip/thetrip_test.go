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

func TestTripexp5(t *testing.T) {
	if v := Tripexp3(3, 1000, 2000, 3000); v != 1000 {
		t.Error("Expecting 1000 but got ", v)
	}
}

func TestTripexp6(t *testing.T) {
	if v := Tripexp3(4, 1500, 1501, 300, 301); v != 1200 {
		t.Error("Expecting 1200 but got ", v)
	}
}
