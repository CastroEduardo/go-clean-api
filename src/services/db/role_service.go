package service

import (
	"reflect"

	"github.com/CastroEduardo/go-clean-api/domain/model"
	"github.com/CastroEduardo/go-clean-api/infra/persistence/database"
	"github.com/CastroEduardo/go-clean-api/pkg/metrics"
	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func NewDbRoleService() *RoleService {
	return &RoleService{
		db: database.GetDb(),
	}
}

func (s *RoleService) recordMetric(operation string, success bool) {
	status := "Success"
	if !success {
		status = "Failed"
	}
	metrics.DbCall.
		WithLabelValues(reflect.TypeOf(model.Role{}).String(), operation, status).
		Inc()
}

func (s *RoleService) Add(role *model.Role) error {
	err := s.db.Create(role).Error
	s.recordMetric("Add", err == nil)
	return err
}

func (s *RoleService) Update(role *model.Role) error {
	err := s.db.Save(role).Error
	s.recordMetric("Update", err == nil)
	return err
}

func (s *RoleService) Remove(id uint) error {
	err := s.db.Delete(&model.Role{}, id).Error
	s.recordMetric("Remove", err == nil)
	return err
}

func (s *RoleService) FindByID(id uint) (*model.Role, error) {
	var role model.Role
	err := s.db.First(&role, id).Error
	s.recordMetric("FindByID", err == nil)
	return &role, err
}

func (s *RoleService) FindAll() ([]model.Role, error) {
	var list []model.Role
	err := s.db.Find(&list).Error
	s.recordMetric("FindAll", err == nil)
	return list, err
}
