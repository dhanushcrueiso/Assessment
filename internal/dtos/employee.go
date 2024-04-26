package dtos

import "time"

type Attendance struct {
	Id       string    `json:"id"`
	ClockIn  time.Time `json:"clockin"`
	ClockOut time.Time `json:"clockout"`
}

type EmployeeDetails struct {
	Id         string    `json:"id"`
	ClockIn    time.Time `json:"clockin"`
	ClockOut   time.Time `json:"clockout"`
	TotalHours int64     `json:"total_hours"`
}
