package main

type User struct {
	name, email string
	wage        float64
}

func (u *User) getName() string {
	return u.name
}

func (u *User) getWage(_ bool) float64 {
	return u.wage
}
