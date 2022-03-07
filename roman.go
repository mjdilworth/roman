package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello")

	fmt.Println(romanTodecimal("x"))

}

func romanTodecimal(rom string) (int, error) {

	iRet := 0
	m := make(map[string]int)
	m["i"] = 1
	m["v"] = 5
	m["x"] = 10
	m["l"] = 50
	m["c"] = 100
	m["d"] = 500
	m["m"] = 1000

	if len(rom) < 2 {
		return m[rom], nil
	}
	/* my good stuff */
	last := 0
	cur := 0

	//for roman numerals count right to left.  if left is small then take away other wise add
	for i := len(rom) - 1; i >= 0; i-- {
		cur = m[string(rom[i])]
		if cur >= last {
			iRet += cur
		} else {
			iRet -= cur
		}
		last = cur

	}

	return iRet, nil
}
