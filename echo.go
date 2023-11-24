package main

import (
	// "fmt"
	"harikedua/config"
	"harikedua/controller"
	auth "harikedua/middleware"

	// "net/http"

	"github.com/labstack/echo/v4"

	middecho "github.com/labstack/echo/v4/middleware"
	// "strings"
	// "echo/controller"
	// "github.com/alecthomas/template"
	_ "harikedua/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type M map[string]interface{}

// @tittle API Documentation Employee
// @version 1.0
// @description This is a simple Employee Server
// @contact.Name Fauzan
// @contact.email fauzan@gmail.com
// @termOfService http://localhost
// @host localhost:9000
// @BasePath /
func main() {
	// defer db.Close()

	config.Connect()

	// tmpl := template.Must(template.ParseGlob("template/*.html"))
	
	r := echo.New()
	r.Use(middecho.CSRFWithConfig(middecho.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHeader,
		ContextKey:  config.CSRFKey,
	}))

	r.GET("/index", controller.Index)
	r.POST("/sayhello", controller.SayHello)

	

	// r.GET("/", controller.HelloWorld)

	// r.GET("/json", controller.JsonMap)

	// r.GET("/page1", controller.Page1)

	// r.Any("/user", controller.User)

	// r.POST("/employee", controller.CreateEmployee)

	// r.PUT("/employee", controller.UpdateEmployee)

	// r.DELETE("/employee/:id", controller.DeleteEmployee)

	//routes for login and register
	r.POST("/login", controller.UserLogin)
	r.POST("/register", controller.CreateEmployee)
	//-------------------------------------

	//gROUP ROUTES FOR EMPLOYEE
	emm := r.Group("/employee")
	emm.Use(auth.Authentication())
	emm.PUT("/", controller.UpdateEmployee)
	emm.DELETE("/:id", controller.DeleteEmployee)

	//group routes for item
	itm := r.Group("/item")
	itm.Use(auth.Authentication())
	itm.POST("/", controller.CreateItem)

	// swagger := e.Group("/swagger")
	// swagger.GET("/*any", echoSwagger.WrapHandler)
	// route for swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	r.Logger.Fatal(r.Start(":9000"))


}
