package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/application/domain"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
)

// Config
type Config struct {
	Service application.Service
	Logger  *domain.Logger
}

// Controller
type Controller struct {
	service application.Service
	logger  *domain.Logger
}

// NewController
func NewController(c *Config) *Controller {
	return &Controller{
		service: c.Service,
		logger:  c.Logger,
	}
}

// Create
func (c *Controller) Create(resource *JSONRequestBody) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*domain.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		if err := ctx.BodyParser(resource); err != nil {
			fmt.Printf("Error in BodyParser %+v\n", err)
			return err
		}
		fmt.Printf("Resource in Create Controller: %+v\n", resource)

		model := resource.Data.Properties
		fmt.Printf("Model in Create Controller: %+v\n", model)

		result, err := c.service.Create(model)
		if err != nil {
			return err
		}
		fmt.Printf("Result in Create Controller: %+v\n", result)

		ctx.Status(http.StatusCreated)
		ctx.JSON(result)
		return nil
	}
}

// Delete
func (c *Controller) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*domain.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Delete Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		if err := c.service.Delete(id); err != nil {
			return err
		}
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

// Detail
func (c *Controller) Detail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*domain.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Detail Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		result, err := c.service.Detail(id)
		if err != nil {
			return err
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}

// List
func (c *Controller) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*domain.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("List Controller called")

		// TODO: get/bind/validate query from request

		query := &domain.ListMeta{} // c.getQueryData(ctx)
		result, err := c.service.List(query)
		if err != nil {
			return err
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}

// Update
func (c *Controller) Update(resource *JSONRequestBody) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*domain.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// TODO: validate body

		if err := ctx.BodyParser(resource); err != nil {
			fmt.Printf("Error in BodyParser %+v\n", err)
			return err
		}
		fmt.Printf("Resource in Update Controller: %+v\n", resource)

		model := resource.Data.Properties // TODO: problem here with ID
		fmt.Printf("Model in Update Controller: %+v\n", model)

		result, err := c.service.Update(model)
		if err != nil {
			return err
		}
		fmt.Printf("Result in Update Controller: %+v\n", result)

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}
