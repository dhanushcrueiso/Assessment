package services

import (
	"Assessment/internal/database/daos"
	"Assessment/internal/dtos"
	"fmt"

	"github.com/google/uuid"
)

func SaveAttendance(req dtos.Attendance) error {

	employee := dtos.EmployeeDetails{
		Id:         req.Id,
		ClockIn:    req.ClockIn,
		ClockOut:   req.ClockOut,
		TotalHours: int64(req.ClockOut.Sub(req.ClockIn)),
	}
	err := daos.SaveDetails(employee)
	if err != nil {
		return err
	}
	return nil
}

func GetAttendance(id uuid.UUID) ([]dtos.EmployeeDetails, error) {

	var employees []dtos.EmployeeDetails
	res, err := daos.GetDetails(id)
	if err != nil {
		return res, err
	}

	for _, v := range res {

		employees = append(employees, v)

	}
	fmt.Println(employees)

	return employees, nil
}
