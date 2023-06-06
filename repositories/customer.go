package repositories

import (
	"errors"
	"strconv"

	"github.com/ahmad20/bri-mini-project/entities"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}
type CustomerRepositoryInterface interface {
	Create(customer *entities.Customer) error
	Read(id string) (*entities.Customer, error)
	ReadAll() ([]*entities.Customer, error)
	Update(id string, customer *entities.Customer) error
	Delete(customer *entities.Customer) error
	CountList(customer *entities.Customer) (int64, error)
	GetCustomersWithConditions(keyword, page, limit string) ([]*entities.Customer, error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepositoryInterface {
	return &CustomerRepository{
		db: db,
	}
}
func (repo CustomerRepository) Create(customer *entities.Customer) error {
	result := repo.db.Create(customer)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo CustomerRepository) Read(id string) (*entities.Customer, error) {
	var customer entities.Customer
	result := repo.db.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}
func (repo CustomerRepository) ReadAll() ([]*entities.Customer, error) {
	var customer []*entities.Customer
	result := repo.db.Find(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}
func (repo CustomerRepository) Update(id string, customer *entities.Customer) error {
	if err := repo.db.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}
func (repo CustomerRepository) Delete(customer *entities.Customer) error {
	if err := repo.db.Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}
func (repo CustomerRepository) CountList(customer *entities.Customer) (int64, error) {
	var count int64
	if err := repo.db.Model(&customer).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}
func (repo CustomerRepository) GetCustomersWithConditions(keyword, page, limit string) ([]*entities.Customer, error) {
	var customer []*entities.Customer
	query := repo.db.Model(&customer)
	if keyword != "" {
		query = query.Where("first_name like ? or last_name like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if page != "" && limit != "" {
		offset, err := strconv.Atoi(page)
		if err != nil {
			return nil, errors.New("invalid page parameter")
		}
		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("invalid limit parameter")
		}
		query = query.Offset((offset - 1) * limitNum).Limit(limitNum)
	}

	if err := query.Find(&customer).Error; err != nil {
		return nil, err
	}
	// Return the retrieved customers
	return customer, nil
}
