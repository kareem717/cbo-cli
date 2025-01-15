package lead

import (
	"github.com/kareem717/k7-cbo/internal/entities/shared"
	"github.com/uptrace/bun"
)

type Lead struct {
	bun.BaseModel `bun:"table:leads"`

	ID          int    `json:"id"`
	LinkedInURL string `json:"linked_in_url"`		

	shared.TimeStamps
}

type CreateLeadParams struct {
	bun.BaseModel `bun:"table:leads,alias:lead"`

	LinkedInURL string `bun:"linked_in_url"`
}
