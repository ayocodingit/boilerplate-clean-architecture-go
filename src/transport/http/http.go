package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	App      *echo.Echo
	Response map[string]interface{}
}

func NewHttp() App {
	r := echo.New()

	plugins(r)
	homepage(r)

	return App{
		App: r,
	}
}

func plugins(r *echo.Echo) {
	r.Use(middleware.CORS())
	r.Use(middleware.Gzip())
	r.Use(middleware.Decompress())
	r.Use(middleware.Secure())
	r.Use(middleware.Recover())
}

func homepage(r *echo.Echo) {
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello, world!",
		})
	})
}

func (app App) Run(port int64) {
	if err := app.App.Start(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server http is running at http://localhost:%d", port)
}
