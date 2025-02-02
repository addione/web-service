package manager

import "github.com/addione/web-service/repository"

type ManagerDIContainer struct {
	managers              *Managers
	repositoryDIContainer *repository.RepositoryDIContainer
}

type Managers struct {
	userManager *UserManager
}

func NewManagerDIContainer() *ManagerDIContainer {
	mdi := &ManagerDIContainer{
		repositoryDIContainer: repository.NewDIContainer(),
	}
	mdi.managers = &Managers{
		userManager: newUserManager(mdi),
	}
	return mdi
}

func (mdi *ManagerDIContainer) GetUserManager() *UserManager {
	return mdi.managers.userManager
}
