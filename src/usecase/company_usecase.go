package usecase

import (
	"context"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/domain/filter"
	model "github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/domain/repository"
	"github.com/CastroEduardo/go-clean-api/usecase/dto"
)

type CompanyUsecase struct {
	base *BaseUsecase[model.Company, dto.CreateCompany, dto.UpdateCompany, dto.Company]
}

func NewCompanyUsecase(cfg *config.Config, repository repository.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{
		base: NewBaseUsecase[model.Company, dto.CreateCompany, dto.UpdateCompany, dto.Company](cfg, repository),
	}
}

// Create
func (u *CompanyUsecase) Create(ctx context.Context, req dto.CreateCompany) (dto.Company, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CompanyUsecase) Update(ctx context.Context, id int, req dto.UpdateCompany) (dto.Company, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CompanyUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CompanyUsecase) GetById(ctx context.Context, id int) (dto.Company, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CompanyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Company], error) {
	return u.base.GetByFilter(ctx, req)
}
