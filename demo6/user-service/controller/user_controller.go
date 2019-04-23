package controller

import (
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/model"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserController() UserController {
	return &userController{
		userService: service.NewUserService(),
	}
}

type userController struct {
	userService service.UserService
}

func (this *userController)GetUsers(c *gin.Context)  {
	users, err := this.userService.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (this *userController)AddUser(c *gin.Context)  {
	u := &model.User{}
	err := c.ShouldBindJSON(u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = this.userService.SaveUser(u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (this *userController)GetUser(c *gin.Context)  {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u, err := this.userService.FindUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (this *userController)UpdateUser(c *gin.Context)  {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u := &model.User{}
	err = c.ShouldBindJSON(u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u, err = this.userService.UpdateUser(userId, u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (this *userController)DeleteUser(c *gin.Context)  {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = this.userService.DeleteById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, userId)
}


