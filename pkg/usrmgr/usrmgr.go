package usrmgr

import (
	"context"
	"fmt"
	"time"

	"github.com/elyseeMB/relay-compiler/pkg/coredata"
	"github.com/elyseeMB/relay-compiler/pkg/crypto/passwdhash"
	"go.gearno.de/kit/pg"
)

type (
	Service struct {
		pg *pg.Client
		hp *passwdhash.Profile
	}

	ErrSignupDisabled struct{}
)

func (e ErrSignupDisabled) Error() string {
	return "signup is disabled"
}

func NewService(ctx context.Context, pgClient *pg.Client, hp *passwdhash.Profile) (*Service, error) {

	return &Service{
		pg: pgClient,
		hp: hp,
	}, nil
}

func (s Service) SignUp(ctx context.Context, fullname, password string) (*coredata.User, *coredata.Session, error) {

	hashedPassword, err := s.hp.HashPassword([]byte(password))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot hash password: %w", err)
	}

	now := time.Now()

	user := &coredata.User{
		FullName: fullname,
		Password: hashedPassword,
	}

	var session *coredata.Session

	if err := s.pg.WithTx(ctx, func(tx pg.Conn) error {
		if err := user.Insert(ctx, tx); err != nil {
			return fmt.Errorf("cannot insert user: %w", err)
		}

		session = &coredata.Session{
			UserId:    user.ID,
			ExpiredAt: now.Add(24 * time.Hour),
			CreateAt:  now,
			UpdateAt:  now,
		}

		if err := session.Insert(ctx, tx); err != nil {
			return fmt.Errorf("cannot insert session : %w", err)
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

	return user, session, nil
}
