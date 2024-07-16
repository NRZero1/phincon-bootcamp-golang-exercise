package service

import (
	"employee_attendance/model"
	"fmt"
	"time"
)

func AddAttendance(listOfEmployee *[]model.Employee) {
	var employeeId int
	var present string
	for {
		fmt.Print("\nPlease enter the id of the employee you want to add the attendance: ")
		_, errEmployeeId := fmt.Scanf("%d\n", &employeeId)

		if errEmployeeId != nil {
			fmt.Println("Incorrect input received, please enter a number")
			continue
		}

		var foundEmployee *model.Employee

		for index := range *listOfEmployee {
			if (*listOfEmployee)[index].Id == employeeId {
				foundEmployee = &(*listOfEmployee)[index]
				break
			}
		}

		if foundEmployee == nil {
			fmt.Println("Employee Not Found")
			continue
		}

		fmt.Print("Please enter the attendance: (present/absent) ")
		_, errAttendance := fmt.Scanf("%s\n", &present)

		if errAttendance != nil || (present != "present" && present != "absent") {
			fmt.Println("Invalid input, please enter the correct input (present/absent)")
			continue
		}

		absent := present == "present"

		attendance := model.NewAttendanceBuilder().
			SetId(len(foundEmployee.Attendances) + 1).
			SetDate(time.Now()).
			SetAbsent(absent).
			SetEmployeeId(foundEmployee.Id).
			Build()

		foundEmployee.Attendances = append(foundEmployee.Attendances, *attendance)
		fmt.Print("\n")
		break
	}
}