package main

import (
	"context"
	handler "task-backend/microservice1/internal/app/handlers"
	"task-backend/microservice1/internal/app/repository"
	"task-backend/microservice1/internal/app/usecase"
	"task-backend/microservice1/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация логгера
	// logger := log.New(os.Stdout, "", log.LstdFlags)

	client := database.SetupDatabase()
	defer client.Disconnect(context.Background())

	studentRepo, _ := repository.NewStudentRepository(client, "taskdb", "students")

	studentUsecase := usecase.NewStudentUsecase(*studentRepo)

	studentHandler := handler.NewStudentHandler(studentUsecase)

	router := gin.Default()

	// Регистрация маршрутов
	api := router.Group("/api/")
	{
		api.POST("/students", studentHandler.CreateStudent)
		api.GET("/students/:id", studentHandler.GetStudentByID)

		auth := api.Group("/auth/")
		{
			auth.POST("/sign-up", studentHandler.CreateStudent)
			auth.POST("/sign-in", studentHandler.SignIn)
		}
	}
	// router.PUT("/students/:id", studentHandler.UpdateStudent)
	// router.DELETE("/students/:id", studentHandler.DeleteStudent)

	// Запуск HTTP-сервера
	router.Run(":8080")
}
