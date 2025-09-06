package db

import "go.gearno.de/kit/pg"

type (
	pgConfig struct {
		Username string `json:username`
		Password string `json:password`
		Database string `json:database`
	}
)

func (cfg pgConfig) Options(options ...pg.Option) []pg.Option {
	opts := []pg.Option{
		pg.WithUser(cfg.Username),
		pg.WithPassword(cfg.Password),
		pg.WithDatabase(cfg.Database),
	}

	opts = append(opts, options...)
	return opts

}
