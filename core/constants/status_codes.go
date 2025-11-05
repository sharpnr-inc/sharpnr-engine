package constants

type StatusCode int

const (
	StatusOK           StatusCode = 200
	StatusBadReq       StatusCode = 400
	StatusUnauthorized StatusCode = 401
	StatusForbidden    StatusCode = 403
	StatusNotFound     StatusCode = 404
	StatusInternalErr  StatusCode = 500
)
