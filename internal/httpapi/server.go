package httpapi

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/application/domain"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// Config defines the input to NewServer
type Config struct {
	Application *application.Application `validate:"required"`
	BaseURL     string                   `validate:"required"`
	Logger      *domain.Logger           `validate:"required"`
	Mode        string                   `validate:"required"`
	Namespace   string                   `validate:"required"`
	Port        uint                     `validate:"required"`
}

// Server defines a server for handling HTTP API requests
type Server struct {
	App         *fiber.App
	Logger      *domain.Logger
	baseURL     string
	controllers *controllerRegistry
	namespace   string
	port        uint
}

// NewServer returns a new Server instance
func NewServer(c *Config) (*Server, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	app := fiber.New(fiber.Config{AppName: "domain"})

	log := c.Logger.Log.With().Str("tags", "httpapi").Logger()
	logger := &domain.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	controllers := registerControllers(logger, c.Application.Services)

	s := &Server{
		Logger: logger,
		App:    app,
		// baseURL:    c.BaseURL,
		controllers: controllers,
		namespace:   c.Namespace,
		port:        c.Port,
	}

	s.configureMiddleware()
	s.registerRoutes()

	return s, nil
}

// Serve starts the HTTP server on the configured address
func (s *Server) Serve() {
	addr := s.baseURL + ":" + strconv.FormatUint(uint64(s.port), 10)
	s.App.Listen(addr)
}
