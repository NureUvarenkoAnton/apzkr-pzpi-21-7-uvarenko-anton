package transport

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/api"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/translate"

	"github.com/gin-gonic/gin"
)

type WalkHalder struct {
	walkService iWalkService
}

func NewWalkHandler(walkService iWalkService) *WalkHalder {
	return &WalkHalder{
		walkService: walkService,
	}
}

type iWalkService interface {
	CreateWalk(ctx context.Context, walkParams core.CreateWalkParams) error
	GetWalksByWalkerId(ctx context.Context, walkerID sql.NullInt64) ([]core.Walk, error)
	GetWalksByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Walk, error)
	UpdateWalkState(ctx context.Context, params core.UpdateWalkStateParams) error
	GetWalksInfoByParams(
		ctx context.Context,
		lang string,
		params core.GetWalkInfoByParamsParams,
	) ([]core.WalkInfo, error)
	GetWalkInfoByWalkId(ctx context.Context, lang string, walkId int64) (core.WalkInfo, error)
}

func (h *WalkHalder) CreateWalkRequest(ctx *gin.Context) {
	userId := ctx.GetInt64("user_id")
	if userId == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	type CreateWalkRequestPayload struct {
		WalkerId  int    `json:"walkerId" binding:"required"`
		PetId     int    `json:"petId" binding:"required"`
		StartTime string `json:"startTime" binding:"required"`
	}
	var payload CreateWalkRequestPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse(time.DateTime, payload.StartTime)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid start time"))
		return
	}

	err = h.walkService.CreateWalk(ctx, core.CreateWalkParams{
		OwnerID:   sql.NullInt64{Int64: int64(userId), Valid: true},
		WalkerID:  sql.NullInt64{Int64: int64(payload.WalkerId), Valid: true},
		PetID:     sql.NullInt64{Int64: int64(payload.PetId), Valid: true},
		StartTime: sql.NullTime{Time: startTime, Valid: true},
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

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *WalkHalder) GetWalksByParams(ctx *gin.Context) {
	type LangUri struct {
		Lang string `uri:"lang"`
	}
	var langPayload LangUri
	err := ctx.ShouldBindUri(&langPayload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if langPayload.Lang != translate.LANG_UA &&
		langPayload.Lang != translate.LANG_EN {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	type QueryParams struct {
		WalkerId int64 `form:"walkerId"`
		OwnerId  int64 `form:"ownerId"`
		PetId    int64 `form:"petId"`
	}
	var payload QueryParams
	err = ctx.BindQuery(&payload)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	walks, err := h.walkService.GetWalksInfoByParams(
		ctx,
		langPayload.Lang,
		core.GetWalkInfoByParamsParams{
			OwnerID:  sql.NullInt64{Int64: payload.OwnerId, Valid: payload.OwnerId != 0},
			WalkerID: sql.NullInt64{Int64: payload.WalkerId, Valid: payload.WalkerId != 0},
			PetID:    sql.NullInt64{Int64: payload.PetId, Valid: payload.PetId != 0},
		})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, api.SliceDbWalkInfoToAPIWalkInfo(walks))
}

func (h *WalkHalder) GetWalkInfoById(ctx *gin.Context) {
	type UriParams struct {
		Id   int64  `uri:"id"`
		Lang string `uri:"lang"`
	}
	var payload UriParams
	err := ctx.ShouldBindUri(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if payload.Lang != translate.LANG_EN &&
		payload.Lang != translate.LANG_UA {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	info, err := h.walkService.GetWalkInfoByWalkId(ctx, payload.Lang, payload.Id)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, api.DbWalkInfoToAPIWalkInfo(info))
}

func (h *WalkHalder) GetWalksBySelfId(ctx *gin.Context) {
	type LangUri struct {
		Lang string `uri:"lang" binding:"required"`
	}
	var langPayload LangUri
	err := ctx.ShouldBindUri(&langPayload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if langPayload.Lang != translate.LANG_EN &&
		langPayload.Lang != translate.LANG_UA {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	type QueryParams struct {
		WalksState string `form:"walkState"`
	}
	var requestPayload QueryParams
	err = ctx.BindQuery(&requestPayload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if requestPayload.WalksState != "" {
		if !slices.Contains([]core.WalksState{
			core.WalksStatePending,
			core.WalksStateAccepted,
			core.WalksStateDeclined,
			core.WalksStateInProccess,
			core.WalksStateFinished,
		}, core.WalksState(requestPayload.WalksState)) {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	id := ctx.GetInt64("user_id")
	if id == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userType := core.UsersUserType(ctx.GetString("user_type"))
	if userType != core.UsersUserTypeWalker &&
		userType != core.UsersUserTypeDefault &&
		userType != core.UsersUserTypeAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	payload := core.GetWalkInfoByParamsParams{}
	if userType == core.UsersUserTypeDefault || userType == core.UsersUserTypeAdmin {
		payload.OwnerID = sql.NullInt64{Int64: id, Valid: true}
	}
	if userType == core.UsersUserTypeWalker {
		payload.WalkerID = sql.NullInt64{Int64: id, Valid: true}
	}

	walks, err := h.walkService.GetWalksInfoByParams(ctx, langPayload.Lang, payload)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, api.SliceDbWalkInfoToAPIWalkInfo(walks))
}

func (h *WalkHalder) UpdateWalkState(ctx *gin.Context) {
	type UpdateWalkPayload struct {
		WalkId int    `json:"walkId"`
		State  string `json:"state"`
	}
	var payload UpdateWalkPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if payload.State != string(core.WalksStateAccepted) &&
		payload.State != string(core.WalksStateInProccess) &&
		payload.State != string(core.WalksStateDeclined) &&
		payload.State != string(core.WalksStateFinished) {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.walkService.UpdateWalkState(ctx, core.UpdateWalkStateParams{
		State: core.NullWalksState{WalksState: core.WalksState(payload.State), Valid: true},
		ID:    int64(payload.WalkId),
	})
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
