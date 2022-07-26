package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "{1}"

    re := strings.NewReplacer("{", "", "}", "")

    res := re.Replace(s)


    v := strings.Split(res, ",")
    fmt.Println(v)
}
