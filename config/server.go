package config

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	InitEnvironment()

	isConnected := InitDBConnection()

	if !isConnected {
		log.Fatalf("Error in database connection!")
	}

	go DBHealthChecker()

	echoInstance := echo.New()

	// // Configuring Middleware Logger
	echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Skipping logging for health checking api
		Skipper: func(c echo.Context) bool {
			if c.Request().RequestURI == "/health" {
				return true
			}
			return false
		},
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	echoInstance.Use(middleware.Recover())

	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	return echoInstance
}
