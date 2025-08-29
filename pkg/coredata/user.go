package coredata

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.gearno.de/kit/pg"
)

type (
	User struct {
		ID       int    `db:"id"`
		FullName string `db:"fullname"`
		Password string `db:"password"`
	}

	Users []*User

	ErrUserNotFound struct {
		Identifier string
	}
)

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user not found: %q", e.Identifier)
}

func (u *User) Insert(ctx context.Context, conn pg.Conn) error {
	q := `INSERT INTO users(fullname, password) VALUES (
	@fullname,
	@password
 )`

	args := pgx.StrictNamedArgs{
		"fullname": u.FullName,
		"password": u.Password,
	}

	_, err := conn.Exec(ctx, q, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return pgErr
		}

		return err
	}

	return nil
}
