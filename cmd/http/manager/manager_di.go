package manager

type ManagerDIContainer struct {
	userManager *UserManager
}

func NewManagerDIContainer() *ManagerDIContainer {
	mdi := &ManagerDIContainer{}
	mdi.userManager = newUserManager(mdi)
	return mdi
}

func (mdi *ManagerDIContainer) GetUserManager() *UserManager {
	return mdi.userManager
}
