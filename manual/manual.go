package manual

import (
	"fmt"
	"time"

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

	timeoutContext := time.Duration(2) * time.Second

	supplierRepository := sp.NewPgSupplierRepository(db)
	supplierUsecase := su.NewSupplierUsecase(supplierRepository, timeoutContext)
	supplierHandler := sh.NewSupplierHandler(supplierUsecase)

	server := service.NewServer(supplierHandler)

	server.Run()
}
