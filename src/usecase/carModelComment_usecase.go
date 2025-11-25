package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type CarModelCommentUsecase struct {
	base *BaseUsecase[model.CarModelComment, dto.CreateCarModelComment, dto.UpdateCarModelComment, dto.CarModelComment]
}

func NewCarModelCommentUsecase(cfg *config.Config, repository repository.CarModelCommentRepository) *CarModelCommentUsecase {
	return &CarModelCommentUsecase{
		base: NewBaseUsecase[model.CarModelComment, dto.CreateCarModelComment, dto.UpdateCarModelComment, dto.CarModelComment](cfg, repository),
	}
}

// Create
func (u *CarModelCommentUsecase) Create(ctx context.Context, req dto.CreateCarModelComment) (dto.CarModelComment, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CarModelCommentUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelComment) (dto.CarModelComment, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarModelCommentUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarModelCommentUsecase) GetById(ctx context.Context, id int) (dto.CarModelComment, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarModelCommentUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelComment], error) {
	return u.base.GetByFilter(ctx, req)
}
