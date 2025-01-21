package storage

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/company"
	"github.com/kareem717/k7-cbo/internal/entities/mom"
)

type CompanyRepository interface {
	Create(ctx context.Context, params company.CreateCompanyParams) error
	GetMany(ctx context.Context) ([]company.CompanyMini, error)
	GetByID(ctx context.Context, id int) (company.Company, error)
}

type MomRepository interface {
	Create(ctx context.Context, params mom.CreateMomTestParams) error
	GetMany(ctx context.Context) ([]mom.MomTestMini, error)
}

type Repository interface {
	Company() CompanyRepository
	Mom() MomRepository
	// Migrate applies all pending migrations
	Migrate(ctx context.Context) error
	HealthCheck(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
