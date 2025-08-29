package main

import (
	"context"

	"github.com/elyseeMB/relay-compiler/pkg/db"
	"go.gearno.de/kit/unit"
)

var (
	version string = "unknown"
	env     string = "unknow"
)

func main() {
	impl := db.New()
	unit := unit.NewUnit(impl, "relay-compiler", version, env)
	err := unit.Run()
	if err != nil && err != context.Canceled {
		panic(err)
	}
}
