package delivery

import (
	"golang-base-code/delivery/authDelivery"
	"golang-base-code/delivery/documentDelivery"
	"golang-base-code/delivery/userDelivery"

	"gorm.io/gorm"
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
