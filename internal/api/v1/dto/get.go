package dto

import "github.com/danic95/users_dashboard/internal/users_dashboard/entity"

// disclaimer: there are non-ok json tag values for some of the *RQ structs, because they are in path or somewhere else, so it is not actual json

// swagger:model getUsersRS
type getUsersRS struct {
	// in: body
	Body struct {
		User []entity.User `json:"user"`
	} `json:"data"`
}

// swagger:parameters getUserRQ
type getUserRQ struct {
	// in: query
	ID string `json:"id"`
}

// swagger:model getUserRS
type getUserRS struct {
	// in: body
	User *entity.User `json:"user"`
}

// swagger:parameters getFriendsRQ
type getFriendsRQ struct {
	// in: path
	ID string `json:"id"`
}

// swagger:model getFriendsRQ
type getFriendsRS struct {
	// in: body
	User []entity.User `json:"user"`
}

