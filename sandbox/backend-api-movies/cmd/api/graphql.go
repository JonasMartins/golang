package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

var movies []*models.Movie

var movieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			}, "description": &graphql.Field{
				Type: graphql.String,
			}, "year": &graphql.Field{
				Type: graphql.Int,
			}, "release_date": &graphql.Field{
				Type: graphql.DateTime,
			}, "runtime": &graphql.Field{
				Type: graphql.Int,
			}, "rating": &graphql.Field{
				Type: graphql.Float,
			}, "created_at": &graphql.Field{
				Type: graphql.DateTime,
			}, "updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// schema definition
var fields = graphql.Fields{
	"movie": &graphql.Field{
		Type:        movieType,
		Description: "Get movie by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(g graphql.ResolveParams) (interface{}, error) {
			id, ok := g.Args["id"]
			if ok {
				for _, movie := range movies {
					if movie.ID == id {
						return movie, nil
					}
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(movieType),
		Description: "Get all movies",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return movies, nil
		},
	},
	"search": &graphql.Field{
		Type:        graphql.NewList(movieType),
		Description: "Search movies by title",
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var list []*models.Movie
			search, ok := params.Args["titleContains"].(string)
			if ok {
				for _, movie := range movies {
					if strings.Contains(movie.Title, search) {
						log.Println("Found")
						list = append(list, movie)
					}
				}
			}
			return list, nil
		},
	},
}

func (app *application) moviesGraphql(w http.ResponseWriter, r *http.Request) {
	movies, _ = app.models.DB.All()

	q, _ := io.ReadAll(r.Body)
	query := string(q)

	log.Println(query)

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		app.errorJSON(w, errors.New("failed to create schema"))
		return
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)
	if len(resp.Errors) > 0 {
		app.errorJSON(w, errors.New(fmt.Sprintf("failed %+v", resp.Errors)))
	}
	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
