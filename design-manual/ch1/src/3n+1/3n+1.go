package main

import (
	"fmt"
)

var m map[int]int

func main() {
	m = map[int]int{}
	max1 := Cal3nplusoneOnArray(1, 100)
	fmt.Printf("max1 = %+v\n", max1)
	max2 := Cal3nplusoneOnArray(100, 200)
	fmt.Printf("max2 = %+v\n", max2)
	max3 := Cal3nplusoneOnArray(201, 210)
	fmt.Printf("max3 = %+v\n", max3)
	max4 := Cal3nplusoneOnArray(900, 1000)
	fmt.Printf("max4 = %+v\n", max4)
	max5 := Cal3nplusoneOnArray(1, 100000)
	fmt.Printf("max5 = %+v\n", max5)
}

// Cal3nplusoneOnArray calculate 3n+1 from i to j
// and return the the one with the max cycles
func Cal3nplusoneOnArray(i, j int) int {
	max := 1
	for n := i; n <= j; n++ {
		possmax := Cal3nplusonewithcache(n)
		// fmt.Printf("possmax of %+v = %+v\n", n, possmax)
		if possmax > max {
			max = possmax
		}
	}
	return max
}

// Cal3nplusone will calculate the 3n+1 algorithm of a
// n number
func Cal3nplusonewithcache(n int) int {
	original := n
	cycles := 0
	if n > 1000000 {
		return -1
	}
	for n != 1 {
		value, ok := m[n]
		if ok {
			m[original] = cycles + value
			// fmt.Printf("n = %+v\n", n)
			// fmt.Printf("value = %+v\n", value)
			return cycles + value
		}
		if n%2 == 0 {
			n /= 2
			cycles++
		} else {
			n = 3*n + 1
			cycles++
		}
	}
	cycles++
	m[original] = cycles
	return cycles
}

// Cal3nplusone will calculate the 3n+1 algorithm of a
// n number
func Cal3nplusone(n int) int {
	cycles := 0
	if n > 1000000 {
		return -1
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
			cycles++
		} else {
			n = 3*n + 1
			cycles++
		}
	}
	cycles++
	return cycles
}
