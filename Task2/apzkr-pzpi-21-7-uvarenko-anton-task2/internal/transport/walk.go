package transport

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"

	"github.com/gin-gonic/gin"
)

type WalkHalder struct {
	walkService walkService
}

func NewWalkHandler(walkService walkService) *WalkHalder {
	return &WalkHalder{
		walkService: walkService,
	}
}

type walkService interface {
	CreateWalk(ctx context.Context, walkParams core.CreateWalkParams) error
	GetWalksByWalkerId(ctx context.Context, walkerID sql.NullInt64) ([]core.Walk, error)
	GetWalksByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Walk, error)
	UpdateWalkState(ctx context.Context, params core.UpdateWalkStateParams) error
}

func (h *WalkHalder) CreateWalkRequest(ctx *gin.Context) {
	userId := ctx.GetInt64("user_id")
	if userId == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	type CreateWalkRequestPayload struct {
		WalkerId  int       `json:"walkedId"`
		PetId     int       `json:"petId"`
		StartTime time.Time `json:"startTime"`
	}
	var payload CreateWalkRequestPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.walkService.CreateWalk(ctx, core.CreateWalkParams{
		OwnerID:   sql.NullInt64{Int64: int64(userId), Valid: true},
		WalkerID:  sql.NullInt64{Int64: int64(payload.WalkerId), Valid: true},
		PetID:     sql.NullInt64{Int64: int64(payload.PetId), Valid: true},
		StartTime: sql.NullTime{Time: payload.StartTime, Valid: true},
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *WalkHalder) GetWalksByWalkerId(ctx *gin.Context) {
	type IdUriParam struct {
		WalkerId int `uri:"walkerId"`
	}
	var payload IdUriParam
	err := ctx.BindUri(&payload)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	walks, err := h.walkService.GetWalksByWalkerId(ctx, sql.NullInt64{Int64: int64(payload.WalkerId), Valid: true})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, walks)
}

func (h *WalkHalder) GetWalksByOwnerId(ctx *gin.Context) {
	type IdUriParam struct {
		OwnerId int `uri:"ownerId"`
	}
	var payload IdUriParam
	err := ctx.ShouldBindUri(&payload)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	walks, err := h.walkService.GetWalksByOwnerId(ctx, sql.NullInt64{Int64: int64(payload.OwnerId), Valid: true})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, walks)
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
