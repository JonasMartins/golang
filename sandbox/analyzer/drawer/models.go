package drawer

import "time"

type Models struct {
	User User
}

// User is the structure which holds one user from the database.
type User struct {
	ID          int       `json:"id"`
	RealeseDate time.Time `json:"realese_date"`
	Bola_1      int       `json:"bola_1"`
	Bola_2      int       `json:"bola_2"`
	Bola_3      int       `json:"bola_3"`
	Bola_4      int       `json:"bola_4"`
	Bola_5      int       `json:"bola_5"`
	Bola_6      int       `json:"bola_6"`
}
