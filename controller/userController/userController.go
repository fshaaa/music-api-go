package userController

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"music-api-go/dto"
	"music-api-go/model"
	"music-api-go/usecase/users"
	"music-api-go/util/middleware"
	"music-api-go/util/uuid"
	"net/http"
	"time"
)

type UserController interface{}

type userController struct {
	user users.UserUsecase
}

func NewUserController(user users.UserUsecase) *userController {
	return &userController{user}
}

func (u *userController) GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := u.user.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get users by id",
		"users":   user,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user model.Users
	c.Bind(&user)
	user.ID = uuid.CreateUUID()
	user.CreatedAt = time.Now().Format(time.RFC1123Z)
	user.UpdatedAt = user.CreatedAt

	err := u.user.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create users",
	})
}

func (u *userController) LoginUser(c echo.Context) error {
	var req model.Users
	c.Bind(&req)

	user, err := u.user.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	fmt.Println(user)

	token, err := middleware.CreateToken(user.Username, user.Email)
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
		"users":   userRes,
	})
}

func (u *userController) UpdateUser(c echo.Context) error {
	var req model.Users
	id := c.Param("id")
	c.Bind(&req)
	req.UpdatedAt = time.Now().Format(time.RFC1123Z)

	user, err := u.user.UpdateUser(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update users",
		"users":   user,
	})
}

func (u *userController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := u.user.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete users",
	})
}

func (u *userController) SearchUser(c echo.Context) error {
	var users []dto.User
	name := c.QueryParam("name")

	users, err := u.user.SearchUser(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success search users",
		"users":   users,
	})
}
