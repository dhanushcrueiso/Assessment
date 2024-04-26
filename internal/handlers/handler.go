package handlers

import (
	"Assessment/internal/database/daos"
	"Assessment/internal/dtos"
	"Assessment/internal/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func SaveAttendance(c *fiber.Ctx) {

	fmt.Println("check log ")
	var emp dtos.Attendance

	if err := c.BodyParser(&emp); err != nil {
		return
	}
	fmt.Println(emp)
	err := services.SaveAttendance(emp)
	if err != nil {
		return
	}

	c.JSON("successfully added details")

}

func GetAttendance(c *fiber.Ctx) {
	fmt.Println("check handler layer")

	eid := c.Query("eid")

	employeeid, err := uuid.Parse(eid)
	if err != nil {
		fmt.Println("unable to parse id")
	}

	res, err := services.GetAttendance(employeeid)
	if err != nil {
		return
	}

	c.JSON(&res)

}

func GetAllWithLess(c *fiber.Ctx) {

	t, err := strconv.Atoi(c.Query("time"))
	if err != nil {
		fmt.Println("unable to parse string to int")
	}
	res, err := daos.GetEmployeesWithLess(t)
	if err != nil {
		return
	}

	c.JSON(res)
}
