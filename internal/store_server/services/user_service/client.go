package user_service

import (
	"log/slog"

	"github.com/19fachri/store-app/internal/store_server/datalayer/actions"
)

type Client struct {
	AuthService    *AuthService
	ProfileService *ProfileService
}

func New(
	logger *slog.Logger,
	action *actions.Client,
) *Client {
	return &Client{
		AuthService:    NewAuthService(logger, action),
		ProfileService: NewProfileService(logger, action),
	}
}
