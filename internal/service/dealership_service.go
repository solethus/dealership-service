package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/solethus/dealership-service/internal/model"
	"github.com/solethus/dealership-service/internal/respository"
)

type DealershipService interface {
	CreateDealership(ctx context.Context, dealership *model.Dealership) error
	GetDealership(ctx context.Context, id uuid.UUID) (*model.Dealership, error)
	UpdateDealership(ctx context.Context, dealership *model.Dealership) error
	SearchDealerships(ctx context.Context, filter model.DealershipFilter) ([]*model.Dealership, int, error)
	GenerateDealershipReport(ctx context.Context) (*model.DealershipStats, error)
	AddEmployeeToDealership(ctx context.Context, dealershipID uuid.UUID, employee *model.Employee) error
	UpdateDealershipInventory(ctx context.Context, dealershipID uuid.UUID, change int) error
}

type dealershipService struct {
	repo respository.DealershipRepository
}

func NewDealershipService(repo respository.DealershipRepository) DealershipService {
	return &dealershipService{
		repo: repo,
	}
}

func (s *dealershipService) CreateDealership(ctx context.Context, dealership *model.Dealership) error {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) GetDealership(ctx context.Context, id uuid.UUID) (*model.Dealership, error) {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) UpdateDealership(ctx context.Context, dealership *model.Dealership) error {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) SearchDealerships(ctx context.Context, filter model.DealershipFilter) ([]*model.Dealership, int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) GenerateDealershipReport(ctx context.Context) (*model.DealershipStats, error) {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) AddEmployeeToDealership(ctx context.Context, dealershipID uuid.UUID, employee *model.Employee) error {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) UpdateDealershipInventory(ctx context.Context, dealershipID uuid.UUID, change int) error {
	//TODO implement me
	panic("implement me")
}

func (s *dealershipService) GetDealershipInfo(ctx context.Context, id string) error {
	//TODO implement me

	panic("implement me")
}

func (s *dealershipService) ListDealerships(ctx context.Context, pageSize, pageNumber int32) ([]*model.Dealership, int32, error) {
	//TODO Implementation
	panic("implement me")
}

func (s *dealershipService) CreateReservation(ctx context.Context, customerID, carID, dealershipID string) (*model.Reservation, error) {
	//TODO Implementation
	panic("implement me")
}
