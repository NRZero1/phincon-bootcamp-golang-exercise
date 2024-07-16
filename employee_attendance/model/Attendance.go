package model

import (
	"slices"
	"time"
)

type Attendance struct {
	Id          int
	Employee_id int
	Date 		time.Time
	Absent  	bool
}

type AttendanceBuilder interface {
	SetId(int) AttendanceBuilder
	SetEmployeeId(int) AttendanceBuilder
	SetDate(time.Time) AttendanceBuilder
	SetAbsent(bool) AttendanceBuilder
	Build() *Attendance
}

type attendanceBuilder struct {
	attendance *Attendance
}

func NewAttendanceBuilder() AttendanceBuilder {
	return &attendanceBuilder{
		attendance: &Attendance{},
	}
}

func (ab *attendanceBuilder) SetId(id int) AttendanceBuilder {
	ab.attendance.Id = id
	return ab
}

func (ab *attendanceBuilder) SetEmployeeId(employee_id int) AttendanceBuilder {
	ab.attendance.Employee_id = employee_id
	return ab
}

func (ab *attendanceBuilder) SetAbsent(absent bool) AttendanceBuilder {
	ab.attendance.Absent = absent
	return ab
}

func (ab *attendanceBuilder) SetDate(date time.Time) AttendanceBuilder {
	ab.attendance.Date = date
	return ab
}

func (ab *attendanceBuilder) Build() *Attendance {
	return ab.attendance
}

func (a *Attendance) GetId() int {
	return a.Id
}

func (a *Attendance) GetEmployeeId() int {
	return a.Employee_id
}

func (a *Attendance) GetAbsent() string {
	if (a.Absent) {
		return "present"
	}
	return "absent"
}

func (a *Attendance) GetDate() time.Time {
	return a.Date
}

func GetAttendanceById(listAttendance []Attendance, id int) (Attendance, bool) {
	index := slices.IndexFunc(listAttendance, func(a Attendance) bool {
		return a.Id == id
	})

	if index == -1 {
		return Attendance{}, false
	}

	return listAttendance[index], true
}