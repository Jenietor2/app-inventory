package main

import (
	"context"

	"github.com/Jenietor2/app-inventory/database"
	"github.com/Jenietor2/app-inventory/internal/repository"
	"github.com/Jenietor2/app-inventory/internal/service"
	"github.com/Jenietor2/app-inventory/settings"
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
		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "jenietor0311@gmail.com", "joanieto", "123.")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "jenietor0311@gmail.com", "123.")
				if err != nil {
					panic(err)
				}

				if u.Username != "joanieto" {
					panic("wrong name")
				}
			},
		),
	)

	app.Run()
}
