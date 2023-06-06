package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ahmad20/bri-mini-project/entities"
	"github.com/ahmad20/bri-mini-project/repositories"
)

type useCase struct {
	accRepo repositories.AccountRepositoryInterface
}

type UseCaseInterface interface {
	Register(account *entities.Account) error
	GetById(id string) (*entities.Account, error)
	GetAdminsWithConditions(keyword, page, limit string) ([]*entities.Account, error)
	GetAll() ([]*entities.Account, error)
	GetWaitingApproval() ([]*entities.Account, error)
	SearchByUsername(username string) (*entities.Account, error)
	UpdateApproval(status string, account *entities.Account) error
	UpdateStatus(status string, account *entities.Account) error
	Delete(account *entities.Account) error
}

func NewUseCase(accrepo repositories.AccountRepositoryInterface) UseCaseInterface {
	return &useCase{
		accRepo: accrepo,
	}
}
func (usecase *useCase) Register(account *entities.Account) error {
	if err := usecase.accRepo.Create(account); err != nil {
		return err
	}
	return nil
}

func (usecase *useCase) GetById(id string) (*entities.Account, error) {
	account, err := usecase.accRepo.Read(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (usecase *useCase) GetAdminsWithConditions(keyword, page, limit string) ([]*entities.Account, error) {
	admins, err := usecase.accRepo.GetAdminsWithConditions(keyword, page, limit)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (usecase *useCase) GetAll() ([]*entities.Account, error) {
	admins, err := usecase.accRepo.ReadAll()

	if err != nil {
		return nil, err
	}
	return admins, nil
}
func (usecase *useCase) GetWaitingApproval() ([]*entities.Account, error) {
	accounts, err := usecase.accRepo.GetWaitingApproval()

	if err != nil {
		return nil, err
	}
	return accounts, nil
}
func (usecase *useCase) SearchByUsername(username string) (*entities.Account, error) {
	account, err := usecase.accRepo.SearchByUsername(username)

	if err != nil {
		return nil, err
	}
	return account, nil
}

func (usecase *useCase) UpdateApproval(status string, account *entities.Account) error {
	if err := usecase.accRepo.UpdateApproval(status, account); err != nil {
		return err
	}
	return nil
}

func (usecase *useCase) UpdateStatus(status string, account *entities.Account) error {
	if err := usecase.accRepo.UpdateStatus(status, account); err != nil {
		return err
	}
	return nil
}

func (usecase *useCase) Delete(account *entities.Account) error {
	if err := usecase.accRepo.Delete(account); err != nil {
		return err
	}
	return nil
}

func FetchData() []entities.Customer {
	url := "https://reqres.in/api/users?page=2"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var data struct {
		Users []entities.Customer `json:"data"`
	}
	json.Unmarshal(body, &data)

	return data.Users
}
