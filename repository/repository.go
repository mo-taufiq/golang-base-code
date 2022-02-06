package repository

import (
	"golang-base-code/repository/userRepository"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepository userRepository.IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: userRepository.NewUserRepository(db),
	}
}
