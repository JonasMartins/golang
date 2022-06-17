package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"src/graph/generated"
	"src/graph/model"
	"src/infra/orm/gorm/models/user"
	jwtLocal "src/main/auth/jwt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.RegisterUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	_user := user.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: string(hashedPassword),
	}

	if result := r.DB.Create(&_user); result.Error != nil {
		return nil, result.Error
	} else {

		expirationTime := time.Now().Add(2 * (time.Hour * 24))

		claims := &jwt.StandardClaims{
			Id:        _user.Base.Id.String(),
			ExpiresAt: expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtLocal.GetJwtSecret())
		if err != nil {
			return nil, err
		}

		response := model.RegisterUserResponse{
			Token: tokenString,
			ID:    _user.Id.String(),
			Name:  _user.Name,
		}

		return &response, nil
	}
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	var _users []*model.User
	var users []user.User
	var userUax model.User
	if result := r.DB.Find(&users).Offset(int(*offset)).Limit(int(math.Min(10, float64(*limit)))); result.Error != nil {
		return nil, result.Error
	}

	for i := 0; i < len(users); i++ {
		userUax.ID = users[i].Base.Id.String()
		userUax.Email = users[i].Email
		userUax.Password = users[i].Password
		userUax.Name = users[i].Name
		_users = append(_users, &userUax)
	}

	return _users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
