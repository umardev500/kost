package constants

import "errors"

type ErrMsg string

const (
	ErrMsgTokenNotFound ErrMsg = "Token not found"
	ErrGetCount         ErrMsg = "Failed to select count"
)

type Error error

var (
	ErrNotAffected Error = errors.New("not affected")
)
