package coredata

import (
	"fmt"
)

type (
	User struct {
		ID       string `db:"id"`
		FullName string `db:"fullname"`
		Email    string `db:"email"`
		Role     string `db:"role"`
	}

	Users []*User

	ErrUserNotFound struct {
		Identifier string
	}
)

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user not found: %q", e.Identifier)
}
