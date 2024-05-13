package transport

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService       iAuthService
	firbaseAuthClient *auth.Client
}

type iAuthService interface {
	RegisterUser(ctx context.Context, user core.CreateUserParams) (string, error)
	Login(ctx context.Context, payload core.CreateUserParams) (string, error)
}

func NewAuthHandler(service iAuthService, authClient *auth.Client) *AuthHandler {
	return &AuthHandler{
		authService:       service,
		firbaseAuthClient: authClient,
	}
}

func (h AuthHandler) RegisterUser(ctx *gin.Context) {
	type RegisterPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var payload RegisterPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrPayloadDecode)
		return
	}

	token, err := h.authService.RegisterUser(ctx, core.CreateUserParams{
		Email:    sql.NullString{String: payload.Email, Valid: true},
		Name:     sql.NullString{String: payload.Name, Valid: true},
		Password: sql.NullString{String: payload.Password, Valid: true},
	})
	if err != nil {
		if errors.Is(err, pkg.ErrEmailDuplicate) {
			ctx.JSON(http.StatusConflict, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (h AuthHandler) Login(ctx *gin.Context) {
	type LoginPayload struct {
		Email    string
		Password string
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
		if errors.Is(err, pkg.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		if errors.Is(err, pkg.ErrWrongPassword) {
			ctx.JSON(http.StatusConflict, err)
			return
		}

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
