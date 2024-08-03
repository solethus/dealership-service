package respository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DealershipRepository interface {
	GetDealership(ctx context.Context, id string) (*model.Dealership, error)
}

type dealershipRepositoryImpl struct {
	db *sqlx.DB
}

func NewDealershipRepository(db *sqlx.DB) DealershipRepository {
	return &dealershipRepositoryImpl{
		db: db,
	}
}

func (d dealershipRepositoryImpl) GetDealership(ctx context.Context, id string) (*model.Dealership, error) {
	//TODO implement me
	panic("implement me")
}
