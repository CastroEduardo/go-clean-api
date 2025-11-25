package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type CarModelPriceHistoryUsecase struct {
	base *BaseUsecase[model.CarModelPriceHistory, dto.CreateCarModelPriceHistory, dto.UpdateCarModelPriceHistory, dto.CarModelPriceHistory]
}

func NewCarModelPriceHistoryUsecase(cfg *config.Config, repository repository.CarModelPriceHistoryRepository) *CarModelPriceHistoryUsecase {
	return &CarModelPriceHistoryUsecase{
		base: NewBaseUsecase[model.CarModelPriceHistory, dto.CreateCarModelPriceHistory, dto.UpdateCarModelPriceHistory, dto.CarModelPriceHistory](cfg, repository),
	}
}

// Create
func (u *CarModelPriceHistoryUsecase) Create(ctx context.Context, req dto.CreateCarModelPriceHistory) (dto.CarModelPriceHistory, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CarModelPriceHistoryUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelPriceHistory) (dto.CarModelPriceHistory, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarModelPriceHistoryUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarModelPriceHistoryUsecase) GetById(ctx context.Context, id int) (dto.CarModelPriceHistory, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarModelPriceHistoryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelPriceHistory], error) {
	return u.base.GetByFilter(ctx, req)
}
