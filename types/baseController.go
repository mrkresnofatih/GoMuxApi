package types

import (
	"net/http"
)

type BaseEndpoint func(w http.ResponseWriter, r *http.Request)

type BaseMiddleware func(next http.Handler) http.Handler
