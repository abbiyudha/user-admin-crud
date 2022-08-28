package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
	"sejutaCita/services/apigateway/entities"
)

func UserProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var client http.Client
		var data entities.UserResponse
		var baseURL = os.Getenv("BASE_URL_USER") + "/user"

		request, err := http.NewRequest("GET", baseURL, nil)
		jwtToken := c.Get("user").(*jwt.Token)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken.Raw))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}

		return c.JSON(http.StatusOK, data)
	}

}

func LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			client  http.Client
			data    LoginUserRequest
			baseURL = os.Getenv("BASE_URL_USER") + "/user/login"
		)
		c.Bind(&data)

		var bodyJson, _ = json.Marshal(data)
		var body = bytes.NewBufferString(string(bodyJson))

		request, err := http.NewRequest("POST", baseURL, body)
		request.Header.Set("Content-Type", "application/json;")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		var resultResponse LoginUserResponse
		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &resultResponse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resultResponse)
		}

		return c.JSON(http.StatusOK, resultResponse)
	}

}

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			client  http.Client
			data    CreateUserRequest
			baseURL = os.Getenv("BASE_URL_USER") + "/create/user"
		)
		c.Bind(&data)

		var bodyJson, _ = json.Marshal(data)
		var body = bytes.NewBufferString(string(bodyJson))

		request, err := http.NewRequest("POST", baseURL, body)
		jwtToken := c.Get("user").(*jwt.Token)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken.Raw))
		request.Header.Set("Content-Type", "application/json;")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		var resultResponse CreateUserResponse
		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &resultResponse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resultResponse)
		}

		return c.JSON(http.StatusOK, resultResponse)
	}

}
