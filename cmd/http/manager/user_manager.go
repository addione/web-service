package manager

import (
	"strconv"

	"github.com/addione/web-service/model/request"
	"github.com/addione/web-service/repository"
)

type UserManager struct {
	userRepo *repository.UserRepo
}

func newUserManager(mdi *ManagerDIContainer) *UserManager {
	return &UserManager{
		userRepo: mdi.repositoryDIContainer.GetUserRepo(),
	}
}

func (um *UserManager) CreateNewUser(user *request.CreateUserParams) (string, error) {
	userId, err := um.userRepo.CreateNewUser(user)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(userId)), nil

}
