package supplier

import (
	"github.com/labstack/echo"
)

type Handler interface {
	FetchAll(e echo.Context) error
}
