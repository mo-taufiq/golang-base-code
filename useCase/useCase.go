package useCase

import (
	"golang-base-code/useCase/authUseCase"
	"golang-base-code/useCase/userUseCase"

	"gorm.io/gorm"
)

type UseCase struct {
	UserUseCase userUseCase.IUserUseCase
	AuthUseCase authUseCase.IAuthUseCase
}

func NewUseCase(db *gorm.DB) *UseCase {
	return &UseCase{
		UserUseCase: userUseCase.NewUserUseCase(db),
		AuthUseCase: authUseCase.NewAuthUseCase(db),
	}
}
