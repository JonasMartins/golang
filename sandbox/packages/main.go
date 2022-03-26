package main

import (
	"fmt"
	"packages/post"
	"packages/user"
)

func main() {

	newPost := post.NewPost("2", "test2", "test2")

	newUser := user.NewUser("1", "Admin", "admin@email.com")
	newUser.AssignPostToUser(newPost)

	fmt.Println("Hello from main")
	fmt.Println(*newPost)
	fmt.Println(newPost.Description())
	fmt.Println(*newUser)
}
