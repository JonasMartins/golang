package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id string) (*Movie, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, rating,  runtime, created_at, updated_at
				FROM movies WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.movie_id, mg.genre_id, g.genre_name
			FROM movies_genres mg
			left join genres g on (g.id = mg.genre_id)
			WHERE mg.movie_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[string]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieId,
			&mg.GenreId,
			&mg.Genre.Name,
		)

		if err != nil {
			return nil, err
		}

		genres[mg.ID] = mg.Genre.Name
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, rating,  runtime, created_at, updated_at
				FROM movies ORDER BY title`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(

			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		genreQuery := `SELECT mg.id, mg.movie_id, mg.genre_id, g.genre_name
				FROM movies_genres mg
				left join genres g on (g.id = mg.genre_id)
				WHERE mg.movie_id = $1`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

		genres := make(map[string]string)
		for rows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieId,
				&mg.GenreId,
				&mg.Genre.Name,
			)

			if err != nil {
				return nil, err
			}

			genres[mg.ID] = mg.Genre.Name
		}

		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}
