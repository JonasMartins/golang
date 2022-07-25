package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"
	"src/graph/model"
	"src/infra/orm/gorm/models"
	"src/main/auth"
	jwtLocal "src/main/auth/jwt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var expirationTime = time.Now().Add(2 * (time.Hour * 24))

// try to soft delete a user given an id
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteAction, error) {
	errArr := []*model.Error{}
	deletedUser := models.User{}
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

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {

	if w := auth.ForResponseWriterContext(ctx); w != nil {
		cookieName := os.Getenv("COOKIE_NAME")

		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now(),
		})
	}
	return true, nil
}

// test a login mutation, and return a token if valid credentials
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {

	errArr := []*model.Error{}
	result := model.AuthResponse{
		Token:  "",
		Errors: errArr,
	}
	_err := model.Error{
		Method:  "Login",
		Message: "",
		Field:   "id",
		Code:    500,
	}
	user := models.User{}
	if foundUser := r.DB.First(&user, "email = ?", input.Email); foundUser.Error != nil {
		_err.Message = fmt.Sprintf("could not found user with email %s", input.Email)
		_err.Code = 404
		_err.Field = "email"
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	matches, err := user.PasswordMatches(input.Password)

	if !matches {
		_err.Message = "wrong password"
		_err.Code = 200
		_err.Field = "password"
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	if err != nil {
		return &result, err
	}
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	myClais := &jwtLocal.ClaimsType{
		Id:             user.Base.ID.String(),
		Name:           user.Name,
		Email:          user.Email,
		StandardClaims: *claims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClais)

	tokenString, err := token.SignedString(jwtLocal.GetJwtSecret())
	if err != nil {
		return &result, err
	}

	if w := auth.ForResponseWriterContext(ctx); w != nil {
		cookieName := os.Getenv("COOKIE_NAME")

		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    tokenString,
			Path:     "/",
			HttpOnly: true,
			Expires:  expirationTime,
		})
	}

	result.Token = tokenString

	return &result, nil
}

// Register a new User and return a token
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.AuthResponse, error) {
	errArr := []*model.Error{}
	result := model.AuthResponse{
		Errors: errArr,
		Token:  "",
	}
	_err := model.Error{
		Method:  "RegisterUser",
		Message: "",
		Field:   "",
		Code:    500,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		_err.Message = err.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	user := models.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: string(hashedPassword),
	}

	if resultCreate := r.DB.Debug().Create(&user); resultCreate.Error != nil {
		_err.Message = resultCreate.Error.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	} else {
		claims := &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}

		myClais := &jwtLocal.ClaimsType{
			Id:             user.Base.ID.String(),
			Name:           user.Name,
			Email:          user.Email,
			StandardClaims: *claims,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClais)

		tokenString, err := token.SignedString(jwtLocal.GetJwtSecret())
		if err != nil {
			_err.Message = err.Error()
			result.Errors = append(result.Errors, &_err)
			return &result, nil
		}

		if w := auth.ForResponseWriterContext(ctx); w != nil {
			cookieName := os.Getenv("COOKIE_NAME")

			http.SetCookie(w, &http.Cookie{
				Name:     cookieName,
				Value:    tokenString,
				HttpOnly: true,
				Path:     "/",
				Expires:  expirationTime,
			})
		}

		result.Token = tokenString

		return &result, nil
	}
}

// Get an amount of users limited and offseted by args
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) (*model.UsersResponse, error) {

	errArr := []*model.Error{}
	users := []*models.User{}
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
	user := models.User{}
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
	user := models.User{}
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
	user := models.User{}
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

func (r *queryResolver) GetUsersChats(ctx context.Context, userId string) (*model.ChatsResponse, error) {
	errArr := []*model.Error{}
	chats := []*models.Chat{}
	messages := []*models.Message{}
	result := model.ChatsResponse{
		Chats:  chats,
		Errors: errArr,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}
	_err := model.Error{
		Method:  "GetUsersChats",
		Message: "",
		Field:   "userId",
		Code:    500,
	}

	rows, err := r.DB.Table("chat_members cm").Select("cm.chat_id").Joins("left join users u on u.id = cm.user_id").Where("u.id = ?", userId).Rows()
	if err != nil {
		_err.Message = err.Error()
		return &result, nil
	}
	defer rows.Close()
	type UsersChats struct {
		Id string
	}

	var chatsIds []string

	for rows.Next() {
		var chat UsersChats
		err := rows.Scan(
			&chat.Id,
		)
		if err != nil {
			_err.Message = err.Error()
		}
		chatsIds = append(chatsIds, chat.Id)
	}

	if foundChats := r.DB.Order("updated_at desc").Where("id", chatsIds).Preload("Members").Find(&chats); foundChats.Error != nil {
		_err.Message = foundChats.Error.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	// if anotherChats := r.DB.Order().Where()

	/*
		if foundChats := r.DB.Order("updated_at desc").Where("id", chatsIds).Preload("Members", func(tx *gorm.DB) *gorm.DB {
			return tx.Joins(`
				JOIN LATERAL (
					SELECT
					FROM users u
					WHERE u.id = chats
				)
			`)
		}).Find(&chats); foundChats.Error != nil {
			_err.Message = foundChats.Error.Error()
			result.Errors = append(result.Errors, &_err)
			return &result, nil
		}
	*/

	rows, err = r.DB.Table("messages m").Order("m.created_at desc").Select("m.id, m.updated_at, m.chat_id, m.author_id, m.body, u.id as user_id, u.name").Joins("left join users u on u.id = m.author_id").Where("m.chat_id", chatsIds).Limit(100).Rows()
	if err != nil {
		_err.Message = err.Error()
		return &result, nil
	}
	defer rows.Close()
	for rows.Next() {
		var message models.Message
		var user models.User
		err := rows.Scan(
			&message.ID,
			&message.CreatedAt,
			&message.ChatId,
			&message.AuthorId,
			&message.Body,
			&user.ID,
			&user.Name,
		)
		if err != nil {
			_err.Message = err.Error()
			return &result, nil
		}
		message.Author = &user
		messages = append(messages, &message)
	}

	for _, c := range chats {
		var messages = []*models.Message{}
		c.Messages = messages
	}

	//helpers.MapGenericRelation(chats, "ID", "Messages", messages, "ChatId")

	for _, m := range messages {
		for _, c := range chats {
			if m.ChatId != "" && m.ChatId == c.ID.String() {
				c.Messages = append(c.Messages, m)
			}
		}
	}

	result.Chats = chats

	return &result, nil

}
