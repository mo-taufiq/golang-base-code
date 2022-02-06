package authDelivery

import (
	"net/http"

	"golang-base-code/helper/request"
	"golang-base-code/helper/responseCode"
	"golang-base-code/model/userModel"
	"golang-base-code/useCase"

	"github.com/gin-gonic/gin"
	gologger "github.com/mo-taufiq/go-logger"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/gorm"
)

type IAuthDelivery interface {
	AuthSignIn(c *gin.Context)
}

type AuthDelivery struct {
	UseCase useCase.UseCase
}

func NewAuthDelivery(db *gorm.DB) IAuthDelivery {
	return &AuthDelivery{
		UseCase: *useCase.NewUseCase(db),
	}
}

func (a *AuthDelivery) AuthSignIn(c *gin.Context) {
	var requestBody userModel.UserWrite
	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}
	errorValidation := request.Validator(c, rules, &requestBody)
	if len(errorValidation) != 0 {
		gologger.Error(responseCode.ErrorValidation.Description)
		c.JSON(http.StatusBadRequest, gin.H{
			"response": responseCode.ErrorValidation,
			"error":    errorValidation,
			"data":     nil,
			"message":  "Failed to sign in.",
		})
		return
	}

	token, err := a.UseCase.AuthUseCase.AuthSignIn(requestBody)
	if err != nil {
		gologger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"response": responseCode.OK,
			"error":    err.Error(),
			"data":     nil,
			"message":  "Failed to sign in.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": responseCode.OK,
		"error":    nil,
		"data": map[string]interface{}{
			"token": token,
		},
		"message": "Successfully signed in.",
	})
}
