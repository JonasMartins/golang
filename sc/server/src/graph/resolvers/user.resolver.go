package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"
	"src/graph/model"
	"src/infra/orm/gorm/models/user"
	auth "src/main/auth"
	jwtLocal "src/main/auth/jwt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var expirationTime = time.Now().Add(2 * (time.Hour * 24))

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

// test a login mutation, and return a token if valid credentials
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {

	user := user.User{}
	if foundUser := r.DB.First(&user, "email = ?", input.Email); foundUser.Error != nil {
		return nil, fmt.Errorf("could not found user with email %s", input.Email)
	}

	matches, err := user.PasswordMatches(input.Password)

	if !matches {
		return nil, fmt.Errorf("worg password")
	}

	if err != nil {
		return nil, err
	}
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	myClais := &jwtLocal.ClaimsType{
		Id:             user.Base.Id.String(),
		Name:           user.Name,
		Email:          user.Email,
		StandardClaims: *claims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClais)

	tokenString, err := token.SignedString(jwtLocal.GetJwtSecret())
	if err != nil {
		return nil, err
	}

	if w := auth.ForResponseWriterContext(ctx); w != nil {
		cookieName := os.Getenv("COOKIE_NAME")

		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   tokenString,
			Expires: expirationTime,
		})
	}

	response := model.AuthResponse{
		Token: "Successfully logged!",
	}

	return &response, nil
}

// Register a new User and return a token
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.AuthResponse, error) {
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

		claims := &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}

		myClais := &jwtLocal.ClaimsType{
			Id:             user.Base.Id.String(),
			Name:           user.Name,
			Email:          user.Email,
			StandardClaims: *claims,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClais)

		tokenString, err := token.SignedString(jwtLocal.GetJwtSecret())
		if err != nil {
			return nil, err
		}

		if w := auth.ForResponseWriterContext(ctx); w != nil {
			cookieName := os.Getenv("COOKIE_NAME")

			http.SetCookie(w, &http.Cookie{
				Name:    cookieName,
				Value:   tokenString,
				Expires: expirationTime,
			})
		}
		response := model.AuthResponse{
			Token: "Successfully registered",
		}

		return &response, nil
	}
}

// Get an amount of users limited and offseted by args
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) (*model.UsersResponse, error) {

	errArr := []*model.Error{}
	users := []*user.User{}
	result := model.UsersResponse{
		Users:  nil,
		Errors: errArr,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}
	err := model.Error{
		Method:  "GetUserById",
		Message: "",
		Field:   "id",
		Code:    500,
	}

	if foundUsers := r.DB.Find(&users).Offset(int(*offset)).Limit(int(math.Min(10, float64(*limit)))); foundUsers.Error != nil {
		err.Message = foundUsers.Error.Error()
		result.Errors = append(result.Errors, &err)
	} else {
		result.Users = users
	}

	return &result, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.UserResponse, error) {

	errArr := []*model.Error{}
	user := user.User{}
	result := model.UserResponse{
		User:   nil,
		Errors: errArr,
	}
	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
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
	errArr := []*model.Error{}
	user := user.User{}
	result := model.UserResponse{
		User:   nil,
		Errors: errArr,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}

	err := model.Error{
		Method:  "GetUserById",
		Message: "",
		Field:   "id",
		Code:    500,
	}
	if findUser := r.DB.First(&user, "email = ?", email); findUser.Error != nil {
		err.Message = findUser.Error.Error()
		result.Errors = append(result.Errors, &err)
	} else {
		result.User = &user
	}

	return &result, nil
}
func (r *queryResolver) GetUserByName(ctx context.Context, name string) (*model.UserResponse, error) {
	errArr := []*model.Error{}
	user := user.User{}
	result := model.UserResponse{
		User:   nil,
		Errors: errArr,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}

	err := model.Error{
		Method:  "GetUserById",
		Message: "",
		Field:   "id",
		Code:    500,
	}
	if findUser := r.DB.First(&user, "name like ?", "%"+name+"%"); findUser.Error != nil {
		err.Message = findUser.Error.Error()
		result.Errors = append(result.Errors, &err)
	} else {
		result.User = &user
	}

	return &result, nil
}
