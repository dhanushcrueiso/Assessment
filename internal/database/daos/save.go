package daos

import (
	"Assessment/internal/db"
	"Assessment/internal/dtos"

	"github.com/google/uuid"
)

func SaveDetails(req dtos.EmployeeDetails) error {
	err := db.DB.Unscoped().Table("employee").Create(req).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDetails(id uuid.UUID) ([]dtos.EmployeeDetails, error) {
	var employee []dtos.EmployeeDetails
	err := db.DB.Unscoped().Table("employee").Find(&employee).Where("id =?", id).Order("clockin DESC").Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func GetEmployeesWithLess(time int) ([]dtos.EmployeeDetails, error) {
	var employee []dtos.EmployeeDetails
	err := db.DB.Unscoped().Raw("select id,clock_in,clock_out,total_hours from employee where total_hours > ?", time).Find(&employee).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}
