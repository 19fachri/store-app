package user_service

import (
	"context"
	"log/slog"

	"github.com/19fachri/store-app/internal/store_server/datalayer/actions"
	"github.com/19fachri/store-app/internal/store_server/datalayer/models"
	"github.com/19fachri/store-app/internal/store_server/utils"
)

type AuthService struct {
	logger              *slog.Logger
	emailPasswordAction actions.AuthEmailPasswordActionInterface
}

func NewAuthService(
	logger *slog.Logger,
	action *actions.Client,
) *AuthService {
	return &AuthService{
		logger:              logger,
		emailPasswordAction: action.AuthEmailPasswordAction,
	}
}

func (s *AuthService) CreateAuthEmailPassword(ctx context.Context, email, password string, profileID uint) (*models.AuthEmailPassword, error) {
	auth := models.AuthEmailPassword{
		Email:     email,
		Password:  password,
		ProfileID: profileID,
	}

	newPassword, err := utils.HashPassword(password)
	if err != nil {
		s.logger.Error("CreateAuthEmailPassword: failed to hash password. Error : %v", err)
		return nil, err
	}
	auth.Password = newPassword

	err = s.emailPasswordAction.SaveAuthEmailPassword(&auth)
	if err != nil {
		s.logger.Error("CreateAuthEmailPassword: failed to save auth for email %v. Error : %v", email, err)
		return nil, err
	}

	return &auth, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.AuthEmailPassword, error) {
	auth, err := s.emailPasswordAction.GetByEmail(email)
	if err != nil {
		s.logger.Error("Login: failed to get auth for email %v. Error : %v", email, err)
		return nil, err
	}

	if err := utils.VerifyPassword(password, auth.Password); err != nil {
		s.logger.Error("Login: failed to check password. Error : %v", err)
		return nil, err
	}

	return auth, nil
}
