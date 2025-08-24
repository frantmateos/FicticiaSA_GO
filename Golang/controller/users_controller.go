package usersController

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	Domain "Golang/domain"

	middle "Golang/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type UserService interface {
	InsertUsuario(usuarioDomain Domain.UserData) (Domain.UserData, error)
	GetUserByName(usuarioDomain Domain.UserData) (Domain.UserData, error)
	UpdateUser(usuarioDomain Domain.UserData) (Domain.UserData, error)
	Login(User Domain.UserData) (Domain.LoginData, error)
	GetAllUsers() ([]Domain.UserData, error)
	GetUserById(userId int) (Domain.UserData, error)
}

type Controller struct {
	service UserService
}

func NewController(service UserService) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) Login(c *gin.Context) {
	var userData Domain.UserData
	c.BindJSON(&userData)

	loginResponse, err := controller.service.Login(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al procesar la solicitud",
		})

		return

	}
	c.JSON(http.StatusOK, loginResponse)

}
func (controller Controller) Extrac(c *gin.Context) {

	data := strings.TrimSpace(c.GetHeader("Authorization"))
	log.Println("token buscado: ", data)
	response, err := middle.ExtractClaims(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al procesar la solicitud",
		})

		return

	}
	c.JSON(http.StatusOK, response)

}

func (controller Controller) GetUserByName(c *gin.Context) {
	fmt.Println("llego al controller")

	var userDomain Domain.UserData
	c.BindJSON(&userDomain)

	userDomain, err := controller.service.GetUserByName(userDomain)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al procesar la solicitud",
		})

		return
	}
	c.JSON(http.StatusOK, userDomain)

}

func (controller Controller) GetUserById(c *gin.Context) {
	userId := c.Param("id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := controller.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller Controller) GetAllUsers(c *gin.Context) {
	users, err := controller.service.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener la lista de usuarios",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (controller Controller) UsuarioInsert(c *gin.Context) {
	var userDomain Domain.UserData
	err := c.BindJSON(&userDomain)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userDomain, er := controller.service.InsertUsuario(userDomain)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al procesar la solicitud",
		})

		return
	}

	c.JSON(http.StatusCreated, userDomain)

}

func (controller Controller) UpdateUser(c *gin.Context) {
	var userDomain Domain.UserData
	err := c.BindJSON(&userDomain)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDomain, er := controller.service.UpdateUser(userDomain)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al procesar la solicitud",
		})

		return
	}

	c.JSON(http.StatusCreated, userDomain)

}
