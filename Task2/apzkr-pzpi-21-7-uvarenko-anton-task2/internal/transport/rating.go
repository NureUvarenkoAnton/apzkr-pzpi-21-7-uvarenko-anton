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

type RatingHandler struct {
	ratingService iRatingService
}

func NewRatingHandler(ratingService iRatingService) *RatingHandler {
	return &RatingHandler{
		ratingService: ratingService,
	}
}

type iRatingService interface {
	GetRatingByIds(ctx context.Context, ids core.RatingByIdsParams) (core.Rating, error)
	GetRatingByRateeId(ctx context.Context, rateeId sql.NullInt32) ([]core.Rating, error)
	GetRatingByRaterId(ctx context.Context, raterId sql.NullInt32) ([]core.Rating, error)
	AddRating(ctx context.Context, params core.AddRatingParams, userType core.UsersUserType) error
	GetAvgRating(ctx context.Context, rateeId int) (int, error)
}

func (h *RatingHandler) AddRating(ctx *gin.Context) {
	type AddRatingPayload struct {
		RateeId int `json:"rateeId" binding:"required"`
		Value   int `json:"value" binding:"required"`
	}
	var payload AddRatingPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	raterId := ctx.GetInt64("user_id")
	userType := core.UsersUserType(ctx.GetString("user_type"))

	err = h.ratingService.AddRating(ctx, core.AddRatingParams{
		RaterID: sql.NullInt32{Int32: int32(raterId), Valid: true},
		RateeID: sql.NullInt32{Int32: int32(payload.RateeId), Valid: true},
		Value:   sql.NullInt32{Int32: int32(payload.Value), Valid: true},
	}, userType)
	if err != nil {
		if errors.Is(err, pkg.ErrEntityDuplicate) {
			ctx.AbortWithStatus(http.StatusConflict)
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

func (h *RatingHandler) GetAvgRating(ctx *gin.Context) {
	type GetRatingPayload struct {
		RateeId int `uri:"rateeId" binding:"required"`
	}
	var payload GetRatingPayload
	err := ctx.ShouldBindUri(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rating, err := h.ratingService.GetAvgRating(ctx, payload.RateeId)
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, api.AvgRatingResponse{
		AvgRating: rating,
	})
}
