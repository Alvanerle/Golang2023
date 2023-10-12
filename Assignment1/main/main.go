package main

import (
	"Assignment1/engineer"
	"Assignment1/hr"
	"Assignment1/manager"
	"Assignment1/sales"
	"Assignment1/teacher"
	"fmt"
)

func main() {
	manager := manager.NewManager("Manager", 60000.0, "123 Main St")
	engineer := engineer.NewEngineer("Engineer", 80000.0, "456 Elm St")
	sales := sales.NewSales("Sales", 55000.0, "789 Oak St")
	hr := hr.NewHR("HR", 50000.0, "101 Pine St")
	teacher := teacher.NewTeacher("Marketing", 55000.0, "202 Maple St")

	fmt.Println("Manager:", manager.GetPosition(), manager.GetSalary(), manager.GetAddress())
	fmt.Println("Engineer:", engineer.GetPosition(), engineer.GetSalary(), engineer.GetAddress())
	fmt.Println("Sales:", sales.GetPosition(), sales.GetSalary(), sales.GetAddress())
	fmt.Println("HR:", hr.GetPosition(), hr.GetSalary(), hr.GetAddress())
	fmt.Println("Teacher:", teacher.GetPosition(), teacher.GetSalary(), teacher.GetAddress())
}
