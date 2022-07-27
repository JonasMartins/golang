package main

import (
	"fmt"
	"strings"
)

func Split() {
	s := "{1}"

	re := strings.NewReplacer("{", "", "}", "")

	res := re.Replace(s)

	v := strings.Split(res, ",")
	fmt.Println(v)
}
