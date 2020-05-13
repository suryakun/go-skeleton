package middleware

import "github.com/labstack/echo/v4"

// GoMiddleware ...
type GoMiddleware struct {
}

// CORS ...
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// InitMiddleware ...
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
