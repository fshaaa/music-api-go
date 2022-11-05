package controller

import (
	"github.com/labstack/echo/v4"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/repository/userRepository"
	"music-api-go/utilities"
	"net/http"
	"time"
)

type UserController interface{}

type userController struct {
	userRepo userRepository.UserRepository
}

func NewUserController(uRepo userRepository.UserRepository) *userController {
	return &userController{uRepo}
}

func (u *userController) GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get user by id",
		"user":    user,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user model.Users
	c.Bind(&user)
	user.ID = utilities.CreateUUID()
	user.CreatedAt = time.Now().Format(time.RFC1123Z)
	user.UpdatedAt = user.CreatedAt

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create user",
		"user":    user,
	})
}

func (u *userController) LoginUser(c echo.Context) error {
	var req model.Users
	c.Bind(&req)

	user, err := u.userRepo.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	token, err := utilities.CreateToken(user.Username, user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes := dto.UserToken{
		user.Username,
		user.Email,
		token,
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login",
		"user":    userRes,
	})
}

func (u *userController) UpdateUser(c echo.Context) error {
	var req model.Users
	id := c.Param("id")
	c.Bind(&req)
	req.UpdatedAt = time.Now().Format(time.RFC1123Z)

	user, err := u.userRepo.UpdateUser(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update user",
		"user":    user,
	})
}

func (u *userController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete user",
	})
}

func (u userController) SearchUser(c echo.Context) error {
	var users []dto.User
	name := c.QueryParam("username")

	users, err := u.userRepo.SearchUser(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success search user",
		"user":    users,
	})
}
