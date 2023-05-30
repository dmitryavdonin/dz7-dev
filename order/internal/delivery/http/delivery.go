package delivery

import (
	"fmt"
	"order/internal/service"

	"github.com/dmitryavdonin/gtools/logger"
	"github.com/gin-gonic/gin"
)

type Delivery struct {
	services *service.Service
	router   *gin.Engine
	logger   logger.Interface
	port     int
	options  Options
}

type Options struct{}

func New(services *service.Service, port int, logger logger.Interface, options Options) (*Delivery, error) {

	var d = &Delivery{
		services: services,
		logger:   logger,
		port:     port,
	}

	d.SetOptions(options)

	d.router = d.initRouter()
	return d, nil
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf(":%d", d.port))
}
