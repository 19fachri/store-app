package user_handler

import (
	"log/slog"
	"net/http"

	"github.com/19fachri/store-app/internal/store_server/responseutils"
	"github.com/19fachri/store-app/internal/store_server/servicemodels"
	"github.com/19fachri/store-app/internal/store_server/services/user_service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AuthHandler struct {
	logger      *slog.Logger
	userService *user_service.Client
}

func NewAuthHandler(
	logger *slog.Logger,
	userService *user_service.Client,
) *AuthHandler {
	return &AuthHandler{
		logger:      logger,
		userService: userService,
	}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var request servicemodels.Register
	if err := ctx.ShouldBindJSON(&request); err != nil {
		a.logger.Error("Register: failed to bind request. Error : %v", err)
		ctx.JSON(http.StatusBadRequest, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	var profileRequest servicemodels.ProfileCreateRequest
	if err := copier.Copy(&profileRequest, &request); err != nil {
		a.logger.Error("Register: failed to copy request. Error : %v", err)
		ctx.JSON(http.StatusBadRequest, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	profile, err := a.userService.ProfileService.CreateProfile(ctx, profileRequest)
	if err != nil {
		a.logger.Error("Register: failed to create profile. Error : %v", err)
		ctx.JSON(http.StatusInternalServerError, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	_, err = a.userService.AuthService.CreateAuthEmailPassword(ctx, request.Email, request.Password, profile.ID)
	if err != nil {
		a.logger.Error("Register: failed to create auth. Error : %v", err)
		ctx.JSON(http.StatusInternalServerError, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	response, err := responseutils.GenerateAuthResponse(profile)
	if err != nil {
		a.logger.Error("Register: failed to generate auth response. Error : %v", err)
		ctx.JSON(http.StatusInternalServerError, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var request servicemodels.Login
	if err := ctx.ShouldBindJSON(&request); err != nil {
		a.logger.Error("Login: failed to bind request. Error : %v", err)
		ctx.JSON(http.StatusBadRequest, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	auth, err := a.userService.AuthService.Login(ctx, request.Email, request.Password)
	if err != nil {
		a.logger.Error("Login: failed to login. Error : %v", err)
		ctx.JSON(http.StatusInternalServerError, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	response, err := responseutils.GenerateAuthResponse(&auth.Profile)
	if err != nil {
		a.logger.Error("Login: failed to generate auth response. Error : %v", err)
		ctx.JSON(http.StatusInternalServerError, servicemodels.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
