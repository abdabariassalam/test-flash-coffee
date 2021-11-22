package http

import (
	"encoding/json"
	"net/http"

	"github.com/bariasabda/test-flash-coffee/domain/entity"
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

func (h *httpHandler) CreateRecipe(c *gin.Context) {
	var req entity.CreateRecipe
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	// if req.Name == "" || req.Description == "" || req.RecipeCategory.Name == "" || req.User.Name == "" {
	// 	c.JSON(http.StatusBadRequest, "All input is required")
	// 	return
	// }
	recipe, err := h.svc.Recipe.CreateRecipe(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success create",
		"data":    recipe,
	})
}
