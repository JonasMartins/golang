package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"src/graph/generated"
	"src/graph/model"
	"src/infra/orm/gorm/models/user"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.RegisterUserResponse, error) {

	/*
		base := base.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	*/

	_user := user.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}

	if result := r.DB.Create(&_user); result.Error != nil {
		fmt.Println(result.Statement.Vars...)
		return nil, result.Error
	} else {
		fmt.Println("result ", result)
		response := model.RegisterUserResponse{
			Token: "",
			ID:    "",
			Name:  "",
		}

		return &response, nil
	}
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
