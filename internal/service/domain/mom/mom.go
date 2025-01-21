package mom

import (
	"context"

	"github.com/kareem717/k7-cbo/internal/entities/mom"
	"github.com/kareem717/k7-cbo/internal/storage"
)

type MomService struct {
	repo storage.MomRepository
}

// NewTestService returns a new instance of test service.
func NewMomService(repo storage.MomRepository) *MomService {
	return &MomService{
		repo: repo,
	}
}

func (s *MomService) Create(ctx context.Context, params mom.CreateMomTestParams) error {
	return s.repo.Create(ctx, params)
}

func (s *MomService) GetMany(ctx context.Context) ([]mom.MomTestMini, error) {
	return s.repo.GetMany(ctx)
}
