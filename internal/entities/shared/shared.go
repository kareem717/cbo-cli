package shared

import "time"

// TimeStamps contains the default timestamp columns used for an entity.
type TimeStamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
