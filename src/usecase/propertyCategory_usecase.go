package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type PropertyCategoryUsecase struct {
	base *BaseUsecase[model.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory]
}

func NewPropertyCategoryUsecase(cfg *config.Config, repository repository.PropertyCategoryRepository) *PropertyCategoryUsecase {
	return &PropertyCategoryUsecase{
		base: NewBaseUsecase[model.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory](cfg, repository),
	}
}

// Create
func (u *PropertyCategoryUsecase) Create(ctx context.Context, req dto.CreatePropertyCategory) (dto.PropertyCategory, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *PropertyCategoryUsecase) Update(ctx context.Context, id int, req dto.UpdatePropertyCategory) (dto.PropertyCategory, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *PropertyCategoryUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *PropertyCategoryUsecase) GetById(ctx context.Context, id int) (dto.PropertyCategory, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *PropertyCategoryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.PropertyCategory], error) {
	return u.base.GetByFilter(ctx, req)
}
