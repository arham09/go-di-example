// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"database/sql"
	"github.com/arham09/go-di-example/service"
	"github.com/arham09/go-di-example/service/supplier/delivery/http"
	"github.com/arham09/go-di-example/service/supplier/repository"
	"github.com/arham09/go-di-example/service/supplier/usecase"
	"time"
)

// Injectors from wire.go:

func StartingService(conn *sql.DB, timeout time.Duration) service.Server {
	supplierRepository := repository.NewPgSupplierRepository(conn)
	supplierUsecase := usecase.NewSupplierUsecase(supplierRepository, timeout)
	handler := http.NewSupplierHandler(supplierUsecase)
	server := service.NewServer(handler)
	return server
}
