package src

import (
	v1 "my-echo/src/v1"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo world!")
	})

	e.GET("health", func(c echo.Context) error {
		return c.String(http.StatusOK, "I am alive!")
	})

	// Resources Routes

	UserRouter(e.Group("api/v1/users"))

}

func UserRouter(g *echo.Group) {
	user := v1.User{}
	g.GET("", user.FindAll)
	g.POST("", user.AddUser)
	g.GET("/:id", user.FindById)
}
