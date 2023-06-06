package customer

import (
	"github.com/ahmad20/bri-mini-project/entities"
	"github.com/ahmad20/bri-mini-project/repositories"
)

type useCase struct {
	custRepo repositories.CustomerRepositoryInterface
}
type UseCaseInterface interface {
	Register(customer *entities.Customer) error
	GetById(id string) (*entities.Customer, error)
	Delete(customer *entities.Customer) error
	GetCustomersWithConditions(keyword, page, limit string) ([]*entities.Customer, error)
	GetAll() ([]*entities.Customer, error)
	CountList(customer *entities.Customer) (int64, error)
}

func NewUseCase(custRepo repositories.CustomerRepositoryInterface) UseCaseInterface {
	return &useCase{
		custRepo: custRepo,
	}
}

func (usecase *useCase) Register(customer *entities.Customer) error {
	if err := usecase.custRepo.Create(customer); err != nil {
		return err
	}
	return nil
}

func (usecase *useCase) GetById(id string) (*entities.Customer, error) {
	customer, err := usecase.custRepo.Read(id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (usecase *useCase) Delete(customer *entities.Customer) error {
	if err := usecase.custRepo.Delete(customer); err != nil {
		return err
	}
	return nil
}
func (usecase *useCase) GetCustomersWithConditions(keyword, page, limit string) ([]*entities.Customer, error) {
	customers, err := usecase.custRepo.GetCustomersWithConditions(keyword, page, limit)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
func (usecase *useCase) GetAll() ([]*entities.Customer, error) {
	customers, err := usecase.custRepo.ReadAll()

	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (usecase *useCase) CountList(customer *entities.Customer) (int64, error) {
	count, _ := usecase.custRepo.CountList(customer)
	return count, nil
}
