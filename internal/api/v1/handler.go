package v1

import (
	validator2 "github.com/sanservices/apicore/validator"

	"github.com/danic95/users_dashboard/internal/users_dashboard"
	"github.com/danic95/users_dashboard/goutils/settings"
)

// Handler handles v1 routes
type Handler struct {
	cfg       *settings.Settings
	service   users_dashboard.Service
	validator *validator2.Validator
}

// NewHandler initialize main *Handler
func NewHandler(cfg *settings.Settings, svc users_dashboard.Service, validator *validator2.Validator) *Handler {

	return &Handler{
		cfg:       cfg,
		service:   svc,
		validator: validator,
	}
}
