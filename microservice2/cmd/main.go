package main

func main() {
	// // Инициализация логгера
	// // logger := log.New(os.Stdout, "", log.LstdFlags)

	// client := database.SetupDatabase()
	// defer client.Disconnect(context.Background())

	// studentRepo, _ := repository.NewCoursesRepository(client, "taskdb", "student")

	// studentUsecase := usecase.NewStudentUsecase(*studentRepo)

	// studentHandler := handler.NewStudentHandler(studentUsecase)

	// router := gin.Default()

	// // Регистрация маршрутов
	// api := router.Group("/api/")
	// {
	// 	api.POST("/students", studentHandler.CreateStudent)
	// 	api.GET("/students/:id", studentHandler.GetStudentByID)

	// 	auth := api.Group("/auth/")
	// 	{
	// 		auth.POST("/sign-up", studentHandler.CreateStudent)
	// 		auth.POST("/sign-in", studentHandler.SignIn)
	// 	}
	// }
	// // router.PUT("/students/:id", studentHandler.UpdateStudent)
	// // router.DELETE("/students/:id", studentHandler.DeleteStudent)

	// // Запуск HTTP-сервера
	// router.Run(":8080")
}