package useCase

import (
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/useCase/authUseCase"
	"taufiq.code/golang-base-code/useCase/userUseCase"
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
