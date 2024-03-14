package main

import (
	"net/http"

	"github.com/adrianlarion/templ_echo_app/internal/view/page"
	"github.com/labstack/echo/v4"
)

func (app *application) home(c echo.Context) error{
	return render(c, http.StatusOK, page.Home("Frodo Baggins"))

}
