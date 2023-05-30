package delivery

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (d *Delivery) initRouter() *gin.Engine {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	var router = gin.Default()

	router.Use(cors.New(corsConfig))

	//var router = gin.New()
	router.GET("/idempotency-key", d.GetIdempotencyKey)
	router.POST("/order", d.CreateOrder)
	d.routerOrder(router.Group("/order"))

	return router
}

func (d *Delivery) routerOrder(router *gin.RouterGroup) {

	router.GET("/:id", d.ReadOrderById)
	router.PUT("/:id", d.UpdateOrder)
	router.DELETE("/:id", d.DeleteOrderById)
}
