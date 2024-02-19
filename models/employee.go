package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	EmployeeNumber int `uadmin:"read_only"`
	Name           string
	//Position string
	// Company
	Gender        Gender
	Position      string
	Company       Company
	CompanyID     uint
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
