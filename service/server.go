package service

import (
	"log"

	"github.com/arham09/go-di-example/service/supplier/delivery/http"
	"github.com/labstack/echo"
)

type Server struct {
	supplierDelivery http.SupplierHandler
}

func NewServer(h http.SupplierHandler) *Server {
	return &Server{supplierDelivery: h}
}

func (s *Server) Run() {
	e := echo.New()

	e.GET("/v1/supplier", s.supplierDelivery.FetchAll)

	log.Fatal(e.Start(":2021"))
}
