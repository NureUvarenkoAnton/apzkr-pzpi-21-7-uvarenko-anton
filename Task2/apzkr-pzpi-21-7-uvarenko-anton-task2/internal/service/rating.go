package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"

	"github.com/go-sql-driver/mysql"
)

type RatingService struct {
	ratingRepo iRatingRepo
}

func NewRatingSrvice(ratingRepo iRatingRepo) *RatingService {
	return &RatingService{
		ratingRepo: ratingRepo,
	}
}

type iRatingRepo interface {
	GetWalksByParams(ctx context.Context, arg core.GetWalksByParamsParams) ([]core.Walk, error)
	AddRating(ctx context.Context, arg core.AddRatingParams) error
	RatingByIds(ctx context.Context, arg core.RatingByIdsParams) (core.Rating, error)
	RatingsByRaterId(ctx context.Context, raterID sql.NullInt32) ([]core.Rating, error)
	RatingsByRateeId(ctx context.Context, rateeID sql.NullInt32) ([]core.Rating, error)
}

func (s *RatingService) AddRating(ctx context.Context, params core.AddRatingParams, userType core.UsersUserType) error {
	getWalkParams := core.GetWalksByParamsParams{
		WalkState: core.NullWalksState{WalksState: core.WalksStateFinished, Valid: true},
	}
	if userType == core.UsersUserTypeWalker {
		getWalkParams.WalkerID = sql.NullInt64{Int64: int64(params.RaterID.Int32), Valid: true}
		getWalkParams.OwnerID = sql.NullInt64{Int64: int64(params.RateeID.Int32), Valid: true}
	}

	if userType == core.UsersUserTypeDefault {
		getWalkParams.OwnerID = sql.NullInt64{Int64: int64(params.RaterID.Int32), Valid: true}
		getWalkParams.WalkerID = sql.NullInt64{Int64: int64(params.RateeID.Int32), Valid: true}
	}
	walks, err := s.ratingRepo.GetWalksByParams(ctx, getWalkParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrForbiden
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return fmt.Errorf("%w: %w", pkg.ErrDbInternal, err)
	}
	if len(walks) == 0 {
		return pkg.ErrForbiden
	}

	err = s.ratingRepo.AddRating(ctx, params)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			err := err.(*mysql.MySQLError)

			if err.Number == 1062 {
				return pkg.ErrEntityDuplicate
			}
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}
	return nil
}

func (s *RatingService) GetRatingByRaterId(ctx context.Context, raterId sql.NullInt32) ([]core.Rating, error) {
	ratings, err := s.ratingRepo.RatingsByRaterId(ctx, raterId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	return ratings, nil
}

func (s *RatingService) GetAvgRating(ctx context.Context, rateeId int) (int, error) {
	ratings, err := s.GetRatingByRateeId(ctx, sql.NullInt32{Int32: int32(rateeId), Valid: true})
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, rating := range ratings {
		sum += int(rating.Value.Int32)
	}

	return sum / len(ratings), nil
}

func (s *RatingService) GetRatingByRateeId(ctx context.Context, rateeId sql.NullInt32) ([]core.Rating, error) {
	ratings, err := s.ratingRepo.RatingsByRateeId(ctx, rateeId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}
	return ratings, nil
}

func (s *RatingService) GetRatingByIds(ctx context.Context, ids core.RatingByIdsParams) (core.Rating, error) {
	rating, err := s.ratingRepo.RatingByIds(ctx, ids)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Rating{}, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return core.Rating{}, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	return rating, nil
}
