package usrmgr

import (
	"context"
	"fmt"

	"github.com/elyseeMB/relay-compiler/pkg/coredata"
	"go.gearno.de/kit/pg"
)

type (
	Service struct {
		pg *pg.Client
	}
)

func NewService(ctx context.Context, pgClient *pg.Client) (*Service, error) {
	return &Service{
		pg: pgClient,
	}, nil
}

func (s Service) SignUp(ctx context.Context, fullname, password string) (*coredata.User, error) {

	user := &coredata.User{
		FullName: fullname,
		Password: password,
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
