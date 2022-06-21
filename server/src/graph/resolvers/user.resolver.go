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

// try to soft delete a user given an id
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteAction, error) {
	errArr := []*model.Error{}
	deletedUser := user.User{}
	result := model.DeleteAction{
		Message: "Server Error",
		Status:  "fail",
		Errors:  errArr,
	}
	err := model.Error{
		Method:  "DeleteUser",
		Message: "",
		Field:   "id",
		Code:    500,
	}
	if findUser := r.DB.First(&deletedUser, "id = ?", id); findUser.Error != nil {
		err.Message = findUser.Error.Error()
		result.Errors = append(result.Errors, &err)
		return &result, nil
	}

	if deletedUserResult := r.DB.Delete(&deletedUser); deletedUserResult.Error != nil {
		err.Message = deletedUserResult.Error.Error()
		result.Errors = append(result.Errors, &err)
		return &result, nil
	}

	result.Message = "User Successfully deleted"
	result.Status = "success"

	return &result, nil
}

// Register a new User and return a token
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.RegisterUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}

	user := user.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: string(hashedPassword),
	}

	if result := r.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	} else {

		expirationTime := time.Now().Add(2 * (time.Hour * 24))

		claims := &jwt.StandardClaims{
			Id:        user.Base.Id.String(),
			ExpiresAt: expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtLocal.GetJwtSecret())
		if err != nil {
			return nil, err
		}

		response := model.RegisterUserResponse{
			Token: tokenString,
			ID:    user.Id.String(),
			Name:  user.Name,
		}

		return &response, nil
	}
}

// Get an amount of users limited and offseted by args
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) (*model.UsersResponse, error) {

	var users []*user.User
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
	response.Errors = nil
	response.Users = users

	return &response, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.UserResponse, error) {

	errArr := []*model.Error{}
	user := user.User{}
	result := model.UserResponse{
		User:   nil,
		Errors: errArr,
	}

	err := model.Error{
		Method:  "GetUserById",
		Message: "",
		Field:   "id",
		Code:    500,
	}
	if findUser := r.DB.First(&user, "id = ?", id); findUser.Error != nil {
		err.Message = findUser.Error.Error()
		result.Errors = append(result.Errors, &err)
	} else {
		result.User = &user
	}

	return &result, nil
}
func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.UserResponse, error) {
	panic("not implemented yet")
}
func (r *queryResolver) GetUserByName(ctx context.Context, name string) (*model.UserResponse, error) {
	panic("not implemented yet")
}
