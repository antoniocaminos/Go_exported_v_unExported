package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "Jhon", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jhonnes", Salary: 50000, FullTime: true},
	{FirstName: "Marck", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: true},
	{FirstName: "Magaret", LastName: "Carter", Salary: 100000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	myStaff.All()
	//log.Println(myStaff.All())
	log.Println(myStaff.Overpaid())
	log.Println(myStaff.Underpaid())
}
