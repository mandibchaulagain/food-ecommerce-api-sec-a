package modals

import "gorm.io/gorm"

//class
type User struct{
	gorm.Model
	CommonModel
	UserDto
}

//data transfer object
//aako data userdto ma convert hannu parxa
type UserDto struct {
	FirstName string `json:"first-name"`
	MiddleName string `json:"middle-name"`
	LastName string `json:"last-name"`
	Age uint16 `json:"age"` //unsigned 
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
}

//cm CommonModel le k garxa?
//[]User is return type
func GetAllUsers()[]User{
	user1:= User{
		Model: gorm.Model{
			ID:1,
		},
		CommonModel: CommonModel{
			Status: true,
			Priority: 1,

		},
		UserDto : UserDto{
			Email: "mandib@gmail.com",
			FirstName: "Dipak",
			LastName: "cgn",
			Age: 22,
			Address: "Kathmandu",
			Password: "password",
		},
	}
	return []User{
		user1,
	}
}

