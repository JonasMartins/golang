package main

import (
	"fmt"
	"packages/post"
)

func main() {

	newPost := post.NewPost("2", "test2", "test2")

	fmt.Println("Hello from main")
	fmt.Println(*newPost)
	fmt.Println(newPost.Description())
}
