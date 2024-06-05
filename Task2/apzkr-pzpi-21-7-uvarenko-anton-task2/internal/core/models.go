// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package core

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type UsersUserType string

const (
	UsersUserTypeDefault UsersUserType = "default"
	UsersUserTypePet     UsersUserType = "pet"
	UsersUserTypeWalker  UsersUserType = "walker"
	UsersUserTypeAdmin   UsersUserType = "admin"
)

func (e *UsersUserType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersUserType(s)
	case string:
		*e = UsersUserType(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersUserType: %T", src)
	}
	return nil
}

type NullUsersUserType struct {
	UsersUserType UsersUserType
	Valid         bool // Valid is true if UsersUserType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersUserType) Scan(value interface{}) error {
	if value == nil {
		ns.UsersUserType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersUserType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersUserType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersUserType), nil
}

type WalksState string

const (
	WalksStatePending    WalksState = "pending"
	WalksStateAccepted   WalksState = "accepted"
	WalksStateDeclined   WalksState = "declined"
	WalksStateInProccess WalksState = "in_proccess"
	WalksStateFinished   WalksState = "finished"
)

func (e *WalksState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WalksState(s)
	case string:
		*e = WalksState(s)
	default:
		return fmt.Errorf("unsupported scan type for WalksState: %T", src)
	}
	return nil
}

type NullWalksState struct {
	WalksState WalksState
	Valid      bool // Valid is true if WalksState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWalksState) Scan(value interface{}) error {
	if value == nil {
		ns.WalksState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WalksState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWalksState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WalksState), nil
}

type Pet struct {
	ID             int64
	OwnerID        sql.NullInt64
	Name           sql.NullString
	Age            sql.NullInt16
	AdditionalInfo sql.NullString
}

type Rating struct {
	RaterID sql.NullInt32
	RateeID sql.NullInt32
	Value   sql.NullInt32
}

type User struct {
	ID        int64
	Name      sql.NullString
	Email     sql.NullString
	Password  sql.NullString
	UserType  NullUsersUserType
	IsBanned  sql.NullBool
	IsDeleted sql.NullBool
	DeletedAt sql.NullTime
	CreatedAt sql.NullTime
}

type Walk struct {
	ID         int64
	OwnerID    sql.NullInt64
	WalkerID   sql.NullInt64
	PetID      sql.NullInt64
	StartTime  sql.NullTime
	FinishTime sql.NullTime
	State      NullWalksState
}

type WalkInfo struct {
	WalkID            int64
	StartTime         sql.NullTime
	FinishTime        sql.NullTime
	State             NullWalksState
	OwnerID           int64
	OwnerName         sql.NullString
	OwnerEmail        sql.NullString
	WalkerID          int64
	WalkerName        sql.NullString
	WalkerEmail       sql.NullString
	PetID             int64
	PetName           sql.NullString
	PetAge            sql.NullInt16
	PetAdditionalInfo sql.NullString
}
