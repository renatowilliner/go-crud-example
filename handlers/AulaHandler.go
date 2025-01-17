package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/dto"
	"github.com/maxilovera/go-crud-example/services"
)

type AulaHandler struct {
	aulaService services.AulaInterface
}

func NewAulaHandler(aulaService services.AulaInterface) *AulaHandler {
	return &AulaHandler{
		aulaService: aulaService,
	}
}

func (handler *AulaHandler) ObtenerAulas(c *gin.Context) {
	aulas := handler.aulaService.ObtenerAulas()

	c.JSON(http.StatusOK, aulas)
}

func (handler *AulaHandler) InsertarAula(c *gin.Context) {
	var aula dto.Aula

	if err := c.ShouldBindJSON(&aula); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.aulaService.InsertarAula(&aula)

	c.JSON(http.StatusCreated, resultado)
}
