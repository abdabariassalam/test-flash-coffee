package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	if req.Name == "" || req.Description == "" || req.RecipeCategory.Name == "" || req.User.Name == "" {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
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

func (h *httpHandler) CreateIngredient(c *gin.Context) {
	var req entity.CreateIngredientRequest
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	defer file.Close()
	json.Unmarshal([]byte(c.Request.FormValue("raw")), &req)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	if req.Name == "" || req.Color == 0 || req.IngredientCategory.Name == "" || req.IngredientCategory.Description == "" {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
	ingredient, err := h.svc.Ingredient.CreateIngredient(entity.CreateIngredientRequest{
		Name:  req.Name,
		Color: req.Color,
		IngredientCategory: entity.IngredientCategory{
			Name:        req.IngredientCategory.Name,
			Description: req.IngredientCategory.Description,
		},
		Image: entity.File{
			ImageContentType: handler.Header.Get("Content-Type"),
			ImageName:        handler.Filename,
			ImageFile:        fileBytes,
			ImageSize:        handler.Size,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success create",
		"data":    ingredient,
	})
}

func (h *httpHandler) GetIngredients(c *gin.Context) {
	recipes, err := h.svc.Ingredient.GetIngredients()
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func (h *httpHandler) CreateStep(c *gin.Context) {
	var req entity.CreateStep
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	defer file.Close()
	json.Unmarshal([]byte(c.Request.FormValue("raw")), &req)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	fmt.Println(req)
	if req.StepNumber == 0 || req.Description == "" || req.Timer == 0 || req.StepIngredient.RecipeId == 0 ||
		req.StepIngredient.IngredientId == 0 || req.StepIngredient.Amount == 0 ||
		req.StepIngredient.Unit == "" {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
	step, err := h.svc.Step.CreateStep(entity.CreateStep{
		StepNumber:  req.StepNumber,
		Description: req.Description,
		Timer:       req.Timer,
		StepIngredient: entity.StepIngredient{
			RecipeId:     req.StepIngredient.RecipeId,
			IngredientId: req.StepIngredient.IngredientId,
			Amount:       req.StepIngredient.Amount,
			Unit:         req.StepIngredient.Unit,
		},
		Image: entity.File{
			ImageContentType: handler.Header.Get("Content-Type"),
			ImageName:        handler.Filename,
			ImageFile:        fileBytes,
			ImageSize:        handler.Size,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success create",
		"data":    step,
	})
}
