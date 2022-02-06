package userUseCase

import (
	"fmt"
	"sort"
	"strings"

	"golang-base-code/helper/array"
	jsonHelper "golang-base-code/helper/json"
	"golang-base-code/helper/pagination"
	"golang-base-code/model/userModel"
	"golang-base-code/repository/userRepository"

	gologger "github.com/mo-taufiq/go-logger"
	"gorm.io/gorm"
)

type IUserUseCase interface {
	CreateUser(requestBody userModel.UserWrite) error
	DeleteUser(id uint) error
	UpdateUser(requestBody userModel.UserWrite) error
	GetUser(p pagination.Pagination, user userModel.UserFilter) ([]userModel.User, *pagination.Pagination, error)
}

type UserUseCase struct {
	UserRepository userRepository.IUserRepository
}

func NewUserUseCase(db *gorm.DB) IUserUseCase {
	return &UserUseCase{
		UserRepository: userRepository.NewUserRepository(db),
	}
}

func (u *UserUseCase) CreateUser(requestBody userModel.UserWrite) error {
	gologger.Info("Start create a new user.")
	gologger.Info("Remove duplicate value of array role id.")
	arrRoleID := array.RemoveDuplicateArrayOfString(requestBody.RoleID)
	gologger.Info("Sort value of array role id.")
	sort.Strings(arrRoleID)
	gologger.Info(fmt.Sprintf("role id, before remove duplicate and sort: %s", requestBody.RoleID))
	gologger.Info(fmt.Sprintf("role id, after remove duplicate and sort: %s", arrRoleID))

	user := userModel.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
		RoleID:   strings.Join(arrRoleID, ","),
	}
	gologger.Info(fmt.Sprintf("User: \n%s", jsonHelper.ConvertStructToTidyJSON(user)))

	gologger.Info("Start checking unique email.")
	existingUser, _ := u.UserRepository.Detail(userModel.UserFilter{Email: user.Email})
	gologger.Info(fmt.Sprintf("Existing user: \n%s", jsonHelper.ConvertStructToTidyJSON(existingUser)))
	if existingUser.Email != "" {
		gologger.Error(fmt.Sprintf("Email: %s, already used.", existingUser.Email))
		gologger.Info("Finish create a new user.")
		return fmt.Errorf("email: %s, already used", existingUser.Email)
	}

	err := u.UserRepository.Create(user)
	gologger.Info("Finish create a new user.")
	return err
}

func (u *UserUseCase) DeleteUser(id uint) error {
	gologger.Info("Start delete a user.")
	gologger.Info(fmt.Sprintf("Id: %d.", id))
	err := u.UserRepository.Delete(id)
	gologger.Info("Finish delete a user.")
	return err
}

func (u *UserUseCase) UpdateUser(requestBody userModel.UserWrite) error {
	gologger.Info("Start update a user.")

	gologger.Info("Check is ID exist.")
	gologger.Info(fmt.Sprintf("ID: %d", requestBody.ID))
	_, err := u.UserRepository.Detail(userModel.UserFilter{ID: requestBody.ID})
	if err != nil {
		gologger.Error(fmt.Sprintf("Error check is ID exist: %s", err.Error()))
		gologger.Info("Finish update a user.")
		return fmt.Errorf("error check is ID exist: %s", err.Error())
	}

	gologger.Info("Remove duplicate value of array role id.")
	arrRoleID := array.RemoveDuplicateArrayOfString(requestBody.RoleID)
	gologger.Info("Sort value of array role id.")
	sort.Strings(arrRoleID)
	gologger.Info(fmt.Sprintf("role id, before remove duplicate and sort: %s", requestBody.RoleID))
	gologger.Info(fmt.Sprintf("role id, after remove duplicate and sort: %s", arrRoleID))

	user := userModel.User{
		ID:     requestBody.ID,
		Name:   requestBody.Name,
		Email:  requestBody.Email,
		RoleID: strings.Join(requestBody.RoleID, ","),
	}
	gologger.Info(fmt.Sprintf("User: \n%s", jsonHelper.ConvertStructToTidyJSON(user)))

	gologger.Info("Start checking unique email.")
	existingUser, _ := u.UserRepository.Detail(userModel.UserFilter{Email: user.Email})
	gologger.Info(fmt.Sprintf("Existing user: \n%s", jsonHelper.ConvertStructToTidyJSON(existingUser)))
	if existingUser.Email != "" && existingUser.ID != user.ID {
		gologger.Error(fmt.Sprintf("Email: %s, already used.", existingUser.Email))
		gologger.Info("Finish update a user.")
		return fmt.Errorf("email: %s, already used", existingUser.Email)
	}

	err = u.UserRepository.Update(user)
	gologger.Info("Finish update a user.")
	return err
}

func (u *UserUseCase) GetUser(p pagination.Pagination, user userModel.UserFilter) ([]userModel.User, *pagination.Pagination, error) {
	gologger.Info("Start get users.")
	gologger.Info(fmt.Sprintf("User filter: \n%s", jsonHelper.ConvertStructToTidyJSON(user)))
	gologger.Info(fmt.Sprintf("User pagination: \n%s", jsonHelper.ConvertStructToTidyJSON(p)))
	total, err := u.UserRepository.Count(user)
	gologger.Info(fmt.Sprintf("Total users: %d", total))
	if err != nil {
		gologger.Error(fmt.Sprintf("Error get user: %s", err.Error()))
		gologger.Info("Finish get users.")
		return nil, nil, err
	}
	p2 := pagination.NewPagination(p.PageNumber, p.TotalItemsPerPage, total)

	users, err := u.UserRepository.Find(user, p2)
	gologger.Info("Finish get users.")
	return users, p2, err
}
