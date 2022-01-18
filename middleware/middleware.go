package middleware

import (
	"github.com/go-redis/redis/v8"
	"taufiq.code/golang-base-code/middleware/authMiddleware"
)

type Middleware struct {
	AuthMiddleware authMiddleware.IAuthMiddleware
}

func NewMiddleware(rc *redis.Client) *Middleware {
	return &Middleware{
		AuthMiddleware: authMiddleware.NewAuthMiddleware(),
	}
}
