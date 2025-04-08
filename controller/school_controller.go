package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"school-service/service"
	"strconv"
)

type SchoolController struct {
	schoolService service.ISchoolService
}

func NewSchoolController(schoolService service.ISchoolService) *SchoolController {
	return &SchoolController{schoolService: schoolService}
}

func (schoolController *SchoolController) RegisterRoutes(echo *echo.Echo) {
	echo.GET("/api/v1/school:id", schoolController.GetById)
}

func (schoolController *SchoolController) GetById(c echo.Context) error {
	paramId := c.Param("id")
	schoolId, _ := strconv.Atoi(paramId)
	response, err := schoolController.schoolService.GetById(int64(schoolId))
	if err == nil {
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusNotFound, err.Error())
}
