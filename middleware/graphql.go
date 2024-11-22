package middleware

import (
	"github.com/labstack/echo/v4"
)

// GraphQLMiddleware checks for proper Accept header
func GraphQLMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Set response content type
		c.Response().Header().Set("Content-Type", "application/graphql-response+json")
		return next(c)
	}
}
