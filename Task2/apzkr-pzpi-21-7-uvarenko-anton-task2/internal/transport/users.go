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

type UserHandler struct {
	userService   iUserService
	ratingService iRatingService
}

func NewUserHandler(userService iUserService, ratingService iRatingService) *UserHandler {
	return &UserHandler{
		userService:   userService,
		ratingService: ratingService,
	}
}

type iUserService interface {
	DeleteUser(ctx context.Context, id int64) error
	BanUser(ctx context.Context, id int64) error
	GetUsers(ctx context.Context, params core.GetUsersParams) ([]api.UserResponse, error)
	GetById(ctx context.Context, id int64, requesterType core.UsersUserType) (api.UserResponse, error)
}

func (h *UserHandler) GetUsersAdmin(ctx *gin.Context) {
	type GetUsersQueryParams struct {
		UserType  string `form:"userType,omitempty"`
		IsBanned  *bool  `form:"isBanned,omitempty"`
		IsDeleted *bool  `form:"isDeleted,omitempty"`
	}
	var payload GetUsersQueryParams
	err := ctx.ShouldBindQuery(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userType := core.UsersUserType(payload.UserType)

	users, err := h.userService.GetUsers(ctx, core.GetUsersParams{
		UserType:  core.NullUsersUserType{UsersUserType: userType, Valid: userType != ""},
		IsBanned:  sql.NullBool{Bool: payload.IsBanned != nil && *payload.IsBanned, Valid: payload.IsBanned != nil},
		IsDeleted: sql.NullBool{Bool: payload.IsDeleted != nil && *payload.IsDeleted, Valid: payload.IsDeleted != nil},
	})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	type IdUriParams struct {
		Id int64 `json:"id"`
	}
	var payload IdUriParams
	err := ctx.ShouldBindUri(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userType := core.UsersUserType(ctx.GetString("user_type"))

	users, err := h.userService.GetById(ctx, payload.Id, userType)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if errors.Is(err, pkg.ErrForbiden) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetWalkers(ctx *gin.Context) {
	users, err := h.userService.GetUsers(ctx, core.GetUsersParams{
		UserType: core.NullUsersUserType{
			UsersUserType: core.UsersUserTypeWalker,
			Valid:         true,
		},
		IsBanned:  sql.NullBool{Bool: false, Valid: true},
		IsDeleted: sql.NullBool{Bool: false, Valid: true},
	})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	type DeleteUserPayload struct {
		Id int64 `uri:"id"`
	}
	var payload DeleteUserPayload
	err := ctx.ShouldBindUri(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(ctx, payload.Id)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) DeleteSelf(ctx *gin.Context) {
	id := ctx.GetInt64("user_id")

	err := h.userService.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) SetBanState(ctx *gin.Context) {
	type BanUserPayload struct {
		Id int64 `json:"id"`
	}
	var payload BanUserPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.userService.BanUser(ctx, payload.Id)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
