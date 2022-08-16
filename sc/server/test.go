package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	now := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if err := os.Mkdir("public/images/"+now, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created")
}
