package service

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/company"
	"github.com/kareem717/k7-cbo/internal/entities/mom"
	companyservice "github.com/kareem717/k7-cbo/internal/service/domain/company"
	momservice "github.com/kareem717/k7-cbo/internal/service/domain/mom"
	"github.com/kareem717/k7-cbo/internal/storage"
)

type CompanyService interface {
	Create(ctx context.Context, params company.CreateCompanyParams) error
	GetMany(ctx context.Context) ([]company.CompanyMini, error)
}

type MomService interface {
	Create(ctx context.Context, params mom.CreateMomTestParams) error
	GetMany(ctx context.Context) ([]mom.MomTestMini, error)
}

type Service struct {
	repositories storage.Repository
	Company      CompanyService
	Mom          MomService
}

// NewService implementation for storage of all services.
func NewService(
	repositories storage.Repository,
) *Service {
	return &Service{
		repositories: repositories,
		Company:      companyservice.NewCompanyService(repositories.Company()),
		Mom:          momservice.NewMomService(repositories.Mom()),
	}
}
