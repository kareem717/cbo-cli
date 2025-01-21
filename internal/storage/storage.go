package storage

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/company"
)

type CompanyRepository interface {
	Create(ctx context.Context, params company.CreateCompanyParams) error
	GetMany(ctx context.Context) ([]company.CompanyMini, error)
	GetByID(ctx context.Context, id int) (company.Company, error)
}

type Repository interface {
	Company() CompanyRepository

	// Migrate applies all pending migrations
	Migrate(ctx context.Context) error
	HealthCheck(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
