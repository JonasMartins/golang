package main

import "fmt"

type Expense interface {
	getName() string
	getWage(annual bool) float64
}

type Person struct {
	name, city string
}

func processObjects(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case User:
			fmt.Println("User:", value.name, "Price:", value.wage)
		case *User:
			fmt.Println("User Pointer:", value.name, "Price:", value.wage)
		case Post:
			fmt.Println("Post:", "Title: ", value.title, "Description: ", value.getDescription())
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

func main() {

	var expense Expense = &User{"Admin", "admin@email.com", 3000.0}

	data := []interface{}{
		expense,
		User{"User", "user@email.com", 48.95},
		Post{"Title test", "test description"},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	processObjects(data...)
}
