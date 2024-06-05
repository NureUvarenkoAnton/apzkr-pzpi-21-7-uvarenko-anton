package transport

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/api"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService iAuthService
}

type iAuthService interface {
	RegisterUser(ctx context.Context, user core.CreateUserParams, toHashPassword bool) (string, error)
	Login(ctx context.Context, payload core.CreateUserParams) (string, error)
	LoginPet(ctx context.Context, petId int64, ownerId int64) (string, error)
}

func NewAuthHandler(service iAuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (h AuthHandler) RegisterUser(ctx *gin.Context) {
	var payload api.RegisterPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, pkg.ErrPayloadDecode)
		return
	}

	if core.UsersUserType(payload.UserType) == core.UsersUserTypeAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	token, err := h.authService.RegisterUser(ctx, core.CreateUserParams{
		Email:    sql.NullString{String: payload.Email, Valid: true},
		Name:     sql.NullString{String: payload.Name, Valid: true},
		Password: sql.NullString{String: payload.Password, Valid: true},
		UserType: core.NullUsersUserType{UsersUserType: core.UsersUserType(payload.UserType), Valid: true},
	}, true)
	if err != nil {
		if errors.Is(err, pkg.ErrEmailDuplicate) {
			ctx.AbortWithError(http.StatusConflict, err)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.TokenResponse{
		Token: token,
	})
}

func (h AuthHandler) Login(ctx *gin.Context) {
	type LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var payload LoginPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrPayloadDecode)
		return
	}

	token, err := h.authService.Login(ctx, core.CreateUserParams{
		Email:    sql.NullString{String: payload.Email, Valid: true},
		Password: sql.NullString{String: payload.Password, Valid: true},
	})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if errors.Is(err, pkg.ErrForbiden) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		if errors.Is(err, pkg.ErrWrongPassword) {
			ctx.AbortWithStatus(http.StatusConflict)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, api.TokenResponse{
		Token: token,
	})
}

func (h AuthHandler) LoginPet(ctx *gin.Context) {
	ownerId := ctx.GetInt64("user_id")
	if ownerId == 0 {

		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	type LoginPetPayload struct {
		PetId int64 `json:"petId"`
	}

	var payload LoginPetPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authService.LoginPet(ctx, payload.PetId, ownerId)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if errors.Is(err, pkg.ErrForbiden) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.TokenResponse{
		Token: token,
	})
}
