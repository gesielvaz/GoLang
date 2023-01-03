package main

import (
	"context"

	"github.com/disturb/inventory/database"
	"github.com/disturb/inventory/internal/repository"
	"github.com/disturb/inventory/internal/service"
	"github.com/disturb/inventory/settings"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(),
	)
	app.Run()
}
