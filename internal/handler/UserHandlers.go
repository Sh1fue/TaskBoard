package handler

import (
	"net/http"
	"trello_parody/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	AuthService *service.AuthService
}

func NewUserHandlers(authService *service.AuthService) *UserHandlers {
	return &UserHandlers{
		AuthService: authService,
	}
}

func (h *UserHandlers) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user, token, err := h.AuthService.CreateUser(
		c.Request.Context(),
		req.Username,
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}
func (h *UserHandlers) Login(c *gin.Context){
	var req struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err:= c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}
	user,token,err:= h.AuthService.LoginUser(
		c.Request.Context(),
		req.Email,
		req.Password,
	)
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"user": user,
		"token": token,
	})

}	