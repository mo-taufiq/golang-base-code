package userDelivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gologger "github.com/mo-taufiq/go-logger"
	"github.com/spf13/cast"
	"github.com/thedevsaddam/govalidator"
	"github.com/ulule/deepcopier"
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/helper/pagination"
	"taufiq.code/golang-base-code/helper/request"
	"taufiq.code/golang-base-code/helper/responseCode"
	"taufiq.code/golang-base-code/model/userModel"
	"taufiq.code/golang-base-code/useCase"
)

type IUserDelivery interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type UserDelivery struct {
	UseCase useCase.UseCase
}

func NewUserDelivery(db *gorm.DB) IUserDelivery {
	return &UserDelivery{
		UseCase: *useCase.NewUseCase(db),
	}
}

func (u *UserDelivery) CreateUser(c *gin.Context) {
	var requestBody userModel.UserWrite
	rules := govalidator.MapData{
		"name":     []string{"required", "between:3,40"},
		"email":    []string{"required", "min:4", "email"},
		"password": []string{"required", "valid_password"},
		"role_id":  []string{"required"},
	}
	errorValidation := request.Validator(c, rules, &requestBody)
	if len(errorValidation) != 0 {
		gologger.Error(responseCode.ErrorValidation.Description)
		c.JSON(http.StatusBadRequest, gin.H{
			"response": responseCode.ErrorValidation,
			"error":    errorValidation,
			"data":     nil,
			"message":  "Failed to create new user.",
		})
		return
	}

	err := u.UseCase.UserUseCase.CreateUser(requestBody)
	if err != nil {
		gologger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": responseCode.OK,
			"error":    err.Error(),
			"data":     nil,
			"message":  "Failed to create new user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": responseCode.OK,
		"error":    nil,
		"data":     "",
		"message":  "Successfully created a new user.",
	})
}

func (u *UserDelivery) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := u.UseCase.UserUseCase.DeleteUser(cast.ToUint(id))
	if err != nil {
		gologger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": responseCode.OK,
			"error":    err.Error(),
			"data":     nil,
			"message":  "Failed to delete user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": responseCode.OK,
		"error":    nil,
		"data":     "",
		"message":  "Successfully deleted a user.",
	})
}

func (u *UserDelivery) UpdateUser(c *gin.Context) {
	var requestBody userModel.UserWrite
	rules := govalidator.MapData{
		"id":      []string{"required"},
		"name":    []string{"required", "between:3,40"},
		"email":   []string{"required", "min:4", "email"},
		"role_id": []string{"required"},
	}
	errorValidation := request.Validator(c, rules, &requestBody)
	if len(errorValidation) != 0 {
		gologger.Error(responseCode.ErrorValidation.Description)
		c.JSON(http.StatusBadRequest, gin.H{
			"response": responseCode.ErrorValidation,
			"error":    errorValidation,
			"data":     nil,
			"message":  "Failed to update new user.",
		})
		return
	}

	err := u.UseCase.UserUseCase.UpdateUser(requestBody)
	if err != nil {
		gologger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": responseCode.OK,
			"error":    err.Error(),
			"data":     nil,
			"message":  "Failed to update new user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": responseCode.OK,
		"error":    nil,
		"data":     "",
		"message":  "Successfully updated a new user.",
	})
}

func (u *UserDelivery) GetUser(c *gin.Context) {
	p := pagination.Pagination{
		PageNumber:        cast.ToInt64(c.Query("page_number")),
		TotalItemsPerPage: cast.ToInt64(c.Query("total_items_per_page")),
	}

	user := userModel.UserFilter{
		ID:            cast.ToUint(c.Query("id")),
		Name:          c.Query("name"),
		Email:         c.Query("email"),
		SortBy:        c.Query("sort_by"),
		SortDirection: c.Query("sort_direction"),
	}

	users, pagination, err := u.UseCase.UserUseCase.GetUser(p, user)
	if err != nil {
		gologger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": responseCode.OK,
			"error":    err.Error(),
			"data":     nil,
			"message":  "Failed to get user data.",
		})
		return
	}

	userReadArray := []userModel.UserRead{}
	for _, user2 := range users {
		userRead := &userModel.UserRead{}
		deepcopier.Copy(user2).To(userRead)
		userReadArray = append(userReadArray, *userRead)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": responseCode.OK,
		"error":    nil,
		"data": map[string]interface{}{
			"users":      userReadArray,
			"pagination": pagination,
		},
		"message": "Successfully get user data.",
	})
}
