package repository

import (
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/repository/userRepository"
)

type Repository struct {
	UserRepository userRepository.IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: userRepository.NewUserRepository(db),
	}
}
