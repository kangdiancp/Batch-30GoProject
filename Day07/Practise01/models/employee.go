package models

import (
	"fmt"
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
	fullName    string
	dateOfBirth time.Time
	salary      float64
	status      Status
}

// constructor
func NewEmployee(fullName string, dateOfBirth time.Time,
	salary float64, status Status) *Employee {
	return &Employee{fullName: fullName, dateOfBirth: dateOfBirth,
		salary: salary, status: status}
}

// overloading constructor
func NewEmployee2(fullName string, salary float64) *Employee {
	return &Employee{fullName: fullName, salary: salary}
}

// func receiver => methods, set public
func (emp *Employee) SetFullName(fullName string) {
	emp.fullName = fullName
}

// use pointer
func (emp *Employee) SetSalary(salary float64) {
	if salary <= 50_000 {
		emp.salary = salary
	}
}

// copy value
func (emp Employee) ToString() string {
	return fmt.Sprintf("Employee : [%s,%s,%.2f,%s]", emp.fullName,
		emp.dateOfBirth, emp.salary, emp.status)
}
