package api

import "errors"

var ErrMissingAuthHeader = errors.New("missing authorization header")
var ErrInvalidAuthHeaderFormat = errors.New("invalid authorization header format")
var ErrUnsupportedAuthType = errors.New("unsupported authorization type in authorization header")
var ErrFetchingUnauthorizedTransaction = errors.New("transaction doesn't belong to the authenticated user")
var ErrSessionNotFound = errors.New("session not found")
var ErrBlockedSession = errors.New("blocked user session")
var ErrIncorrectSessionUser = errors.New("incorrect username for session")
var ErrIncorrectSessionToken = errors.New("incorrect refresh_token for session")
var ErrExpiredSession = errors.New("session has expired")
