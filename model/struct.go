package model

import (
	"fmt"
	"harikedua/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type M map[string]interface{}

type Item struct{
	Name 	string `json:"name" form:"name"`
	EmployeeId int 
	Employee	*Employee `json:"-"`
}

// type Employee struct{
// 	ID			int `json:"id" form:"id"`
// 	Full_name	string `json:"full_name" form:"full_name"`
// 	Email		string `json:"email" form:"email"`
// 	Age			int `json:"age" form:"age"`
// 	Division	string `json:"division" form:"division"`

// }

type Employee struct{
	ID			int `json:"id" form:"id" swagger:"description(ID)"`
	Full_name	string `json:"full_name" form:"full_name" swagger:"description(Full_name)" valid:"required"`
	Email		string `json:"email" form:"email" swagger:"description(Email)" valid:"required"`
	Password	string `json:"password" form:"password" swagger:"description(Password)" valid:"required"`
	Age			int `json:"age" form:"age" swagger:"description(Age)" valid:"required"`
	Division	string `json:"division" form:"division" swagger:"description(Division)" valid:"required"`
	Item		[]Item 
}

type DeleteResponse struct{
	Status 		int `json:"status"`
	Message 	string `json:"message"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator. ValidateStruct(e)
	if errCreate != nil {
		err = errCreate
		fmt.Println("error guys" , err)
		return
	}

	e.Password = helpers.HashPass(e.Password)
	err = nil

	return
}