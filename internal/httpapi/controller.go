package httpapi

import "github.com/gin-gonic/gin"

type Controller struct{}

func newController() *Controller {
	return &Controller{}
}

func (c *Controller) Create(ctx *gin.Context) {
	// get correlation, type from c.state
	// get body, method from request
	// validate body
	// get properties data from body, send props, type to domain layer, return result
	// serialize result
	// set body, http status, mime type (at a minimum)
}

func (c *Controller) Delete(ctx *gin.Context) {
	// get correlation, type from c.state
	// get id from path params
	// send type, id to domain layer, return result
	// set http status (at a minimum)
}

func (c *Controller) Detail(ctx *gin.Context) {
	// get correlation, type from c.state
	// get id from path params
	// send type, id to domain layer, return result
	// serialize result
	// set body, http status, mime type (at a minimum)
}

func (c *Controller) List(ctx *gin.Context) {
	// get correlation, type from c.state
	// get query from request
	// parse query ??
	// validate query, get filters, page, sort data
	// send filters, page, sort, type to domain layer, return result
	// serialize result
	// set body, http status, mime type (at a minimum)
}

func (c *Controller) Update(ctx *gin.Context) {
	// get correlation, type from c.state
	// get id from request
	// get body, method from request
	// validate body, id based on method, type
	// get properties data from body, send props, type, id to domain layer, return result
	// serialize result
	// set body, http status, mime type (at a minimum)
}