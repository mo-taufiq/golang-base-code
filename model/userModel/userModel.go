package userModel

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/helper/encryption"
	"taufiq.code/golang-base-code/model"
	"taufiq.code/golang-base-code/model/roleModel"
)

type UserFilter struct {
	ID            uint
	Name          string
	Email         string
	SortBy        string
	SortDirection string
}

type UserWrite struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	RoleID   []string `json:"role_id"`
	Password string   `json:"password"`
}

type UserRead struct {
	ID        uint             `json:"id" deepcopier:"field:ID"`
	Name      string           `json:"name" deepcopier:"field:Name"`
	Email     string           `json:"email" deepcopier:"field:Email"`
	Role      []roleModel.Role `json:"role" deepcopier:"field:RoleID"`
	OtherData *string          `json:"other_data" deepcopier:"field:OtherData"`
	model.ModelFieldsDefault
}

type User struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	RoleID    string  `json:"role_id"`
	OtherData *string `json:"other_data"`
	model.ModelFieldsDefault
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	encryptedPasswordByte, err := encryption.HashPassword(u.Password)
	u.Password = cast.ToString(encryptedPasswordByte)
	return err
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		encryptedPasswordByte, err := encryption.HashPassword(u.Password)
		u.Password = cast.ToString(encryptedPasswordByte)
		return err
	}
	return
}
