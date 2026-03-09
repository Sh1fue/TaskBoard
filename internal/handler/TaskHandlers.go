package handler

import (
	"net/http"
	"strconv"
	"trello_parody/internal/domain"
	"trello_parody/internal/service"

	"github.com/gin-gonic/gin"
)


type TaskHandlers struct {
	TaskService *service.TaskService
}

func NewTaskHandlers(taskService *service.TaskService) *TaskHandlers {
	return &TaskHandlers{
		TaskService: taskService,
	}
}

func (t *TaskHandlers) Create(c *gin.Context){
	var task domain.Task

	if err:= c.ShouldBindJSON(&task); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	err:= t.TaskService.CreateTask(c.Request.Context(), &task)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated,task)
}

func (t *TaskHandlers) Update(c *gin.Context){
	idParam := c.Param("id")

	id,err:= strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task_id",
		})
		return
	}
	var task domain.Task
	if err:= c.ShouldBindJSON(&task); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H {
			"error":err.Error(),
		})
	}
	task.ID = id

	err = t.TaskService.UpdateTask(c.Request.Context(), &task)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task is updated",
	})
}