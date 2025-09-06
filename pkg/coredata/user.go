package coredata

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/elyseeMB/relay-compiler/pkg/gid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.gearno.de/kit/pg"
)

type (
	User struct {
		ID       gid.GID `db:"id"`
		FullName string  `db:"fullname"`
		Password []byte  `db:"password"`
	}

	Users []*User

	ErrUserNotFound struct {
		Identifier string
	}

	ErrUserAlreadyExists struct {
		message string
	}
)

func (e ErrUserAlreadyExists) Error() string {
	return e.message
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user not found: %q", e.Identifier)
}

func (u *User) Insert(ctx context.Context, conn pg.Conn) error {

	q := `INSERT INTO users(id, fullname, password)
	VALUES(
		@id,
		@fullname,
		@password
	)
	`

	args := pgx.StrictNamedArgs{
		"id":       u.ID,
		"fullname": u.FullName,
		"password": u.Password,
	}

	_, err := conn.Exec(ctx, q, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && strings.Contains(pgErr.ConstraintName, "FullName") {
				return &ErrUserAlreadyExists{
					message: fmt.Sprintf("user with email %s already exists", u.FullName),
				}
			}
		}

		return err
	}

	return nil
}
