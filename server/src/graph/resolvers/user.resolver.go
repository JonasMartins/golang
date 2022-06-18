package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"math"
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
