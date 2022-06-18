package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"src/graph/model"
	"src/infra/orm/gorm/models/user"
)

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) (*model.UsersResponse, error) {
	var _users []*model.User
	var users []user.User
	var response model.UsersResponse

	if result := r.DB.Find(&users).Offset(int(*offset)).Limit(int(math.Min(10, float64(*limit)))); result.Error != nil {
		response.Errors = append(response.Errors, &model.Error{
			Message: result.Error.Error(),
			Method:  "Users",
			Field:   "-",
			Code:    500,
		})
		response.Users = nil

		return &response, nil
	}

	for i := 0; i < len(users); i++ {
		_users = append(_users, &model.User{
			ID:       users[i].Base.Id.String(),
			Email:    users[i].Email,
			Password: users[i].Password,
			Name:     users[i].Name,
		})
	}
	response.Errors = nil
	response.Users = _users

	return &response, nil
}
