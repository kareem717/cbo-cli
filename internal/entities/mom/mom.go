package mom

import (
	"github.com/kareem717/k7-cbo/internal/entities/shared"
	"github.com/uptrace/bun"
)

type MomTest struct {
	bun.BaseModel `bun:"table:mom_tests,alias:mom_test"`

	ID         int    `json:"id"`
	CompanyID  int    `json:"company_id"`
	Hypothesis string `json:"hypothesis"`
	shared.TimeStamps
}

// MomTestMini is a simplified version of a MomTest entity for listing.
type MomTestMini struct {
	bun.BaseModel `bun:"table:mom_tests,alias:mom_test"`

	ID         int    `json:"id"`
	CompanyID  int    `json:"company_id"`
	Hypothesis string `json:"hypothesis"`
}

type CreateMomTestParams struct {
	bun.BaseModel `bun:"table:mom_tests,alias:mom_test"`

	CompanyID  int    `json:"company_id"`
	Hypothesis string `json:"hypothesis"`
}
