// users_dashboard service
//
// Schemes: http
// Host: localhost:8080
// BasePath: /v1
// Version: 1.0
//
// Security:
//     - api_key:
//
// SecurityDefinitions:
//  api_key:
//   type: apiKey
//   name: api-key
//   in: header
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"

	"github.com/danic95/users_dashboard/goutils/dbfactory"
	"github.com/danic95/users_dashboard/goutils/settings"
	"github.com/danic95/users_dashboard/internal/api"
	"github.com/danic95/users_dashboard/internal/api/healthcheck"
	v1 "github.com/danic95/users_dashboard/internal/api/v1"
	"github.com/danic95/users_dashboard/internal/users_dashboard"
	"github.com/danic95/users_dashboard/internal/users_dashboard/repository"
	"github.com/danic95/users_dashboard/internal/users_dashboard/service"
	"github.com/sanservices/apicore/validator"
	logger "github.com/sanservices/apilogger/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Set logger to custom one
		// fx.Logger(logger.New()),

		fx.Provide(
			// Provide new instances of structs
			// Empty context
			context.Background,
			// Logger
			logger.New,
			// Settings
			settings.New,
			// database connection
			dbfactory.New,
			// Repo
			repository.New,
			// Service
			service.New,
			// Validator
			validator.NewValidator,
			// New server
			api.NewServer,
			// Add all handlers here
			func(cfg *settings.Settings, srv users_dashboard.Service, vld *validator.Validator) []api.Handler {
				return []api.Handler{
					healthcheck.NewHandler(),     // For Healthchecks
					v1.NewHandler(cfg, srv, vld), // v1
				}
			},
		),
		fx.Invoke(
			// Use logger to initialize it globally
			func(ctx context.Context, l *logger.Logger) {
				logger.Info(ctx, logger.LogCatStartUp, "Initializing the app")
			},

			// Register routes
			api.RegisterRoutes,
		),
	)

	app.Run()
}
