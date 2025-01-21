package company

import (
	"github.com/kareem717/k7-cbo/internal/entities/shared"
	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel `bun:"table:companies"`

	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	shared.TimeStamps
}

// CompanyMini is a minified version of a Company entity, used for mass indexing/listing.
type CompanyMini struct {
	bun.BaseModel `bun:"table:companies"`

	ID   int    `json:"id"`
	Name string `json:"name"`

	// Description is a truncated version of the company description
	Description string `json:"description"`

	shared.TimeStamps
}

type CreateCompanyParams struct {
	bun.BaseModel `bun:"table:companies,alias:company"`

	Name        string `json:"name"`
	Description string `json:"description"`
}
