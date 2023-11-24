package controller

import (
	"fmt"
	"harikedua/config"
	"harikedua/model"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World!")
}

func JsonMap(ctx echo.Context) error {
	data := model.M{"Message": "hello", "counter": 2, "statusKode": http.StatusOK}
	return ctx.JSON(http.StatusOK, data)
}

func Page1(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	data := "Hello " + name
	result := fmt.Sprintf("%s", data)
	return ctx.JSON(http.StatusOK, result)
}

func User(ctx echo.Context) error {
	user := model.Item{}
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// func CreateEmployee(ctx echo.Context) error{
// 	db, err := config.Connect()
// 	if err != nil{
// 		fmt.Println(err)
// 	}

// 	employee :=model.Employee{}
// 	if err := ctx.Bind(&employee); err != nil{
// 		return err
// 	}

// 	sqlStatment := `INSERT INTO employees
// 	(full_name, email, age, division)
// 	VALUES ($1,$2,$3,$4)`

// 	_, err = db.Exec(sqlStatment, employee.Full_name,
// 		employee.Email, employee.Age, employee.Division)

// 	if err != nil{
// 		panic(err)
// 	}

// 	return ctx.JSON(http.StatusOK, employee)
// }

// CreateEmployee Godoc
// @Summary create new Employee
// @Description Create a new Employee
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param model.Employee body model.Employee true "Employee"
// @Success 200 {object} model.Employee
// @Router /employee [post]
func CreateEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := model.Employee{}

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	db.Debug().Create(&employee)

	fmt.Println("create Employee")

	return ctx.JSON(http.StatusOK, employee)
}

func CreateItem(ctx echo.Context) error {
	db := config.GetDB()
	item := model.Item{}

	userData, ok := ctx.Get("userData").(jwt.MapClaims)

	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   userData,
			"message": "Failed to get user data",
		})
	}

	userID := uint(userData["id"].(float64))

	if err := ctx.Bind(&item); err != nil {
		return err
	}

	item.EmployeeId = int(userID)
	db.Debug().Create(&item)

	fmt.Println("CreateItem")
	return ctx.JSON(http.StatusOK, item)
}

func UpdateEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := model.Employee{}

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	//jika tidak menggunakan ID maka akan membuat data baru
	// db.Save(&employee)

	//jika idnya tidak ditulis maka data tidak dibuat baru dan tidak diupdate
	db.Model(&employee).Where("id=?", employee.ID).
		Updates(model.Employee{
			Full_name: employee.Full_name,
			Email:     employee.Email,
			Age:       employee.Age,
			Division:  employee.Division,
		})

	fmt.Println("Update Employee")

	return ctx.JSON(http.StatusOK, employee)
}

func DeleteEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := model.Employee{}

	delResp := model.DeleteResponse{
		Status:  http.StatusOK,
		Message: "Employee deleted Success",
	}

	paramId := ctx.Param("id")

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	db.Model(&employee).Where("id=?", paramId).Delete(&employee)

	fmt.Println("Delete Employee")
	return ctx.JSON(http.StatusOK, delResp)
}

func Index(c echo.Context) error {
	tmpl := template.Must(template.ParseGlob("template/*.html"))

	type M map[string]interface{}
	data := make(M)

	data[config.CSRFKey] = c.Get(config.CSRFKey)
	return tmpl.Execute(c.Response(), data)
}

func SayHello(c echo.Context) error {
	type M map[string]interface{}
	data := make(M)

	if err := c.Bind(&data); err != nil {
		return err
	}

	message := fmt.Sprintf("Hello %s , My Gender %s", data["name"], data["gender"])
	return c.JSON(http.StatusOK, message)
}
