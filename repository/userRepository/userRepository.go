package userRepository

import (
	"fmt"

	"golang-base-code/helper/pagination"
	"golang-base-code/model/userModel"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Delete(id uint) error
	Create(user userModel.User) error
	Update(user userModel.User) error
	Find(user userModel.UserFilter, p *pagination.Pagination) ([]userModel.User, error)
	Count(user userModel.UserFilter) (int64, error)
	Detail(f userModel.UserFilter) (userModel.User, error)
}

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (u *UserRepository) Delete(id uint) error {
	m := userModel.User{}
	db := u.Database.Delete(&m, id)
	return db.Error
}

func (u *UserRepository) Create(m userModel.User) error {
	db := u.Database.Create(&m)
	return db.Error
}

func (u *UserRepository) Update(m userModel.User) error {
	db := u.Database.Updates(&m)
	return db.Error
}

func (u *UserRepository) Find(f userModel.UserFilter, p *pagination.Pagination) ([]userModel.User, error) {
	m := []userModel.User{}
	db := u.Database.Model(m)
	sortBy := "id"
	if f.SortBy != "" {
		sortBy = f.SortBy
	}
	sortDirection := "asc"
	if f.SortDirection != "" {
		sortDirection = f.SortDirection
	}
	db.Order(fmt.Sprintf("%s %s", sortBy, sortDirection))
	if f.ID != 0 {
		db.Where("id = ?", f.ID)
	}
	if f.Name != "" {
		db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", f.Name))
	}
	if f.Email != "" {
		db.Where("email LIKE ?", fmt.Sprintf("%%%s%%", f.Email))
	}
	if p.PageNumber == 0 {
		db.Find(&m)
	} else {
		db.Limit(int(p.Limit)).Offset(int(p.Offset)).Find(&m)
	}
	return m, db.Error
}

func (u *UserRepository) Count(f userModel.UserFilter) (int64, error) {
	m := userModel.User{}
	var totalRecords int64
	db := u.Database.Model(m)
	if f.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", f.Name))
	}
	if f.Email != "" {
		db = db.Where("email LIKE ?", fmt.Sprintf("%%%s%%", f.Email))
	}
	db.Count(&totalRecords)
	return totalRecords, db.Error
}

func (u *UserRepository) Detail(f userModel.UserFilter) (userModel.User, error) {
	m := userModel.User{}
	db := u.Database.Model(m)
	if f.ID != 0 {
		db = db.First(&m, "id = ?", f.ID)
	}
	if f.Email != "" {
		db = db.First(&m, "email = ?", f.Email)
	}
	return m, db.Error
}
