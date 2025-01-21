package mom

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/kareem717/k7-cbo/internal/entities/mom"
)

type MomRepository struct {
	db bun.IDB
}

func NewMomRepository(db bun.IDB) *MomRepository {
	return &MomRepository{db}
}

func (r *MomRepository) Create(ctx context.Context, params mom.CreateMomTestParams) error {
	_, err := r.db.NewInsert().Model(&params).Exec(ctx)
	return err
}

func (r *MomRepository) GetMany(ctx context.Context) ([]mom.MomTestMini, error) {
	var tests []mom.MomTestMini
	err := r.db.NewSelect().Model(&tests).Scan(ctx)
	return tests, err
}
