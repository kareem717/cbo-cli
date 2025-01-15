package storage

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/lead"
)

type LeadRepository interface {
	Create(ctx context.Context, params lead.CreateLeadParams) error
	GetMany(ctx context.Context) ([]lead.Lead, error)
}

type Repository interface {
	Lead() LeadRepository

	// Migrate applies all pending migrations
	Migrate(ctx context.Context) error
	HealthCheck(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
