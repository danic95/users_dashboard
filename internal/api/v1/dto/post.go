package dto

import "github.com/danic95/users_dashboard/internal/users_dashboard/entity"

// createUserRQ represents the parameters for Swagger documentation
//
// swagger:parameters createUserRQ
type createUserRQ struct {
	// in:body
	Body struct {
		Id_user string `json:"id_user" validate:"required"`
		Pass    string `json:"pass" validate:"required"`
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required"`
		Role    string `json:"role" validate:"required"`
	}
}

// createUserRS represents the response for Swagger documentation
//
// swagger:response createUserRS
type createUserRS struct {
	// in:body
	Msg string `json:"msg"`
}

// logInUserRQ represents the parameters for Swagger documentation
//
// swagger:parameters logInUserRQ
type logInUserRQ struct {
	// in:path
	Id_user string `json:"id_user" validate:"required"`
	Pass    string `json:"pass" validate:"required"`
}

// logInUserRS represents the response for Swagger documentation
//
// swagger:response logInUserRS
type logInUserRS struct {
	// in:body
	User *entity.User `json:"user"`
}

// sendRequestRQ represents the parameters for Swagger documentation
//
// swagger:parameters sendRequestRQ
type sendRequestRQ struct {
	// in:path
	User_from string `json:"user_from" validate:"required"`
	User_to   string `json:"user_to" validate:"required"`
}

// sendRequestRS represents the response for Swagger documentation
//
// swagger:response sendRequestRS
type sendRequestRS struct {
	// in:body
	Msg string `json:"msg"`
}

// deleteFriendRQ represents the parameters for Swagger documentation
//
// swagger:parameters deleteFriendRQ
type deleteFriendRQ struct {
	// in:path
	Id_user_1 string `json:"id_user_1" validate:"required"`
	Id_user_2 string `json:"id_user_2" validate:"required"`
}

// deleteFriendRS represents the response for Swagger documentation
//
// swagger:response deleteFriendRS
type deleteFriendRS struct {
	// in:body
	Msg string `json:"msg"`
}
