package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"io"

	//"io"
	"math"
	"net/http"
	"os"
	"src/cmd/helpers/data"
	"src/cmd/utils"
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
	result := model.ChatsResponse{
		Chats:  nil,
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

	var query = `
		SELECT cm.chat_id,
		c.updated_at as chat_updated_at,
		m1.id AS message_id,
		m1.chat_id as message_chat_id,
		m1.body AS message_body, 
		m1.created_at AS message_created_at, 
		m1.author_id, 
		m1.author_name,
		m1.seen
		FROM chat_members cm
		LEFT JOIN chats c on cm.chat_id = c.id
		LEFT JOIN users u on cm.user_id = u.id
		JOIN LATERAL (
			SELECT m.id, m.chat_id, m.body, m.created_at, u1.id as author_id, 
			u1.name as author_name, m.seen
			FROM messages m
			INNER JOIN users u1 on m.author_id = u1.id
			WHERE m.chat_id = cm.chat_id
			ORDER BY m.created_at DESC
			LIMIT 10
		) AS m1 ON m1.chat_id = cm.chat_id
		WHERE u.id = ? 
		ORDER BY c.updated_at DESC
	`
	var queryResults []*utils.ResultGetUsersChats
	rows, err := r.DB.Raw(query, userId).Rows()
	if err != nil {
		_err.Message = err.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}
	defer rows.Close()

	for rows.Next() {
		var row utils.ResultGetUsersChats
		err := rows.Scan(
			&row.ChatId,
			&row.ChatUpdatedAt,
			&row.MessageId,
			&row.MessageChatId,
			&row.MessageBody,
			&row.MessageCreatedAt,
			&row.AuthorId,
			&row.AuthorName,
			&row.Seen,
		)
		if err != nil {
			_err.Message = err.Error()
			result.Errors = append(result.Errors, &_err)
			return &result, nil
		}

		queryResults = append(queryResults, &row)
	}

	query = `
		SELECT cm1.chat_id, u.id AS member_id, u.name
		FROM chat_members cm1
		LEFT JOIN users u ON u.id = cm1.user_id
		WHERE cm1.chat_id IN (
			SELECT cm.chat_id
			FROM chat_members cm
			WHERE cm.user_id = ?
		)`

	var chatMembersResult []*utils.ResultChatMembersByMemberId
	rows, err = r.DB.Raw(query, userId).Rows()
	if err != nil {
		_err.Message = err.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}
	defer rows.Close()

	for rows.Next() {
		var row utils.ResultChatMembersByMemberId
		err := rows.Scan(
			&row.ChatId,
			&row.MemberId,
			&row.MemberName,
		)
		if err != nil {
			_err.Message = err.Error()
			result.Errors = append(result.Errors, &_err)
			return &result, nil
		}

		chatMembersResult = append(chatMembersResult, &row)
	}

	_chats, err := data.GetUsersChatsFromRaw(queryResults, chatMembersResult)
	if err != nil {
		_err.Message = err.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}
	result.Chats = _chats

	return &result, nil
}

func (r *mutationResolver) ChangeProfilePicture(ctx context.Context, input model.UploadProfilePicture) (*model.CreateAction, error) {
	errArr := []*model.Error{}
	user := models.User{}
	result := model.CreateAction{
		Created: true,
		Errors:  errArr,
	}
	_err := model.Error{
		Method:  "ChangeProfilePicture",
		Message: "",
		Field:   "",
		Code:    500,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}

	if findUser := r.DB.First(&user, "id = ?", input.UserID); findUser.Error != nil {
		_err.Message = findUser.Error.Error()
		_err.Field = "userId"
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	fileContent, err := io.ReadAll(input.File.File)
	if err != nil {
		_err.Message = err.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	path, err := utils.HandleUploads(fileContent, input.File.Filename)
	if err != nil {
		_err.Message = err.Error()
		result.Errors = append(result.Errors, &_err)
		return &result, nil
	}

	user.Profile = path
	r.DB.Save(&user)

	return &result, nil
}
