package types

import "github.com/elyseeMB/relay-compiler/pkg/coredata"

func NewUser(u *coredata.User) *User {
	return &User{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Role:     Role(u.Role),
	}
}
