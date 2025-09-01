package usrmgr

import (
	"context"
	"fmt"

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

func (s Service) SignUp(ctx context.Context, fullname, password string) (*coredata.User, error) {

	hashedPassword, err := s.hp.HashPassword([]byte(password))
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &coredata.User{
		FullName: fullname,
		Password: hashedPassword,
	}

	if err := s.pg.WithTx(ctx, func(tx pg.Conn) error {
		if err := user.Insert(ctx, tx); err != nil {
			return fmt.Errorf("cannot insert user: %w", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return user, nil
}
