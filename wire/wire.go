package wire

import (
	"database/sql"
	"time"

	"github.com/arham09/go-di-example/service"
	sh "github.com/arham09/go-di-example/service/supplier/delivery/http"
	sp "github.com/arham09/go-di-example/service/supplier/repository"
	su "github.com/arham09/go-di-example/service/supplier/usecase"
	"github.com/google/wire"
)

func StartingService(conn *sql.DB, t time.Duration) service.Server {
	wire.Build(service.NewServer, sh.NewSupplierHandler, su.NewSupplierUsecase, sp.NewPgSupplierRepository(conn))

	return service.Server{}
}
