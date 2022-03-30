package resolvers

import (
	"database/sql"
	user "graphql_two/entities"
	"graphql_two/utils"

	"github.com/graphql-go/graphql"
)

func GetUserType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "user",
		Description: "An user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The identifier of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*user.User); ok {
						return user.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*user.User); ok {
						return user.Name, nil
					}

					return nil, nil
				},
			},
			"email": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The email address of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*user.User); ok {
						return user.Email, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The created_at date of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*user.User); ok {
						return user.CreatedAt, nil
					}

					return nil, nil
				},
			},

			"updated_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The updated_at date of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*user.User); ok {
						return user.CreatedAt, nil
					}

					return nil, nil
				},
			},
		},
	})
}

func GetUserField(userType *graphql.Object, db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Get an author.",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)

			author := &user.User{}
			err := db.QueryRow("select id, name, email from public.user where id = $1", id).Scan(&author.ID, &author.Name, &author.Email)
			utils.CheckError(err)

			return author, nil
		},
	}
}

func GetUsersField(userType *graphql.Object, db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Get a list of users.",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)

			author := &user.User{}
			err := db.QueryRow("select id, name, email from authors where id = $1", id).Scan(&author.ID, &author.Name, &author.Email)
			utils.CheckError(err)

			return author, nil
		},
	}
}
