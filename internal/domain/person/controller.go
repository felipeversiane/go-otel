package person

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/go-otel/internal/infra/config/validation"
	"github.com/gin-gonic/gin"
)

type personController struct {
	service PersonServiceInterface
}

type PersonControllerInterface interface {
	InsertOneController(c *gin.Context)
	GetOneController(c *gin.Context)
	GetAllController(c *gin.Context)
	UpdateController(c *gin.Context)
	DeleteController(c *gin.Context)
}

func NewPersonController(service PersonServiceInterface) PersonControllerInterface {
	return &personController{service}
}

func (controller *personController) InsertOneController(c *gin.Context) {
	var req PersonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationError := validation.ValidateError(err)
		c.JSON(validationError.Code, validationError)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := controller.service.InsertOneService(req, ctxTimeout)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})

}

func (controller *personController) GetOneController(c *gin.Context) {

}
func (controller *personController) GetAllController(c *gin.Context) {

}

func (controller *personController) UpdateController(c *gin.Context) {

}

func (controller *personController) DeleteController(c *gin.Context) {

}
