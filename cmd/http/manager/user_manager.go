package manager

import "fmt"

type UserManager struct {
}

func newUserManager(mdi *ManagerDIContainer) *UserManager {
	return &UserManager{}
}

func (um *UserManager) CreateNewUser() {
	fmt.Println("creating user---- manager")
}
