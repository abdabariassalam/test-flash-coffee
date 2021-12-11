package http

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	Handler httpHandler
}

func NewRouter(handler httpHandler) *Router {
	return &Router{
		router:  gin.Default(),
		Handler: handler,
	}
}

func (r *Router) routes() {
	r.router.GET("/recipes", r.Handler.GetRecipe)
	r.router.POST("/recipes", r.Handler.CreateRecipe)
	r.router.GET("/ingredient", r.Handler.GetIngredients)
	r.router.POST("/ingredient", r.Handler.CreateIngredient)
	r.router.POST("/step", r.Handler.CreateStep)
}
