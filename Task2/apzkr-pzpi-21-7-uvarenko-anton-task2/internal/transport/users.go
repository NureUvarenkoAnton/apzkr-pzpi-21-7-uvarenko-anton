package transport

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/api"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService iUserService
}

func NewUserHandler(userService iUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type iUserService interface {
	DeleteUser(ctx context.Context, id int64) error
	BanUser(ctx context.Context, id int64) error
	GetUsers(ctx context.Context, params core.GetUsersParams) ([]core.User, error)
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	type GetUsersQueryParams struct {
		Name      string `form:"name,omitempty"`
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

	requesterUserType := core.UsersUserType(ctx.GetString("user_type"))
	userType := core.UsersUserType(payload.UserType)

	if requesterUserType != core.UsersUserTypeAdmin &&
		userType == core.UsersUserTypeAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	// if !slices.Contains(
	// 	[]core.UsersUserType{
	// 		core.UsersUserTypeAdmin,
	// 		core.UsersUserTypeDefault,
	// 		core.UsersUserTypeWalker,
	// 	},
	// 	userType) && userType != "" {
	// 	ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid user type"))
	// 	return
	// }

	users, err := h.userService.GetUsers(ctx, core.GetUsersParams{
		Name:      sql.NullString{String: payload.Name, Valid: payload.Name != ""},
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

	result := api.SliceDbUserToAPIUser(users)
	ctx.JSON(http.StatusOK, result)
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
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
