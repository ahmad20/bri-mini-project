package customer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ahmad20/bri-mini-project/entities"
	mocks "github.com/ahmad20/bri-mini-project/mocks/repositories"
)

func Test_useCase_Register(t *testing.T) {
	type args struct {
		customer *entities.Customer
	}
	mockRepo := &mocks.CustomerRepositoryInterface{}
	tests := []struct {
		name    string
		usecase *useCase
		args    args
		want    error
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Successful registration",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				customer: &entities.Customer{
					ID:         1,
					Email:      "testuser@email.com",
					First_Name: "test",
					Last_Name:  "user",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Create", tt.args.customer).Return(tt.want)
			err := tt.usecase.Register(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("useCase.Register() = %v, want %v", err, tt.want)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_GetById(t *testing.T) {
	type args struct {
		id string
	}
	type wants struct {
		cust *entities.Customer
		err  error
	}
	mockRepo := &mocks.CustomerRepositoryInterface{}
	tests := []struct {
		name    string
		usecase *useCase
		args    args
		want    wants
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Valid ID",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				id: "123",
			},
			want: wants{
				cust: &entities.Customer{
					ID:         123,
					Email:      "testuser@mail.com",
					First_Name: "test",
					Last_Name:  "user",
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid ID",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				id: "456",
			},
			want: wants{
				cust: nil,
				err:  errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Read", tt.args.id).Return(tt.want.cust, tt.want.err)
			got, err := tt.usecase.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want.cust) {
				t.Errorf("useCase.GetById() = %v, want %v", got, tt.want.cust)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_Delete(t *testing.T) {
	type args struct {
		customer *entities.Customer
	}
	mockRepo := &mocks.CustomerRepositoryInterface{}
	tests := []struct {
		name    string
		usecase *useCase
		args    args
		want    error
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				customer: &entities.Customer{
					ID:         1,
					Email:      "testuser@email.com",
					First_Name: "test",
					Last_Name:  "user",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "failed: user not found",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				customer: &entities.Customer{},
			},
			want:    errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Delete", tt.args.customer).Return(tt.want)
			err := tt.usecase.Delete(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("useCase.Delete() = %v, want %v", err, tt.want)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_GetCustomersWithConditions(t *testing.T) {
	type args struct {
		keyword string
		page    string
		limit   string
	}
	type wants struct {
		custs []*entities.Customer
		err   error
	}
	mockRepo := &mocks.CustomerRepositoryInterface{}
	tests := []struct {
		name    string
		usecase *useCase
		args    args
		want    wants
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Valid Conditions",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				keyword: "jo",
				page:    "1",
				limit:   "10",
			},
			want: wants{
				custs: []*entities.Customer{
					{
						ID:         1,
						Email:      "johndoe@email.com",
						First_Name: "john",
						Last_Name:  "doe",
					},
					{
						ID:         1,
						Email:      "davidjoel@email.com",
						First_Name: "david",
						Last_Name:  "joel",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid Conditions",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			args: args{
				keyword: "jo",
				page:    "invalid",
				limit:   "10",
			},
			want: wants{
				custs: nil,
				err:   errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetCustomersWithConditions", tt.args.keyword, tt.args.page, tt.args.limit).Return(tt.want.custs, tt.want.err)
			got, err := tt.usecase.GetCustomersWithConditions(tt.args.keyword, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetCustomersWithConditions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want.custs) {
				t.Errorf("useCase.GetCustomersWithConditions() = %v, want %v", got, tt.want.custs)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	mockRepo := &mocks.CustomerRepositoryInterface{}
	type wants struct {
		custs []*entities.Customer
		err   error
	}
	tests := []struct {
		name    string
		usecase *useCase
		want    wants
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			usecase: &useCase{
				custRepo: mockRepo,
			},
			want: wants{
				custs: []*entities.Customer{
					{
						ID:         1,
						Email:      "johndoe@email.com",
						First_Name: "john",
						Last_Name:  "doe",
					},
					{
						ID:         1,
						Email:      "davidjoel@email.com",
						First_Name: "david",
						Last_Name:  "joel",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("ReadAll").Return(tt.want.custs, tt.want.err)
			got, err := tt.usecase.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetAll() error = %v, wantErr %v", (err != nil), tt.wantErr)
			}
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("useCase.GetAll() error = %v, want %v", err, tt.want.err)
			}
			if !reflect.DeepEqual(got, tt.want.custs) {
				t.Errorf("useCase.GetAll() = %v, want %v", got, tt.want.custs)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
