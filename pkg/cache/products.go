package cache

import (
	"gorillamux/pkg/common/models"

	"golang.org/x/net/context"
)

type ProdCache interface {
	Set(ctx context.Context, key string, value *models.Product)
	Get(ctx context.Context, key string) *models.Product
}
