package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdulmanafc2001/testing_ap/database"
	"github.com/abdulmanafc2001/testing_ap/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	return r
}

func TestCreateUser(t *testing.T) {
	db := database.InitDatabase()
	router := routerInit()

	router.POST("/", CreateUser(db))

	input := models.User{
		Name:  "manaf",
		Email: "manaf@gmail.com",
		Phone: "9995657894",
	}
	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	t.Log(res.Body)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)

}

func TestGetUserById(t *testing.T) {
	db := database.InitDatabase()
	router := routerInit()

	router.GET("/:id", GetUserById(db))

	req, _ := http.NewRequest(http.MethodGet, "/1", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestUpdateUser(t *testing.T) {
	db := database.InitDatabase()
	router := routerInit()

	router.PUT("/:id", UpdateUser(db))

	input := models.User{
		Name:  "shaheed",
		Email: "shaheed@gmail.com",
		Phone: "8527419630",
	}
	reqBody, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)

}

func TestDeleteUser(t *testing.T) {
	db := database.InitDatabase()
	router := routerInit()

	router.DELETE("/:id", DeleteUser(db))

	req, _ := http.NewRequest(http.MethodDelete, "/1", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}
