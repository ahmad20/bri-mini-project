package repositories

import (
	"errors"
	"strconv"

	"github.com/ahmad20/bri-mini-project/entities"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}
type AccountRepositoryInterface interface {
	Create(account *entities.Account) error
	Read(id string) (*entities.Account, error)
	ReadAll() ([]*entities.Account, error)
	Update(id string, account *entities.Account) error
	Delete(account *entities.Account) error
	SearchByUsername(username string) (*entities.Account, error)
	UpdateApproval(status string, account *entities.Account) error
	UpdateStatus(status string, account *entities.Account) error
	GetWaitingApproval() ([]*entities.Account, error)
	GetAdminsWithConditions(keyword, page, limit string) ([]*entities.Account, error)
}

func NewAccountRepository(db *gorm.DB) AccountRepositoryInterface {
	return &AccountRepository{
		db: db,
	}
}
func (repo AccountRepository) Create(account *entities.Account) error {
	err := repo.db.Create(account).Error

	if err != nil {
		return err
	}
	return nil
}

func (repo AccountRepository) Read(id string) (*entities.Account, error) {
	var account entities.Account
	result := repo.db.Where("id = ?", id).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}
func (repo AccountRepository) ReadAll() ([]*entities.Account, error) {
	var account []*entities.Account
	result := repo.db.Find(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}
func (repo AccountRepository) Update(id string, account *entities.Account) error {
	if err := repo.db.Save(&account).Error; err != nil {
		return err
	}
	return nil
}
func (repo AccountRepository) Delete(account *entities.Account) error {
	if err := repo.db.Delete(&account).Error; err != nil {
		return err
	}
	return nil
}
func (repo AccountRepository) SearchByUsername(username string) (*entities.Account, error) {
	var account entities.Account
	result := repo.db.Where("username = ?", username).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (repo AccountRepository) UpdateApproval(status string, account *entities.Account) error {
	if err := repo.db.Model(&account).Update("approval_status", status).Error; err != nil {
		return err
	}
	return nil
}
func (repo AccountRepository) GetWaitingApproval() ([]*entities.Account, error) {
	var account []*entities.Account
	result := repo.db.Where("approval_status = ?", "waiting").Find(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}
func (repo AccountRepository) UpdateStatus(status string, account *entities.Account) error {
	if err := repo.db.Model(&account).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
func (repo AccountRepository) GetAdminsWithConditions(keyword, page, limit string) ([]*entities.Account, error) {
	var account []*entities.Account
	query := repo.db.Model(&account)
	if keyword != "" {
		query = query.Where("username like ?", "%"+keyword+"%")
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
	// Retrieve the admins based on the applied conditions
	if err := query.Find(&account).Error; err != nil {
		return nil, err
	}
	// Return the retrieved admins
	return account, nil
}
