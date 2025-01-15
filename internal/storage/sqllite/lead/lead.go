package lead

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/lead"
	"github.com/uptrace/bun"
)

type LeadRepository struct {
	db *bun.DB
}

func NewLeadRepository(db *bun.DB) *LeadRepository {
	return &LeadRepository{db}
}

func (r *LeadRepository) Create(ctx context.Context, params lead.CreateLeadParams) error {
	_, err := r.db.NewInsert().Model(&params).Exec(ctx)

	return err
}

func (r *LeadRepository) GetMany(ctx context.Context) ([]lead.Lead, error) {
	leads := []lead.Lead{}

	err := r.db.NewSelect().Model(&leads).Scan(ctx)

	return leads, err
}
