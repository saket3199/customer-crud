package service

import (
	"customerCrud/model"
	"customerCrud/repository"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	Repository repository.Repository
	Db         *gorm.DB
}

func NewService(re *repository.Repository, d *gorm.DB) *Service {
	return &Service{
		Repository: *re,
		Db:         d,
	}
}
func (s *Service) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")

}

func (s *Service) GetAll(customers *[]model.Order) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.GetAll(uow, &customers, []string{"Orders"})
}

func (s *Service) CreateNewOrder(customer *model.Order) error {

	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Add(uow, &customer)
	if err != nil {
		uow.DB.Rollback()
		return err

	}
	uow.Commit()
	return nil
}
func (s *Service) UpdateOrder(Order *model.Order) error {

	fmt.Println(Order.ID)
	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Save(uow, &Order)
	if err != nil {
		uow.DB.Rollback()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) ReturnSingleOrder(Order *model.Order, id uuid.UUID) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.Get(uow, &Order, id, []string{"Orders"})
}

func (s *Service) DeleteOrder(Order *model.Order) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.Delete(uow, &Order)
}
