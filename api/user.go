package api

import (
	"echo-gorm/db"
	"echo-gorm/model"
	"net/http"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, users)
}
