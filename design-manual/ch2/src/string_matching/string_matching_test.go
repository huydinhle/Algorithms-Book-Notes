package main

import (
	"testing"
)

//positive flow of Strmatch
func TestStringMatchPositive(t *testing.T) {
	if match := Strmatch("huyle", "huy", 0); !match {
		t.Error("huy is a match of huyle but it's not")
	}
}

//str = "", substr="", index =0
func TestStringMatchZero(t *testing.T) {
	if match := Strmatch("", "", 0); !match {
		t.Error("match of empty str at 0 is supposed to be empty string")
	}
}

//len(substr) > len(str)
func TestStringMatchSubStrToobig(t *testing.T) {
	if match := Strmatch("a", "bcdik", 0); match {
		t.Error("if substr is > str then can't be a match")
	}
}

//index less than 0
func TestStringMatchIndexLessThanZero(t *testing.T) {
	if match := Strmatch("a", "bcdik", -1); match {
		t.Error("index less than zero meaning no match")
	}
}

//test substring match
func TestSubStringMatch(t *testing.T) {
	if match := SubStrMatch("iamgod", "god"); match == 3 {
		t.Error("god is the substring of iamgod")
	}
}
