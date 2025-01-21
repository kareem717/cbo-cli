package service

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/company"
	companyservice "github.com/kareem717/k7-cbo/internal/service/domain/company"
	"github.com/kareem717/k7-cbo/internal/storage"
)

type CompanyService interface {
	Create(ctx context.Context, params company.CreateCompanyParams) error
	GetMany(ctx context.Context) ([]company.CompanyMini, error)
}

type Service struct {
	repositories storage.Repository
	Company      CompanyService
}

// NewService implementation for storage of all services.
func NewService(
	repositories storage.Repository,
) *Service {
	return &Service{
		repositories: repositories,
		Company:      companyservice.NewCompanyService(repositories.Company()),
	}
}
