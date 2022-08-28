package admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sejutaCita/services/admin-service/delivery/helper"
	"sejutaCita/services/admin-service/delivery/middleware"
	"sejutaCita/services/admin-service/entities"
	"sejutaCita/services/admin-service/usecase/admin"
)

type AdminHandler struct {
	adminUseCase admin.AdminUsecaseInferface
}

func NewAdminHandler(adminUsecase admin.AdminUsecaseInferface) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminUsecase,
	}
}

func (ah *AdminHandler) CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			adminParam entities.Admin
			result     CreateAdminResponse
		)

		c.Bind(&adminParam)
		err := ah.adminUseCase.CreateAdmin(adminParam)
		if err != nil {
			result = CreateAdminResponse{
				Status:   "Failed",
				Messages: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, result)
		}
		result = CreateAdminResponse{
			Status:   "Succes",
			Messages: "Succes Create admmin",
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (ah *AdminHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login loginAdminRequest
		var result LoginAdminResponse

		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data"))
		}

		token, errorLogin := ah.adminUseCase.LoginAdmin(login.Email, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errorLogin.Error()))
		}

		result = LoginAdminResponse{
			Status: "Succes",
			Token:  token,
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (ah *AdminHandler) GetAdminById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := middleware.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, UnAuthorizeResponse{
				Status:   "Failed",
				Messages: "Unauthorized",
			})
		}
		admin, err := ah.adminUseCase.GetAdminById(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("user not found"))
		}
		var AdminResponseData = GetAdminResponse{
			Id:    admin.IdString,
			Name:  admin.Name,
			Email: admin.Email,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Succes", AdminResponseData))
	}

}

func (ah *AdminHandler) UpdateAdmin() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, role := middleware.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, UnAuthorizeResponse{
				Status:   "Failed",
				Messages: "Unauthorized",
			})
		}
		var admin entities.Admin
		var result UpdateAdminResponse
		c.Bind(&admin)

		err := ah.adminUseCase.UpdateAdmin(id, admin)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to update admin"))
		}

		result = UpdateAdminResponse{
			Name:     admin.Name,
			Email:    admin.Email,
			Password: admin.Password,
		}
		return c.JSON(http.StatusOK, result)

	}

}

func (ah *AdminHandler) DeleteAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := middleware.ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, UnAuthorizeResponse{
				Status:   "Failed",
				Messages: "Unauthorized",
			})
		}

		err := ah.adminUseCase.DeleteAdmin(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to delete admin"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete admin"))

	}

}
