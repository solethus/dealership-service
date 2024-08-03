package respository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"

	"github.com/solethus/dealership-service/internal/model"
)

type DealershipRepository interface {
	CreateDealership(ctx context.Context, dealership *model.Dealership) error
	GetDealership(ctx context.Context, id uuid.UUID) (*model.Dealership, error)
	UpdateDealership(ctx context.Context, dealership *model.Dealership) error
	DeleteDealership(ctx context.Context, id uuid.UUID) error
	ListDealerships(ctx context.Context, filter model.DealershipFilter) ([]*model.Dealership, int, error)
	GetDealershipStats(ctx context.Context) (*model.DealershipStats, error)
	AddEmployee(ctx context.Context, dealershipID uuid.UUID, employee *model.Employee) error
	UpdateInventoryCount(ctx context.Context, dealershipID uuid.UUID, change int) error
}

type dealershipRepositoryImpl struct {
	db *sql.DB
}

func NewDealershipRepository(db *sql.DB) DealershipRepository {
	return &dealershipRepositoryImpl{
		db: db,
	}
}

func (r *dealershipRepositoryImpl) CreateDealership(ctx context.Context, dealership *model.Dealership) error {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) GetDealership(ctx context.Context, id uuid.UUID) (*model.Dealership, error) {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) UpdateDealership(ctx context.Context, dealership *model.Dealership) error {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) DeleteDealership(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) ListDealerships(ctx context.Context, filter model.DealershipFilter) ([]*model.Dealership, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) GetDealershipStats(ctx context.Context) (*model.DealershipStats, error) {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) AddEmployee(ctx context.Context, dealershipID uuid.UUID, employee *model.Employee) error {
	//TODO implement me
	panic("implement me")
}

func (r *dealershipRepositoryImpl) UpdateInventoryCount(ctx context.Context, dealershipID uuid.UUID, change int) error {
	//TODO implement me
	panic("implement me")
}
