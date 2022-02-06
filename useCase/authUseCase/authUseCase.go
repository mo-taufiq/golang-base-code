package authUseCase

import (
	"errors"

	"golang-base-code/helper/encryption"
	jsonWebToken "golang-base-code/helper/jsonWebToken"
	"golang-base-code/model/userModel"
	"golang-base-code/repository/userRepository"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type IAuthUseCase interface {
	AuthSignIn(requestBody userModel.UserWrite) (*string, error)
}

type AuthUseCase struct {
	UserRepository userRepository.IUserRepository
}

func NewAuthUseCase(db *gorm.DB) IAuthUseCase {
	return &AuthUseCase{
		UserRepository: userRepository.NewUserRepository(db),
	}
}

func (u *AuthUseCase) AuthSignIn(requestBody userModel.UserWrite) (*string, error) {
	uf := userModel.UserFilter{
		Email: requestBody.Email,
	}

	user, err := u.UserRepository.Detail(uf)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email or password is not valid")
		}
	}

	isMatch := encryption.CheckPasswordHash(requestBody.Password, user.Password)
	if !isMatch {
		return nil, errors.New("email or password is not valid")
	}

	token, err := jsonWebToken.GenerateJWT(jwt.MapClaims{
		"user_id":    user.ID,
		"user_email": user.Email,
		"user_name":  user.Name,
		"role_id":    user.RoleID,
	}, 60*60*1)

	return token, err
}
