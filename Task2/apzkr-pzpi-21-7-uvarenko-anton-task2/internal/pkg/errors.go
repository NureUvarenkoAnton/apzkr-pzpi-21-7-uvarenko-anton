package pkg

import "errors"

// transport errors
var ErrPayloadDecode = errors.New("can't decode payload")

// db erorrs
var (
	ErrDbInternal     = errors.New("db intenrnal error")
	ErrEmailDuplicate = errors.New("email is already taken")
	ErrRetrievingUser = errors.New("can't retrieve user")
	ErrUserNotFound   = errors.New("user does not exist")
)

// service errors
var (
	ErrEncryptingPassword = errors.New("can't encrypt password")
	ErrWrongPassword      = errors.New("wrong password")
)

// firebase errors
var ErrCreatingToken = errors.New("something failed generating jwt token")
