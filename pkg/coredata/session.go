package coredata

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"go.gearno.de/kit/pg"
)

type (
	Session struct {
		ID        int         `db:"id"`
		UserId    int         `db:"user_id"`
		Data      SessionData `db:"data"`
		ExpiredAt time.Time   `db:"expired_at"`
		CreateAt  time.Time   `db:"created_at"`
		UpdateAt  time.Time   `db:"update_at"`
	}

	SessionData struct{}
)

func (s *Session) Insert(ctx context.Context, conn pg.Conn) error {

	q := `INSERT INTO sessions
	(id,
	user_id,
	expired_at,
	created_at,
	updated_at
	)
	VALUES (
		@session_id,
		@user_id,
		@data,
		@expired_at,
		@created_at,
		@updated_at
	)`

	args := pgx.StrictNamedArgs{
		"session_id": s.ID,
		"user_id":    s.UserId,
		"data":       s.Data,
		"expired_at": s.ExpiredAt,
		"created_at": s.CreateAt,
		"update_at":  s.UpdateAt,
	}

	_, err := conn.Exec(ctx, q, args)

	return err
}
