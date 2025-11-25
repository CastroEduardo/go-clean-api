package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type ColorUsecase struct {
	base *BaseUsecase[model.Color, dto.CreateColor, dto.UpdateColor, dto.Color]
}

func NewColorUsecase(cfg *config.Config, repository repository.ColorRepository) *ColorUsecase {
	return &ColorUsecase{
		base: NewBaseUsecase[model.Color, dto.CreateColor, dto.UpdateColor, dto.Color](cfg, repository),
	}
}

// Create
func (u *ColorUsecase) Create(ctx context.Context, req dto.CreateColor) (dto.Color, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *ColorUsecase) Update(ctx context.Context, id int, req dto.UpdateColor) (dto.Color, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *ColorUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *ColorUsecase) GetById(ctx context.Context, id int) (dto.Color, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *ColorUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Color], error) {
	return u.base.GetByFilter(ctx, req)
}
