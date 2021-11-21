package http

import (
	"net/http"

	"github.com/bariasabda/test-flash-coffee/domain/service"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	svc service.Service
}

func NewHandler(service service.Service) *httpHandler {
	return &httpHandler{
		svc: service,
	}
}

func (h *httpHandler) GetRecipe(c *gin.Context) {
	recipes, err := h.svc.Recipe.GetRecipes()
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, recipes)
}
