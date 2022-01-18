package authMiddleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gologger "github.com/mo-taufiq/go-logger"
	"github.com/spf13/cast"
	jsonWebToken "taufiq.code/golang-base-code/helper/jsonWebToken"
	"taufiq.code/golang-base-code/helper/responseCode"
)

type IAuthMiddleware interface {
	AuthWithCheckRoleMiddleware(roles []string) gin.HandlerFunc
}

type authMiddleware struct {
}

func NewAuthMiddleware() IAuthMiddleware {
	return &authMiddleware{}
}

func (u *authMiddleware) AuthWithCheckRoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		gologger.Info("Start checking valid roles and user roles.")
		token := u.getToken(c)
		claims, err := jsonWebToken.ParseJWT(token)
		userRoles := strings.Split(cast.ToString(claims["role_id"]), ",")
		gologger.Info(fmt.Sprintf("Valid roles: %s.", roles))
		gologger.Info(fmt.Sprintf("User roles: %s.", userRoles))
		if err != nil {
			gologger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"response": responseCode.OK,
				"error":    err.Error(),
				"data":     nil,
				"message":  "Failed to validating token.",
			})
			gologger.Info("Finish checking valid roles and user roles.")
			c.Abort()
			return
		}

		for _, role := range roles {
			for _, userRole := range userRoles {
				if strings.EqualFold(role, userRole) {
					gologger.Info("Finish checking valid roles and user roles.")
					c.Next()
					return
				}
			}
		}

		err = errors.New("role_id is not authorized")
		gologger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"response": responseCode.OK,
			"error":    err.Error(),
			"data":     nil,
			"message":  "Failed to validating token.",
		})
		gologger.Info("Finish checking valid roles and user roles.")
		c.Abort()
	}
}

func (u *authMiddleware) getToken(c *gin.Context) string {
	tokenBearer := c.Request.Header.Get("Authorization")
	token := strings.Replace(tokenBearer, "Bearer ", "", 1)
	return token
}
