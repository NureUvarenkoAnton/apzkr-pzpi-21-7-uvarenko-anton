package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo   iAuthUserRepo
	jwtHandler jwt.JWT
}

const tokenTimeToLive = time.Hour * 24 * 7

type iAuthUserRepo interface {
	CreateUser(ctx context.Context, arg core.CreateUserParams) error
	GetUserByEmail(ctx context.Context, email sql.NullString) (core.User, error)
}

func NewAuthService(repo iAuthUserRepo, jwtHandler jwt.JWT) *AuthService {
	return &AuthService{
		userRepo:   repo,
		jwtHandler: jwtHandler,
	}
}

func (s AuthService) RegisterUser(ctx context.Context, user core.CreateUserParams) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%w: [%w]", pkg.ErrEncryptingPassword, err)
	}

	user.Password.String = string(pass)

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		err, ok := err.(*mysql.MySQLError)
		if !ok {
			return "", err
		}

		// duplicate entry
		if err.Number == 1062 {
			return "", pkg.ErrEmailDuplicate
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return "", pkg.ErrDbInternal
	}

	dbUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", pkg.ErrRetrievingUser
	}

	token, err := s.jwtHandler.GenUserToken(dbUser.ID, dbUser.UserType.UsersUserType, time.Now().Add(tokenTimeToLive))
	if err != nil {
		pkg.PrintErr(pkg.ErrCreatingToken, err)
		return "", pkg.ErrCreatingToken
	}

	return token, nil
}

func (s AuthService) Login(ctx context.Context, payload core.CreateUserParams) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		err, ok := err.(*mysql.MySQLError)
		if !ok {
			return "", err
		}

		// not found
		if err.Number == 1339 {
			return "", pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return "", fmt.Errorf("[ERROR] %w: [%w]", pkg.ErrDbInternal, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(payload.Password.String))
	if err != nil {
		return "", pkg.ErrWrongPassword
	}

	token, err := s.jwtHandler.GenUserToken(user.ID, user.UserType.UsersUserType, time.Now().Add(tokenTimeToLive))
	if err != nil {
		pkg.PrintErr(pkg.ErrCreatingToken, err)
		return "", pkg.ErrCreatingToken
	}

	return token, nil
}
