package user_handler

import (
	"log/slog"

	"github.com/19fachri/store-app/internal/store_server/services/user_service"
)

type Client struct {
	AuthHandler *AuthHandler
}

func New(
	logger *slog.Logger,
	userService *user_service.Client,
) *Client {
	return &Client{
		AuthHandler: NewAuthHandler(logger, userService),
	}
}
