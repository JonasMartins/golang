package drawer

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const dbTimeout = time.Second * 3

func (d *Drawer) GeneratePot() *[]uint8 {
	d.Pot = nil
	pot := []uint8{}
	var j uint8

	for j = 1; j <= MAX; j++ {
		pot = append(pot, j)
	}

	return &pot
}

func (d *Drawer) Draw(n uint8, round int) (*[]uint8, error) {

	if n > MAX {
		return nil, fmt.Errorf("invalid number to draw")
	}

	index := d.Find(n)

	if index >= (MAX - round) {
		return nil, fmt.Errorf("number already withdrawn")
	} else {
		d.Pot = d.WithDraw(index)
	}

	return d.Pot, nil

}

// Finds a number from pot and return its index
// If not Found, return MAX + 1
func (d *Drawer) Find(n uint8) int {
	//fmt.Printf("\npot length %d\n", len(*d.Pot))
	for i, x := range *d.Pot {
		if n == x {
			return i
		}
	}

	return MAX + 1
}

// make sure the number to withdraw exists
// then withdraw it from the pot then
// the pot gets rebuilt
func (d *Drawer) WithDraw(n int) *[]uint8 {

	var aux = *d.Pot
	aux = append(aux[:n], aux[n+1:]...)
	return &aux
}

// true if a int belongs to an array of ints
func (d *Drawer) CheckBallBelongs(n uint8, arr *[]uint8) bool {
	for _, x := range *arr {
		if n == x {
			return true
		}
	}
	return false
}

// generates the data and insert its into database
func (d *Drawer) GenerateData(amount int) {
	min := 1
	max := MAX
	draws := []uint8{}
	var ball uint8
	var x int
	var alreadyWithdrown bool
	contest := Contest{}
	d.Pot = d.GeneratePot()

	for i := 0; i < amount; i++ {
		for j := 0; j < 6; j++ {
			rand.Seed(time.Now().UnixNano())
			x = rand.Intn(max-min+1) + min
			ball = uint8(x)
			alreadyWithdrown = d.CheckBallBelongs(ball, &draws)
			if alreadyWithdrown {
				for {
					x = rand.Intn(max-min+1) + min
					ball = uint8(x)
					alreadyWithdrown = d.CheckBallBelongs(ball, &draws)
					if !alreadyWithdrown {
						break
					}
				}
			}
			draws = append(draws, ball)
			d.Pot, _ = d.Draw(ball, j)
		}

		contest.ID = i + 1
		contest.RealeseDate = time.Now()
		contest.Bola_1 = int(draws[0])
		contest.Bola_2 = int(draws[1])
		contest.Bola_3 = int(draws[2])
		contest.Bola_4 = int(draws[3])
		contest.Bola_5 = int(draws[4])
		contest.Bola_6 = int(draws[5])

		err := d.InsertGame(contest)
		if err != nil {
			panic("error inserting contest")
		}
		fmt.Print("\n")
		draws = draws[:0]
		d.Pot = d.GeneratePot()
	}

}

func (d *Drawer) InsertGame(contest Contest) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var Id int
	stmt := `insert into contests (realese_date, bola_1, bola_2, bola_3, bola_4, bola_5, bola_6)
				values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err := d.DB.QueryRowContext(ctx, stmt,
		contest.RealeseDate,
		contest.Bola_1,
		contest.Bola_2,
		contest.Bola_3,
		contest.Bola_4,
		contest.Bola_5,
		contest.Bola_6,
	).Scan(&Id)

	if err != nil {
		return err
	}
	return nil
}

func (d *Drawer) GetSome(chunk int, offset int) ([]*Contest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, realese_date, bola_1, bola_2, bola_3, bola_4, bola_5, bola_6 from contests limit $1 offset $2`

	rows, err := d.DB.QueryContext(ctx, query, chunk, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contests []*Contest
	for rows.Next() {
		var contest Contest
		err := rows.Scan(
			&contest.ID,
			&contest.RealeseDate,
			&contest.Bola_1,
			&contest.Bola_2,
			&contest.Bola_3,
			&contest.Bola_4,
			&contest.Bola_5,
			&contest.Bola_6,
		)
		if err != nil {
			log.Println("Error scanning ", err)
			return nil, err
		}

		contests = append(contests, &contest)
	}
	return contests, nil

}
func (d *Drawer) GetAll() ([]*Contest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, realese_date, bola_1, bola_2, bola_3, bola_4, bola_5, bola_6 from contests`

	rows, err := d.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contests []*Contest
	var contest Contest
	for rows.Next() {
		err := rows.Scan(
			&contest.ID,
			&contest.RealeseDate,
			&contest.Bola_1,
			&contest.Bola_2,
			&contest.Bola_3,
			&contest.Bola_4,
			&contest.Bola_5,
			&contest.Bola_6,
		)
		if err != nil {
			log.Println("Error scanning ", err)
			return nil, err
		}

		contests = append(contests, &contest)
	}

	return contests, nil

}
