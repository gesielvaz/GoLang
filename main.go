package main

import (
	"context"

	"github.com/disturb/inventory/database"
	"github.com/disturb/inventory/settings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from  USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()
}
