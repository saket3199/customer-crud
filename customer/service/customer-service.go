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

func (s *Service) GetAll(customers *[]model.Customer) error {

	uow := repository.NewUnitOfWork(s.Db, true)
	err := s.Repository.GetAll(uow, &customers, []string{"Orders"})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateCustomer(customer *model.Customer) error {

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
	err := s.Repository.Update(uow, &customer)
	if err != nil {
		uow.DB.Rollback()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) GetCustomer(customer *model.Customer, id uuid.UUID) error {

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
	if err != nil {
		uow.DB.Rollback()
		return err
	}
	err = s.Repository.Delete(uow, &model.Order{}, map[string]interface{}{"customer_id = ?": customer.ID})
	if err != nil {
		uow.DB.Rollback()
		return err
	}
	uow.Commit()
	return nil
}
