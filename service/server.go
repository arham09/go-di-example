package service

import (
	"log"

	"github.com/arham09/go-di-example/service/supplier"
	"github.com/labstack/echo"
)

type Server struct {
	delivery supplier.Handler
}

func NewServer(h supplier.Handler) Server {
	return Server{delivery: h}
}

func (s *Server) Run() {
	e := echo.New()

	e.GET("/v1/supplier", s.delivery.FetchAll)

	log.Fatal(e.Start(":2021"))
}
