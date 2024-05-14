package transport

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService iUserProfileService
}

func NewProfileHandler(service iUserProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: service,
	}
}

type iUserProfileService interface {
	AddPet(ctx context.Context, pet core.AddPetParams) error
	GetPetById(ctx context.Context, id int64) (*core.Pet, error)
	UpdateUserData(ctx context.Context, userData core.UpdateUserParams) error
	UpdatePet(ctx context.Context, pet core.UpdatePetParams) error
	GetAllPetsByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Pet, error)
}

func (h ProfileHandler) AddPet(ctx *gin.Context) {
	type AddPetPayload struct {
		Name           string `json:"name"`
		Age            int    `json:"age"`
		AdditionalInfo string `json:"additional_info"`
	}

	var payload AddPetPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = h.profileService.AddPet(ctx, core.AddPetParams{
		OwnerID:        sql.NullInt64{Int64: ctx.GetInt64("user_id"), Valid: true},
		Name:           sql.NullString{String: payload.Name, Valid: true},
		Age:            sql.NullInt16{Int16: int16(payload.Age), Valid: true},
		AdditionalInfo: sql.NullString{String: payload.AdditionalInfo, Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
}

func (h ProfileHandler) GetOwnerPets(ctx *gin.Context) {
	pets, err := h.profileService.GetAllPetsByOwnerId(ctx, sql.NullInt64{
		Int64: ctx.GetInt64("user_id"),
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}
	}

	ctx.JSON(http.StatusOK, pets)
}

func (h ProfileHandler) UpdatePet(ctx *gin.Context) {
	type UpdatePetPayload struct {
		PetId          int    `json:"pet_id"`
		Name           string `json:"name"`
		Age            int    `json:"age"`
		AdditionalInfo string `json:"additional_info"`
	}
	var payload UpdatePetPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = h.profileService.UpdatePet(ctx, core.UpdatePetParams{
		Name:           sql.NullString{String: payload.Name, Valid: payload.Name != ""},
		Age:            sql.NullInt16{Int16: int16(payload.Age), Valid: payload.Age != 0},
		AdditionalInfo: sql.NullString{String: payload.AdditionalInfo, Valid: payload.AdditionalInfo != ""},
		ID:             int64(payload.PetId),
	})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
}

func (h ProfileHandler) UpdateUser(ctx *gin.Context) {
	type UpdateUserPayload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var payload UpdateUserPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = h.profileService.UpdateUserData(ctx, core.UpdateUserParams{
		Name:  sql.NullString{String: payload.Name, Valid: payload.Name != ""},
		Email: sql.NullString{String: payload.Email, Valid: payload.Email != ""},
		ID:    ctx.GetInt64("user_id"),
	})
	if err != nil {
		if errors.Is(err, pkg.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
}
