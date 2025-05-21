package auth

import "errors"

type TokenType string

const (
	TokenTypeAccess TokenType = "househunt-helper-access"
)

var ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")
