package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mulaydm10/backend/internal/api/routes"
	"github.com/mulaydm10/backend/internal/middleware"

	_ "github.com/mulaydm10/backend/docs" // import generated docs
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Fiber Tools API
// @version 1.0
// @description API documentation for your tool-based Go backend.
// @host localhost:3000
// @BasePath /api
func main() {
	app := fiber.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	// Routes
	api := app.Group("/api")
	routes.RegisterToolRoutes(api)

	// Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Println("Server running at http://localhost:8002")
	log.Fatal(app.Listen(":8002"))
}
