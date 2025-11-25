package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type GearboxUsecase struct {
	base *BaseUsecase[model.Gearbox, dto.Name, dto.Name, dto.IdName]
}

func NewGearboxUsecase(cfg *config.Config, repository repository.GearboxRepository) *GearboxUsecase {
	return &GearboxUsecase{
		base: NewBaseUsecase[model.Gearbox, dto.Name, dto.Name, dto.IdName](cfg, repository),
	}
}

// Create
func (u *GearboxUsecase) Create(ctx context.Context, req dto.Name) (dto.IdName, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *GearboxUsecase) Update(ctx context.Context, id int, req dto.Name) (dto.IdName, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *GearboxUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *GearboxUsecase) GetById(ctx context.Context, id int) (dto.IdName, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *GearboxUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.IdName], error) {
	return u.base.GetByFilter(ctx, req)
}
