package company

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/company"
	"github.com/kareem717/k7-cbo/internal/storage"
)

type CompanyService struct {
	repo storage.CompanyRepository
}

// NewTestService returns a new instance of test service.
func NewCompanyService(repo storage.CompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (s *CompanyService) Create(ctx context.Context, params company.CreateCompanyParams) error {
	return s.repo.Create(ctx, params)
}

func (s *CompanyService) GetMany(ctx context.Context) ([]company.CompanyMini, error) {
	return s.repo.GetMany(ctx)
}
