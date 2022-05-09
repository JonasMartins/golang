/*
An "even ended number" is a number whose first and last digit are the same.
You mission, should you choose to accept it, is to count how many "even ended numbers" are
there that are a multiplication of two 4 digit numbers.
*/

package main

import (
	"fmt"
)

func main() {

	count := 0

	for a := 1000; a <= 9999; a ++ {
		for b := a; b <= 9999; b ++ {
			n :=  a * b

			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count ++
			}
		}
	}
	
	fmt.Println(count)
}