package admin

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sejutaCita/services/admin-service/delivery/middleware"
	"sejutaCita/services/admin-service/entities"
	"sejutaCita/services/admin-service/repository/admin"
)

type AdminUsecaseInferface interface {
	CreateAdmin(adminParam entities.Admin) error
	LoginAdmin(email string, password string) (string, error)
	GetAdminById(id string) (entities.Admin, error)
	UpdateAdmin(id string, adminParam entities.Admin) error
	DeleteAdmin(id string) error
}

type AdminUseCase struct {
	AdminRepository admin.AdminRepositoryInferface
}

func NewAdminUsecase(adminRepo admin.AdminRepositoryInferface) AdminUsecaseInferface {
	return &AdminUseCase{
		AdminRepository: adminRepo,
	}
}

func (auc *AdminUseCase) CreateAdmin(adminParam entities.Admin) error {
	fmt.Println("hit")
	err := auc.AdminRepository.CreateAdmin(adminParam)
	return err
}

func (auc *AdminUseCase) LoginAdmin(email string, password string) (string, error) {
	adminData, err := auc.AdminRepository.LoginAdmin(email)
	if err != nil {
		return "", errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(password))
	if err != nil {
		return "", errors.New("Wrong Password")
	}

	adminData.IdString = adminData.Id.Hex()
	token, _ := middleware.CreateToken(adminData.IdString, adminData.Name)

	return token, err
}

func (auc *AdminUseCase) GetAdminById(id string) (entities.Admin, error) {
	adminParam, err := auc.AdminRepository.GetAdminById(id)

	adminParam.IdString = adminParam.Id.Hex()
	return adminParam, err
}

func (auc *AdminUseCase) UpdateAdmin(id string, adminParam entities.Admin) error {
	err := auc.AdminRepository.UpdateAdmin(id, adminParam)
	if err != nil {
		errors.New("Failed to update admin")
	}
	return err
}

func (auc *AdminUseCase) DeleteAdmin(id string) error {
	err := auc.AdminRepository.DeleteAdmin(id)
	if err != nil {
		errors.New("Failed to delete admin")
	}
	return err
}
