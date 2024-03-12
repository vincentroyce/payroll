package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	User       uadmin.User
	UserID     uint
	IdNumber   string
	FirstName  string `uadmin:"list_exclude"`
	MiddleName string `uadmin:"list_exclude"`
	LastName   string `uadmin:"list_exclude"`
	SuffixName string `uadmin:"list_exclude"`
	Password   string `uadmin:"password;list_exclude"`

	Gender       Gender `uadmin:"list_exclude"`
	Address      string `uadmin:"html"`
	Worksite     Worksite
	WorksiteID   uint
	Department   Department
	DepartmentID uint
	Position     Position
	PositionID   uint
	CivilStatus  CivilStatus `uadmin:"list_exclude"`

	ContactNumber string
	Email         string `uadmin:"list_exclude"`
	Status        Status `uadmin:"list_exclude"`
	Supervisor    bool   `uadmin:"list_exclude"`
}

func (s *Employee) String() string {
	return s.LastName + ", " + s.FirstName + " " + s.MiddleName + " " + s.SuffixName
}

func (e *Employee) Save() {
	user := uadmin.User{}
	user.Username = e.IdNumber
	user.FirstName = e.FirstName
	user.LastName = e.LastName
	user.Password = "Password11"
	//user.Email = e.Username
	//e.UserID = e.Username
	user.Active = true



	
	user.Save()
	uadmin.Save(e)
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

type CivilStatus int

func (CivilStatus) Single() int {
	return 1
}
func (CivilStatus) Married() int {
	return 2
}
func (CivilStatus) Widowed() int {
	return 3
}
func (CivilStatus) Seperated() int {
	return 4
}
