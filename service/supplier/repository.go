package supplier

import (
	"context"

	"github.com/arham09/go-di-example/service/models"
)

type Repository interface {
	FetchAll(ctx context.Context) (res []*models.Supplier, err error)
}
