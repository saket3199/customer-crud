package service

import (
	"fmt"
	"net/http"

	"customerCrud/model"
	"customerCrud/repository"

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

func (s *Service) GetAll(customers *[]model.Customer) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.GetAll(uow, &customers, []string{"Orders"})
}

func (s *Service) CreateNewCustomer(customer *model.Customer) error {

	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Add(uow, &customer)
	if err != nil {
		uow.DB.Rollback()
		return err

	}
	uow.Commit()
	return nil
}
func (s *Service) UpdateCustomer(customer *model.Customer) error {

	fmt.Println(customer.ID)
	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Save(uow, &customer)
	if err != nil {
		uow.DB.Rollback()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) ReturnSingleCustomer(customer *model.Customer, id uuid.UUID) error {

	uow := repository.NewUnitOfWork(s.Db, true)
	err := s.Repository.Get(uow, &customer, id, []string{"Orders"})
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) DeleteCustomer(customer *model.Customer) error {

	uow := repository.NewUnitOfWork(s.Db, true)
	err := s.Repository.Delete(uow, &customer)
	orders := customer.Orders
	for i := 0; i < len(orders); i++ {
		err := s.Repository.Delete(uow, &i)
		if err != nil {
			return err
		}
		uow.Commit()
	}
	// db.Where("customerId = ?", customer.ID).Delete(&model.Order{})
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}
