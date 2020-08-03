package manual

import (
	"fmt"

	"github.com/arham09/go-di-example/service"
	"github.com/arham09/go-di-example/service/configs/database"
	sh "github.com/arham09/go-di-example/service/supplier/delivery/http"
	sp "github.com/arham09/go-di-example/service/supplier/repository"
	su "github.com/arham09/go-di-example/service/supplier/usecase"
)

func Main() {
	dbHost := "localhost"
	dbUser := "medea"
	dbPassword := "developer"
	dbName := "battlefield"

	dsn := fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, dbUser, dbPassword, dbHost, dbName)

	db, err := database.NewDB(dsn)

	if err != nil {
		panic(err)
	}

	supplierRepository := sp.NewPgSupplierRepository(db)
	supplierUsecase := su.NewSupplierUsecase(supplierRepository)
	supplierHandler := sh.NewSupplierHandler(supplierUsecase)

	server := service.NewServer(supplierHandler)

	server.Run()
}
