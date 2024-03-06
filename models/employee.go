package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	FirstName string
	LastName  string
	Username  string
	Password  string `uadmin:"password"`
	//Position string
	// Company
	Gender        Gender
	Company       Company
	CompanyID     uint
	Department    Department
	DepartmentID  uint
	ContactNumber string
	Email         string
	Status        Status
}

type Gender int

func (Gender) Male() int {
	return 1
}

func (Gender) Female() int {
	return 2
}

type Status int

func (Status) Trainee() int {
	return 1
}

func (Status) Probationary() int {
	return 2
}
func (Status) Regular() int {
	return 3
}

func (s *Employee) String() string {
	return s.LastName + " " + s.FirstName
}

func (e *Employee) Save() {
	user := uadmin.User{}
	user.Username = e.Username
	user.FirstName = e.FirstName
	user.LastName = e.LastName
	user.Password = e.Password
	user.Email = e.Username
	user.Active = true
	user.Save()
	uadmin.Save(e)
}
