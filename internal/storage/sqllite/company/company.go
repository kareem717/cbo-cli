package company

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/kareem717/k7-cbo/internal/entities/company"
)

type CompanyRepository struct {
	db bun.IDB
}

func NewCompanyRepository(db bun.IDB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (r *CompanyRepository) Create(ctx context.Context, params company.CreateCompanyParams) error {
	_, err := r.db.
		NewInsert().
		Model(&params).
		Exec(ctx)

	return err
}

func (r *CompanyRepository) GetMany(ctx context.Context) ([]company.CompanyMini, error) {
	companies := []company.CompanyMini{}

	err := r.db.
		NewSelect().
		Model(&companies).
		Scan(ctx)

	return companies, err
}

func (r *CompanyRepository) GetByID(ctx context.Context, id int) (company.Company, error) {
	company := company.Company{}

	err := r.db.
		NewSelect().
		Model(&company).
		Where("id = ?", id).
		Scan(ctx)

	return company, err
}
