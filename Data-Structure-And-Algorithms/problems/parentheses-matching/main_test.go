package main

import (
	"testing"
)

func TestMatch1(t *testing.T) {
	valid, index := matchParentheses("()()")
	if !valid {
		t.Error("valid has to be true")
	}
	if index != -1 {
		t.Error("index has to be -1. but it was ", index)
	}
}

func TestMatch2(t *testing.T) {
	valid, index := matchParentheses("()()(")
	if valid {
		t.Error("valid has to be false")
	}
	if index != 4 {
		t.Error("index has to be 4. but it was ", index)
	}
}

func TestMatch3(t *testing.T) {
	valid, index := matchParentheses("()(())))")
	if valid {
		t.Error("valid has to be false")
	}
	if index != 6 {
		t.Error("index has to be 6. but it was ", index)
	}
}

func TestMatch4(t *testing.T) {
	valid, index := matchParentheses("")
	if !valid {
		t.Error("valid has to be true")
	}
	if index != -1 {
		t.Error("index has to be -1. but it was ", index)
	}
}
