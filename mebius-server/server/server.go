package server

import (
	"github.com/PIETEP/mebius/mebius-server/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	u := r.Group("/users")
	{
		ctrl := controller.UserController{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	api := r.Group("/api/v1")

	c := api.Group("/cards")
	{
		ctrl := controller.CardController{}
		c.GET("", ctrl.Index)
		c.POST("", ctrl.Create)
		c.GET("/:id", ctrl.Show)
		c.PUT("/:id", ctrl.Update)
		c.DELETE("/:id", ctrl.Delete)
	}

	d := api.Group("/decks")
	{
		ctrl := controller.DeckController{}
		d.GET("", ctrl.Index)
		d.POST("", ctrl.Create)
		d.GET("/:id", ctrl.Show)
		d.PUT("/:id", ctrl.Update)
		d.DELETE("/:id", ctrl.Delete)
	}

	return r
}
