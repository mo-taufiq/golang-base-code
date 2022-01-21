package delivery

import (
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/delivery/authDelivery"
	"taufiq.code/golang-base-code/delivery/documentDelivery"
	"taufiq.code/golang-base-code/delivery/userDelivery"
)

type Delivery struct {
	UserDelivery     userDelivery.IUserDelivery
	AuthDelivery     authDelivery.IAuthDelivery
	DocumentDelivery documentDelivery.IDocumentDelivery
}

func NewDelivery(db *gorm.DB) *Delivery {
	return &Delivery{
		UserDelivery:     userDelivery.NewUserDelivery(db),
		AuthDelivery:     authDelivery.NewAuthDelivery(db),
		DocumentDelivery: documentDelivery.NewDocumentDelivery(db),
	}
}
