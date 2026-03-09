package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.Engine,
	userHandlers *UserHandlers,
	taskHandlers *TaskHandlers,
) {

	r.POST("/register", userHandlers.Register)
	r.POST("/login", userHandlers.Login)

	r.POST("/tasks", taskHandlers.Create)
	r.PUT("/tasks/:id", taskHandlers.Update)
}
