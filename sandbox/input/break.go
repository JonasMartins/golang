package main

var arr1 = []string{"1", "2", "3"}
var arr2 = []string{"1", "2", "3", "4"}

func Break() {

	for _, a := range arr1 {
		for _, b := range arr2 {
			if a == b {
				break
			}
		}
	}
}
