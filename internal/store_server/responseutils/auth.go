package responseutils

import (
	"log/slog"

	"github.com/19fachri/store-app/internal/store_server/datalayer/models"
	"github.com/19fachri/store-app/internal/store_server/servicemodels"
	"github.com/19fachri/store-app/internal/store_server/utils"
)

func GenerateAuthResponse(profile *models.Profile) (*servicemodels.AuthResponse, error) {
	_, token, err := utils.GenerateToken(profile.ExternalID)
	if err != nil {
		slog.Error("GenerateAuthResponse: failed to generate token. Error : %v", err)
		return nil, err
	}

	authResponse := servicemodels.AuthResponse{
		AccessToken: token,
		UID:         profile.ExternalID,
	}

	return &authResponse, nil
}
