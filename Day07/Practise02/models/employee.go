package models

import (
	"time"
)

type (
	Status string
)

const (
	Permanents Status = "PERMANENT" //Permanents = iota + 1
	Contracts         = "CONTRACT"  //2
	Trainees          = "TRAINEE"   //3
)

const (
	Male   = iota + 1
	Female //2
)

type Employee struct {
	empId       int
	fullName    string
	dateOfBirth time.Time
	salary      float64
	status      Status
}

type EmployeeInterface interface {
	SetFullName(fullName string)
	SetDateOfBirth(hireDate time.Time)
	SetSalary(salary float64)
	SetStatus(status Status)
	ToString()
}

type Info interface {
	ToString()
}
