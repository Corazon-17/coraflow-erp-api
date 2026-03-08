package main

import (
	"log"

	db "coraflow-erp-api/services/hr-service/db/sqlc"
	"coraflow-erp-api/services/hr-service/internal/grpc"
	"coraflow-erp-api/services/hr-service/internal/grpc/handler"
	"coraflow-erp-api/services/hr-service/internal/repository"
	"coraflow-erp-api/services/hr-service/internal/service"
	"coraflow-erp-api/shared/config"
	"coraflow-erp-api/shared/database"
)

func main() {
	cfg := config.Load()

	log.Println("starting hr-service")

	pool, err := database.NewPostgres(cfg.HRDBUrl)
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(pool)

	departmentRepo := repository.NewDepartmentRepository(q)
	employeeRepo := repository.NewEmployeeRepository(q)

	departmentService := service.NewDepartmentService(departmentRepo)
	employeeService := service.NewEmployeeService(employeeRepo)

	departmentHandler := handler.NewDepartmentHandler(departmentService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	err = grpc.Start(cfg.HRServicePort, departmentHandler, employeeHandler)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg

}
