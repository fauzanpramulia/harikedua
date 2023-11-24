package controller

import (
	"fmt"
	"harikedua/config"
	"harikedua/model"
	"net/http"
	"harikedua/helpers"
	"github.com/labstack/echo/v4"
	
)


func UserLogin(ctx echo.Context) error{
	db := config.GetDB()
	Emp := model.Employee{}
	
	if err := ctx.Bind(&Emp); err != nil{
		return err
	}

	password :=""
	password = Emp.Password

	if err := db.Debug().Where("email =?", Emp.Email).Take(&Emp).Error; err != nil{
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":	"Unauthorized",
			"Message":	"Invalid Email",
		})
	}
	
	fmt.Println(password)
	fmt.Println(Emp.Password)
	comparePass := helpers.ComparePass([]byte(Emp.Password), []byte(password))

	if !comparePass{
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":	"Unauthorized",
			"Message":	"Invalid password",
		})
	}

	token := helpers.GenerateToken(uint(Emp.ID), Emp.Email)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}