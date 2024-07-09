package models

type Permanent struct {
	Employee //embedded
	//emps      Employee //not emmbedded
	insurance float64
}
