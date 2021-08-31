//结构题是一种聚合数据类型，是由零个或多个任意类型聚合成的实体
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
	//a         Employee
}

var dilbert Employee

func main() {

	dilbert.Salary = 5000
	position := &dilbert.Position
	*position = "Senior" + *position
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Println(EmployeeByID(dilbert.ManagerID).ID)

}
func EmployeeByID(id int) *Employee { return EmployeeByID(dilbert.ManagerID) }
