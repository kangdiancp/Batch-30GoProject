package models

import (
	"fmt"
	"time"
)

type Contract struct {
	Employee //embedded
	overTime float64
}

type ContractInterface interface {
	SetOvertime()
}

func NewContract(fullName string,
	dateOfBirth time.Time,
	salary float64,
	overtime float64) *Contract {

	return &Contract{*NewEmployee(fullName, dateOfBirth, salary, Contracts), overtime}
}

func (contract *Contract) SetOvertime(overtime float64) {
	contract.overTime = overtime
}

func (contract *Contract) Info() string {
	return fmt.Sprintf("%v, Overtime : [%.2f]", contract.Employee, contract.overTime)
}
