package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Salary    float64
}

type EmployeeManager struct {
	employees []Employee
}

func (em *EmployeeManager) AddEmployee(employee Employee) {
	em.employees = append(em.employees, employee)
}

func (em *EmployeeManager) RemoveEmployee(id int) {
	for i, emp := range em.employees {
		if emp.ID == id {
			em.employees = append(em.employees[:i], em.employees[i+1:]...)
			break
		}
	}
}

func (em *EmployeeManager) UpdateEmployee(id int, newEmployee Employee) {
	for i, emp := range em.employees {
		if emp.ID == id {
			em.employees[i] = newEmployee
			break
		}
	}
}

func (em *EmployeeManager) DisplayEmployees() {
	fmt.Println("Employee Details:")
	for _, emp := range em.employees {
		fmt.Printf("ID: %d, Name: %s %s, Age: %d, Salary: %.2f\n", emp.ID, emp.FirstName, emp.LastName, emp.Age, emp.Salary)
	}
}

func main() {
	employeeManager := EmployeeManager{}

	// Adding employees
	employeeManager.AddEmployee(Employee{ID: 1, FirstName: "John", LastName: "Doe", Age: 30, Salary: 50000.0})
	employeeManager.AddEmployee(Employee{ID: 2, FirstName: "Jane", LastName: "Smith", Age: 25, Salary: 45000.0})
	employeeManager.AddEmployee(Employee{ID: 3, FirstName: "Michael", LastName: "Johnson", Age: 35, Salary: 60000.0})

	// Displaying employees
	employeeManager.DisplayEmployees()

	// Removing an employee
	employeeManager.RemoveEmployee(2)
	fmt.Println("After removing employee with ID 2:")
	employeeManager.DisplayEmployees()

	// Updating an employee
	employeeManager.UpdateEmployee(1, Employee{ID: 1, FirstName: "John", LastName: "Doe", Age: 32, Salary: 55000.0})
	fmt.Println("After updating employee with ID 1:")
	employeeManager.DisplayEmployees()
}
