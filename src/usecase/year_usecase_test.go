package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type PersianYearTestUsecase struct {
	base *BaseUsecase[
		model.PersianYearTest,
		dto.CreatePersianYearTest,
		dto.UpdatePersianYearTest,
		dto.PersianYearTest,
	]
}

func NewPersianYearTestUsecase(cfg *config.Config, repo repository.PersianYearTestRepository) *PersianYearTestUsecase {
	return &PersianYearTestUsecase{
		base: NewBaseUsecase[
			model.PersianYearTest,
			dto.CreatePersianYearTest,
			dto.UpdatePersianYearTest,
			dto.PersianYearTest,
		](cfg, repo),
	}
}

// Create
func (u *PersianYearTestUsecase) Create(ctx context.Context, req dto.CreatePersianYearTest) (dto.PersianYearTest, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *PersianYearTestUsecase) Update(ctx context.Context, id int, req dto.UpdatePersianYearTest) (dto.PersianYearTest, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *PersianYearTestUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *PersianYearTestUsecase) GetById(ctx context.Context, id int) (dto.PersianYearTest, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *PersianYearTestUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.PersianYearTest], error) {
	return u.base.GetByFilter(ctx, req)
}
