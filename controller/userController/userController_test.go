package userController

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"music-api-go/model"
	"music-api-go/repository/users"
	"net/http/httptest"
	"testing"
)

type suiteUsers struct {
	suite.Suite
	control *userController
	mock    *users.MockUserRepository
}

func (s *suiteUsers) SetupSuite() {
	s.mock = &users.MockUserRepository{}
	//s.control = &userController{s.mock}
}

func (s *suiteUsers) TestCreateUser() {
	mockUser := model.Users{
		Username: "fathazhar",
		Password: "azhar",
		Email:    "fathazhar@gmail.com",
		Fullname: "Fath Azhar",
	}
	s.mock.On("CreateUser", mockUser).Return(nil)

	testcases := []struct {
		Name             string
		ExpectStatusCode int
		Method           string
		HasReturnBody    bool
		ExpectBody       map[string]interface{}
	}{
		{
			"get all users",
			200,
			"POST",
			true,
			map[string]interface{}{
				"message": "success create users",
			},
		},
	}
	for _, testcase := range testcases {
		s.T().Run(testcase.Name, func(t *testing.T) {
			r := httptest.NewRequest(testcase.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.control.CreateUser(ctx)
			s.NoError(err)
			s.Equal(testcase.ExpectStatusCode, w.Result().StatusCode)
		})
	}
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
