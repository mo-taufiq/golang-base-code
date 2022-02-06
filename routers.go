package main

import (
	"net/http"

	"golang-base-code/delivery"
	"golang-base-code/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func Routers(router *gin.Engine, db *gorm.DB, rc *redis.Client) {
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Up",
		})
	})

	middleware := middleware.NewMiddleware(rc)
	delivery := delivery.NewDelivery(db)

	userRoute := router.Group("api/v1/users").Use(middleware.AuthMiddleware.AuthWithCheckRoleMiddleware([]string{"1"}))
	{
		userRoute.POST("", delivery.UserDelivery.CreateUser)
		userRoute.PUT("", delivery.UserDelivery.UpdateUser)
		userRoute.DELETE(":id", delivery.UserDelivery.DeleteUser)
		userRoute.GET("", delivery.UserDelivery.GetUser)
		userRoute.POST("upload-profile-image", delivery.UserDelivery.UploadProfileImage)
	}

	authRoute := router.Group("api/v1/auth")
	{
		authRoute.POST("sign-in", delivery.AuthDelivery.AuthSignIn)
	}

	documentRoute := router.Group("api/v1/document")
	{
		documentRoute.GET("pdf/invoice", delivery.DocumentDelivery.GetInvoicePDF)
		documentRoute.GET("excel/invoice", delivery.DocumentDelivery.GetInvoiceExcel)
	}
}
