package main

import (
	"employee_attendance/model"
	"employee_attendance/service"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() error

func init() {
	clear = make(map[string]func() error)
	clear["linux"] = func() error {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	clear["windows"] = func() error {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	clear["darwin"] = clear["linux"]
}

func main() {
	var menu int
	var listOfEmployee []model.Employee

	for {
		fmt.Println("\nEMPLOYEE ATTENDANCE MANAGEMENT SYSTEM")
		fmt.Println("==================================================")
		fmt.Println("Menu:")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Add Attendance")
		fmt.Println("3. Show All Employees and Attendances")
		fmt.Println("4. Delete an Employee")
		fmt.Println("5. Find Employee by Id")
		fmt.Println("6. Clear Console")
		fmt.Println("7. Exit")
		fmt.Print("\nInput menu: ")
		_, err := fmt.Scanf("%d\n", &menu)

		if (err != nil) {
			fmt.Println("Invalid input detected, please enter a number")
			continue
		}

		switch menu {
		case 1:
			service.AddEmployee(&listOfEmployee)
		case 2:
			service.AddAttendance(&listOfEmployee)
		case 3:
			service.ShowAll(&listOfEmployee)
		case 4:
			var id int
			fmt.Print("Please enter the employee id you wish to delete: ")
			_, errId := fmt.Scanf("%d\n", &id)

			if errId != nil {
				fmt.Println("Invalid input, please input a number")
				continue
			}
			result := service.DeleteEmployeeById(&listOfEmployee, id)

			if (result == 1) {
				fmt.Println("Delete success")
			} else {
				fmt.Println(result)
			}
		case 5:
			var id int
			fmt.Print("Please enter the employee id you wish to delete: ")
			_, errId := fmt.Scanf("%d\n", &id)

			if errId != nil {
				fmt.Println("Invalid input, please input a number")
				continue
			}

			service.FindEmployeeById(&listOfEmployee, id)
		case 6:
			if clearFunc, ok := clear[runtime.GOOS]; ok {
				err := clearFunc()
				if err != nil {
					fmt.Println("Failed to clear console:", err)
				}
			} else {
				fmt.Println("Platform is not supported")
			}
		case 7:
			fmt.Println("Press any key to exit...")
			fmt.Scanln()
			return
		default:
			fmt.Println("Invalid menu inputted, please try again")
		}
	}
}