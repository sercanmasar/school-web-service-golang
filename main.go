package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"school-service/controller"
	"school-service/infrastructure"
	"school-service/persistence"
	"school-service/service"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	e := echo.New()
	ctx := context.Background()

	pool := infrastructure.GetConnectionPool(ctx, infrastructure.Config{
		Host:                  "localhost",
		Port:                  "5423",
		UserName:              "postgres",
		Password:              "postgres",
		DbName:                "school_service_db",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	repository := persistence.NewSchoolRepository(pool)
	schoolService := service.NewSchoolService(repository)
	schoolController := controller.NewSchoolController(schoolService)
	schoolController.RegisterRoutes(e)

	e.Start("localhost:7070")
}
