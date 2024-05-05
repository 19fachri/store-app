package user_service

import (
	"context"
	"log/slog"

	"github.com/19fachri/store-app/internal/store_server/datalayer/actions"
	"github.com/19fachri/store-app/internal/store_server/datalayer/models"
	"github.com/19fachri/store-app/internal/store_server/servicemodels"
	"github.com/jinzhu/copier"
)

type ProfileService struct {
	logger        *slog.Logger
	profileAction actions.ProfileActionInterface
}

func NewProfileService(
	logger *slog.Logger,
	action *actions.Client,
) *ProfileService {
	return &ProfileService{
		logger:        logger,
		profileAction: action.ProfileAction,
	}
}

func (s *ProfileService) CreateProfile(ctx context.Context, request servicemodels.ProfileCreateRequest) (*models.Profile, error) {
	var profile models.Profile
	err := copier.Copy(&profile, &request)
	if err != nil {
		s.logger.Error("CreateProfile: failed to copy profile. Error : %v", err)
		return nil, err
	}

	err = s.profileAction.SaveProfile(&profile)
	if err != nil {
		s.logger.Error("CreateProfile: failed to save profile. Error : %v", err)
		return nil, err
	}

	return &profile, nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, request servicemodels.ProfileUpdateRequest, profile *models.Profile) error {
	if err := copier.Copy(profile, &request); err != nil {
		s.logger.Error("UpdateProfile: failed to copy profile. Error : %v", err)
		return err
	}

	err := s.profileAction.SaveProfile(profile)
	if err != nil {
		s.logger.Error("UpdateProfile: failed to save profile. Error : %v", err)
		return err
	}

	return nil
}

func (s *ProfileService) IsEmailUsed(ctx context.Context, email string) (bool, error) {
	count, err := s.profileAction.CountProfileByEMail(email)
	if err != nil {
		s.logger.Error("IsEmailUsed: failed to count profile. Error : %v", err)
		return false, err
	}

	return count > 0, nil
}

func (s *ProfileService) IsUsernameUsed(ctx context.Context, username string) (bool, error) {
	count, err := s.profileAction.CountProfileByUsername(username)
	if err != nil {
		s.logger.Error("IsUsernameUsed: failed to count profile. Error : %v", err)
		return false, err
	}

	return count > 0, nil
}
