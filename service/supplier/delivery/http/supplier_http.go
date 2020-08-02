package http

import (
	"context"
	"net/http"

	"github.com/arham09/go-di-example/service/helpers"
	"github.com/arham09/go-di-example/service/supplier"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type SupplierHandler struct {
	SUsecase supplier.Usecase
}

func NewSupplierHandler(us supplier.Usecase) *SupplierHandler {
	return &SupplierHandler{SUsecase: us}
}

func (s *SupplierHandler) FetchAll(c echo.Context) error {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	listSup, err := s.SUsecase.FetchAll(ctx)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listSup)
}
