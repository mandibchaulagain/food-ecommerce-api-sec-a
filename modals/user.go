package modals

import "gorm.io/gorm"

//class
type User struct{
	gorm.Model
	FirstName string `json:"first-name"`
	MiddleName string `json:"middle-name"`
	LastName string `json:"last-name"`
	Age uint16 `json:"age"` //unsigned 
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	CommonModel
}
//cm CommonModel le k garxa?
//[]User is return type
func GetAllUsers()[]User{
	return []User{
		User{
			FirstName: "Mandib",
			MiddleName: "",
			LastName: "Chaulagain",
			Age: 22,
			Email:"mandibchaulagain8@gmail.com",
			Password: "mandib123",
			Address: "Kathmandu",
		},
		User{
			FirstName: "Dipak",
			MiddleName: "",
			LastName: "Shrestha",
			Age: 22,
			Email:"dipakshrestha@gmail.com",
			Password: "dipak123",
			Address: "Kathmandu",
		},
	}
}
