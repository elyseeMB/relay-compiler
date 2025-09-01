package services

import (
	"context"

	"go.gearno.de/kit/pg"
)

type (
	Service struct {
		pg *pg.Client
	}
)

func NewService(ctx context.Context, pgClient *pg.Client) (*Service, error) {

	svc := &Service{
		pg: pgClient,
	}

	return svc, nil
}
