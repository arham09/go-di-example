package supplier

import (
	"context"

	"github.com/arham09/go-di-example/service/models"
)

type Usecase interface {
	FetchAll(ctx context.Context) ([]*models.Supplier, error)
}
