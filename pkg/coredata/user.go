package coredata

import (
	"fmt"
)

type (
	User struct {
		ID       string `db:"id"`
		fullName string `db:"fullname"`
		Email    string `db:"email"`
	}

	Users []*User

	ErrUserNotFound struct {
		Identifier string
	}
)

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user not found: %q", e.Identifier)
}

// func (u *User) Insert(ctx context.Context, conn pg.Conn) error {

// }
