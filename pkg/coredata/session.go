package coredata

import (
	"context"
	"fmt"
	"time"

	"github.com/elyseeMB/relay-compiler/pkg/gid"
	"github.com/jackc/pgx/v5"
	"go.gearno.de/kit/pg"
)

type (
	Session struct {
		ID        gid.GID     `db:"id"`
		UserId    gid.GID     `db:"user_id"`
		Data      SessionData `db:"data"`
		ExpiredAt time.Time   `db:"expired_at"`
		CreateAt  time.Time   `db:"created_at"`
		UpdateAt  time.Time   `db:"update_at"`
	}

	SessionData struct{}
)

func (s *Session) Insert(ctx context.Context, conn pg.Conn) error {

	q := `INSERT INTO sessions
	(
	user_id,
	data,
	expired_at,
	created_at,
	updated_at
	)
	VALUES (
		@user_id,
		@data,
		@expired_at,
		@created_at,
		@updated_at
	)`

	args := pgx.StrictNamedArgs{
		"user_id":    s.UserId,
		"data":       s.Data,
		"expired_at": s.ExpiredAt,
		"created_at": s.CreateAt,
		"updated_at": s.UpdateAt,
	}

	fmt.Printf("log %s", args)

	_, err := conn.Exec(ctx, q, args)

	return err
}
