package delivery

import (
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/delivery/authDelivery"
	"taufiq.code/golang-base-code/delivery/userDelivery"
)

type Delivery struct {
	UserDelivery userDelivery.IUserDelivery
	AuthDelivery authDelivery.IAuthDelivery
}

func NewDelivery(db *gorm.DB) *Delivery {
	return &Delivery{
		UserDelivery: userDelivery.NewUserDelivery(db),
		AuthDelivery: authDelivery.NewAuthDelivery(db),
	}
}
