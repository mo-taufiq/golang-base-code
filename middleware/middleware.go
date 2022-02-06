package middleware

import (
	"golang-base-code/middleware/authMiddleware"

	"github.com/go-redis/redis/v8"
)

type Middleware struct {
	AuthMiddleware authMiddleware.IAuthMiddleware
}

func NewMiddleware(rc *redis.Client) *Middleware {
	return &Middleware{
		AuthMiddleware: authMiddleware.NewAuthMiddleware(),
	}
}
