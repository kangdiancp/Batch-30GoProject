package main

import (
	"day07/practise01/models"
	"fmt"
	"time"
)

func main() {
	//declare object employee
	emps1 := new(models.Employee)
	emps1.SetFullName("kang dian")
	emps1.SetSalary(45_000)

	fmt.Println(emps1.ToString())

	emps2 := models.NewEmployee("Jhone", time.Now(), 50_000, models.Contracts)
	fmt.Println(emps2.ToString())

	emps2.SetSalary(100_000)

	emps3 := models.NewEmployee2("Herlin", 45_909)
	fmt.Println(emps3.ToString())

	emps3.SetFullName("Daemon")
	fmt.Println(emps3.ToString())
}
