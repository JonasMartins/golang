package main

type Post struct {
	title, description string
}

func (p Post) getDescription() string {
	return p.description
}
