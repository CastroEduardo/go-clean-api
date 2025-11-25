package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type CarTypeUsecase struct {
	base *BaseUsecase[model.CarType, dto.Name, dto.Name, dto.IdName]
}

func NewCarTypeUsecase(cfg *config.Config, repository repository.CarTypeRepository) *CarTypeUsecase {
	return &CarTypeUsecase{
		base: NewBaseUsecase[model.CarType, dto.Name, dto.Name, dto.IdName](cfg, repository),
	}
}

// Create
func (u *CarTypeUsecase) Create(ctx context.Context, req dto.Name) (dto.IdName, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CarTypeUsecase) Update(ctx context.Context, id int, req dto.Name) (dto.IdName, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarTypeUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarTypeUsecase) GetById(ctx context.Context, id int) (dto.IdName, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarTypeUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.IdName], error) {
	return u.base.GetByFilter(ctx, req)
}
