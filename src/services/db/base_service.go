package service

// import (
// 	"reflect"

// 	"github.com/CastroEduardo/go-clean-api/pkg/metrics"
// 	"gorm.io/gorm"
// )

// // BaseService se puede usar para cualquier modelo
// type BaseService[T any] struct {
// 	db *gorm.DB
// }

// // NewBaseService retorna un nuevo servicio genérico
// func NewBaseService[T any](db *gorm.DB) *BaseService[T] {
// 	return &BaseService[T]{db: db}
// }

// // helper interno: registra métricas
// func (s *BaseService[T]) recordMetric(operation string, success bool) {
// 	var model T
// 	status := "Success"
// 	if !success {
// 		status = "Failed"
// 	}

// 	metrics.DbCall.
// 		WithLabelValues(reflect.TypeOf(model).String(), operation, status).
// 		Inc()
// }

// // Add crea un registro
// func (s *BaseService[T]) Add(entity *T) error {
// 	err := s.db.Create(entity).Error
// 	s.recordMetric("Add", err == nil)
// 	return err
// }

// // Update actualiza un registro existente
// func (s *BaseService[T]) Update(entity *T) error {
// 	err := s.db.Save(entity).Error
// 	s.recordMetric("Update", err == nil)
// 	return err
// }

// // Remove elimina por ID
// func (s *BaseService[T]) Remove(id uint) error {
// 	err := s.db.Delete(new(T), id).Error
// 	s.recordMetric("Remove", err == nil)
// 	return err
// }

// // FindByID obtiene por el ID
// func (s *BaseService[T]) FindByID(id uint) (*T, error) {
// 	entity := new(T)
// 	err := s.db.First(entity, id).Error
// 	s.recordMetric("FindByID", err == nil)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return entity, nil
// }

// // FindAll devuelve todos los registros
// func (s *BaseService[T]) FindAll() ([]T, error) {
// 	var list []T
// 	err := s.db.Find(&list).Error
// 	s.recordMetric("FindAll", err == nil)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return list, nil
// }
