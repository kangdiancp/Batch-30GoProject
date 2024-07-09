package models

import (
	"fmt"
	"math/rand"
	"time"
)

//generate interface:
// emp *Employee EmployeeInterface , tekan enter

// constructor
func NewEmployee(fullName string, dateOfBirth time.Time,
	salary float64, status Status) *Employee {
	return &Employee{empId: rand.Intn(102), fullName: fullName, dateOfBirth: dateOfBirth,
		salary: salary, status: status}
}

func (emp *Employee) SetFullName(fullName string) {
	emp.fullName = fullName
}

func (emp *Employee) SetDateOfBirth(dateOfBirth time.Time) {
	emp.dateOfBirth = dateOfBirth
}

func (emp *Employee) SetSalary(salary float64) {
	emp.salary = salary
}

func (emp *Employee) SetStatus(status Status) {
	emp.status = status
}

func (emp *Employee) ToString() string {
	return fmt.Sprintf("Employee : [%d,%s,%s,%.2f,%s]", emp.empId, emp.fullName,
		emp.dateOfBirth, emp.salary, emp.status)
}
