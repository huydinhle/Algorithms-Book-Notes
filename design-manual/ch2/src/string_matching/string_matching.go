package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	printstrbychar("huy")
}

// SubStrMatch finds if string1 is a substring of
// string2 and return the starting index from string2
func SubStrMatch(str, substr string) int {
	for i := 0; i < len(str); i++ {
		if match := Strmatch(str, substr, i); match {
			return i
		}
	}
	return -1
}

// Strmatch returns if a string1 is a substring of
// string2 at index i
func Strmatch(str, substr string, index int) bool {
	// substr has to be smaller than len of the string - index
	if len(substr) > (len(str) - index) {
		fmt.Println("can't have substring too big")
		return false
	}
	//index has to be bigger than zero
	if index < 0 {
		return false
	}
	i := 0 // iterate on substr from 0 to end
	for j := 0; j < len(substr); j++ {
		if str[i] != substr[j] {
			fmt.Printf("str[i] = %+v\n", str[i])
			fmt.Printf("substr[j] = %+v\n", substr[j])
			return false
		}
		i++
	}
	return true
}

func printstrbychar(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[i] = %+v\n", str[i])
		fmt.Printf("type of str[i] = %T\n", str[i])
	}
}
