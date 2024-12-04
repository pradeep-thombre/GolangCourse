package main

import (
	"GolangCourse/apis"
	_ "GolangCourse/apis/docs"
	"GolangCourse/commons/apploggers"
	"GolangCourse/configs"
	"GolangCourse/internals/db"
	"GolangCourse/internals/services"
	"context"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @User Management
// @version 1.0
// @description This is a sample API using Echo and Swagger.
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:3000
// @BasePath /
func main() {
	context, logger := apploggers.NewLoggerWithCorrelationid(context.Background(), "")
	err := configs.NewApplicationConfig(context)
	if err != nil {
		logger.Errorf("Error in Appconfig:", err)
	}

	// Initialize services
	userDbService := db.NewUserDbService(configs.AppConfig.DbClient)
	userEventService := services.NewUserEventService(userDbService)

	topicDbService := db.NewTopicDbService(configs.AppConfig.DbClient)
	topicEventService := services.NewTopicEventService(topicDbService)

	pageDbService := db.NewPageDbService(configs.AppConfig.DbClient)
	pageEventService := services.NewPageEventService(pageDbService)

	// Echo instance
	e := echo.New()

	// user api Routes
	userController := apis.NewUserController(userEventService)
	e.GET("/users", userController.GetUsers)
	e.GET("/users/:id", userController.GetUserById)
	e.DELETE("/users/:id", userController.DeleteUserById)
	e.POST("/users", userController.CreateUser)
	e.PATCH("/users/:id", userController.UpdateUser)

	topicController := apis.NewTopicController(topicEventService)
	e.GET("/topic", topicController.GetAllTopics)
	e.POST("/topic", topicController.CreateTopic)
	e.PUT("/topic/:id", topicController.UpdateTopic)
	e.DELETE("/topic/:id", topicController.DeleteTopic)
	e.PUT("/topic/:id/hide", topicController.HideTopic)

	pageController := apis.NewPageController(pageEventService)
	e.GET("/topic/:topicId/pages", pageController.GetPagesByTopicId)
	e.POST("/topic/:topicId/pages", pageController.CreatePage)
	e.GET("/topic/:topicId/pages/:id", pageController.GetPageById)
	e.PATCH("/topic/:topicId/pages/:id", pageController.UpdatePage)
	e.PATCH("/topic/:topicId/pages/:id/hide", pageController.HidePage)
	e.DELETE("/topic/:topicId/pages/:id", pageController.DeletePageById)

	// Swagger UI route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	logger.Infof("starting http server on localhost:%v", configs.AppConfig.HttpPort)
	e.Logger.Fatal(e.Start(":" + configs.AppConfig.HttpPort))
}
