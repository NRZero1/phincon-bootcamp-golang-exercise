package service

import (
	"employee_attendance/model"
	"fmt"
)

func AddEmployee(listOfEmployee *[]model.Employee) {
	var name string
	listEmployeeLength := len(*listOfEmployee)
	for {
		fmt.Print("\nPlease input employee name's: ")
		_, err := fmt.Scanf("%s\n", &name)

		if (err != nil) {
			fmt.Println("Invalid input")
			continue
		}

		employee := model.NewEmployeeBuilder().
			SetId(listEmployeeLength + 1).
			SetName(name).
			Build()
		
		*listOfEmployee = append(*listOfEmployee, *employee)
		fmt.Print("\n")
		break
	}
}

func FindEmployeeById(listOfEmployee *[]model.Employee, id int) {
	for index := range *listOfEmployee {
		if (*listOfEmployee)[index].GetId() == id {
			fmt.Println("Employee Id: ", (*listOfEmployee)[index].GetId())
			fmt.Println("Employee Name: ", (*listOfEmployee)[index].GetName())
			for _, value := range (*listOfEmployee)[index].Attendances {
				fmt.Println("Id attendance: ", value.GetId())
				fmt.Println("Date: ", value.GetDate())
				fmt.Println("Status: ", value.GetAbsent())
				fmt.Println()
			}
		}
	}
}

func DeleteEmployeeById(listOfEmployee *[]model.Employee, id int) interface{} {
	employees := *listOfEmployee
	for index := range employees {
		if employees[index].Id == id {
			*listOfEmployee = append(employees[:index], employees[index+1:]...)
			return 1
		}
	}
	return fmt.Sprintf("No data found with employee id: %d", id)
}

func ShowAll(listOfEmployee *[]model.Employee) {
	for _, value := range *listOfEmployee {
		fmt.Print("\n")
		fmt.Println("Id Employee: ", value.GetId())
		fmt.Println("Name Employee: ", value.GetName())
		fmt.Println("Attendances: ")
		for _, value := range value.Attendances {
			fmt.Println("Id attendance: ", value.GetId())
			fmt.Println("Date: ", value.GetDate())
			fmt.Println("Status: ", value.GetAbsent())
			fmt.Println()
		}
	}
	fmt.Print("\n")
}