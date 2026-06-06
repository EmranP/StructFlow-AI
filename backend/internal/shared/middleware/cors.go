package middleware

import (
	"github.com/EmranP/Design-Struct-Project-AI/backend/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS(cfg *configs.Config) fiber.Handler {
	return cors.New(
		cors.Config{
			AllowOrigins: cfg.OriginUrl,

			AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",

			AllowHeaders: "Origin,Content-Type,Accept,Authorization",

			ExposeHeaders: "Content-Length",

			AllowCredentials: false,
		},
	)
}
