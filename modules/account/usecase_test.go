package account

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ahmad20/bri-mini-project/entities"
	mocks "github.com/ahmad20/bri-mini-project/mocks/repositories"
)

func Test_useCase_Register(t *testing.T) {
	type args struct {
		account *entities.Account
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
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
				accRepo: mockRepo,
			},
			args: args{
				account: &entities.Account{
					Username: "testuser",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Failed registration: bad request",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				account: &entities.Account{
					Username: "testuser",
					Password: "password",
					Role:     "testrole",
				},
			},
			want:    errors.New("bad request"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Create", tt.args.account).Return(tt.want)
			err := tt.usecase.Register(tt.args.account)
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
		acc *entities.Account
		err error
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
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
				accRepo: mockRepo,
			},
			args: args{
				id: "123",
			},
			want: wants{
				acc: &entities.Account{
					ID:       "123",
					Username: "testuser",
					Password: "password",
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid ID",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				id: "456",
			},
			want: wants{
				acc: nil,
				err: errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Read", tt.args.id).Return(tt.want.acc, tt.want.err)
			got, err := tt.usecase.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want.acc) {
				t.Errorf("useCase.GetById() = %v, want %v", got, tt.want.acc)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_GetAdminsWithConditions(t *testing.T) {
	type args struct {
		keyword string
		page    string
		limit   string
	}
	type wants struct {
		accs []*entities.Account
		err  error
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
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
				accRepo: mockRepo,
			},
			args: args{
				keyword: "admin",
				page:    "1",
				limit:   "10",
			},
			want: wants{
				accs: []*entities.Account{
					{
						ID:       "1",
						Username: "admin1",
						Password: "password1",
					},
					{
						ID:       "2",
						Username: "admin2",
						Password: "password2",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid Conditions",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				keyword: "admin",
				page:    "invalid",
				limit:   "10",
			},
			want: wants{
				accs: nil,
				err:  errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetAdminsWithConditions", tt.args.keyword, tt.args.page, tt.args.limit).Return(tt.want.accs, tt.want.err)
			got, err := tt.usecase.GetAdminsWithConditions(tt.args.keyword, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetAdminsWithConditions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want.accs) {
				t.Errorf("useCase.GetAdminsWithConditions() = %v, want %v", got, tt.want.accs)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	mockRepo := &mocks.AccountRepositoryInterface{}
	type wants struct {
		accs []*entities.Account
		err  error
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
				accRepo: mockRepo,
			},
			want: wants{
				accs: []*entities.Account{
					{
						ID:       "1",
						Username: "admin1",
						Password: "password1",
					},
					{
						ID:       "2",
						Username: "admin2",
						Password: "password2",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("ReadAll").Return(tt.want.accs, tt.want.err)
			got, err := tt.usecase.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetAll() error = %v, wantErr %v", (err != nil), tt.wantErr)
			}
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("useCase.GetAll() error = %v, want %v", err, tt.want.err)
			}
			if !reflect.DeepEqual(got, tt.want.accs) {
				t.Errorf("useCase.GetAll() = %v, want %v", got, tt.want.accs)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_GetWaitingApproval(t *testing.T) {
	mockRepo := &mocks.AccountRepositoryInterface{}
	type wants struct {
		accs []*entities.Account
		err  error
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
				accRepo: mockRepo,
			},
			want: wants{
				accs: []*entities.Account{
					{
						ID:       "1",
						Username: "admin1",
						Password: "password1",
					},
					{
						ID:       "2",
						Username: "admin2",
						Password: "password2",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetWaitingApproval").Return(tt.want.accs, tt.want.err)
			got, err := tt.usecase.GetWaitingApproval()
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetWaitingApproval() error = %v, wantErr %v", (err != nil), tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want.accs) {
				t.Errorf("useCase.GetWaitingApproval() = %v, want %v", got, tt.want.accs)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_SearchByUsername(t *testing.T) {
	type args struct {
		username string
	}
	type wants struct {
		accs *entities.Account
		err  error
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
	tests := []struct {
		name    string
		usecase *useCase
		args    args
		want    wants
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Valid username",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				username: "admin1",
			},
			want: wants{
				accs: &entities.Account{
					ID:       "123",
					Username: "testuser",
					Password: "password",
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Invalid username",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				username: "nonexistent",
			},
			want: wants{
				accs: nil,
				err:  errors.New("not found"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("SearchByUsername", tt.args.username).Return(tt.want.accs, tt.want.err)
			got, err := tt.usecase.SearchByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetAdminsWithConditions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want.accs) {
				t.Errorf("useCase.GetAdminsWithConditions() = %v, want %v", got, tt.want.accs)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_UpdateApproval(t *testing.T) {
	type args struct {
		status  string
		account *entities.Account
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
	tests := []struct {
		name       string
		usecase    *useCase
		args       args
		want       error
		wantErr    bool
		mockStatus string
	}{
		// TODO: Add test cases.
		{
			name: "Update status from waiting to approved",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				status: "approved",
				account: &entities.Account{
					ID:             "1",
					Username:       "admin1",
					Password:       "password1",
					ApprovalStatus: "waiting", // Set the initial status to "waiting"
				},
			},
			want:       nil,
			wantErr:    false,
			mockStatus: "waiting", // Set the initial status of the mock repository to "waiting"
		},
		{
			name: "Update status from waiting to approved (failure)",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				status: "approved",
				account: &entities.Account{
					ID:       "2",
					Username: "admin2",
					Password: "password2",
					Status:   "waiting",
				},
			},
			want:       errors.New("update failed"),
			wantErr:    true,
			mockStatus: "waiting",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("UpdateApproval", tt.args.status, tt.args.account).Return(tt.want)
			err := tt.usecase.UpdateApproval(tt.args.status, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.UpdateApproval() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("useCase.UpdateApproval() = %v, want %v", err, tt.want)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_useCase_UpdateStatus(t *testing.T) {
	type args struct {
		status  string
		account *entities.Account
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
	tests := []struct {
		name       string
		usecase    *useCase
		args       args
		want       error
		wantErr    bool
		mockStatus string
	}{
		// TODO: Add test cases.
		{
			name: "Update status from inactive to active",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				status: "active",
				account: &entities.Account{
					ID:       "1",
					Username: "admin1",
					Password: "password1",
					Status:   "inactive", // Set the initial status to "waiting"
				},
			},
			want:       nil,
			wantErr:    false,
			mockStatus: "inactive", // Set the initial status of the mock repository to "waiting"
		},
		{
			name: "Update status from inactive to active (failure)",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				status: "active",
				account: &entities.Account{
					ID:       "2",
					Username: "admin2",
					Password: "password2",
					Status:   "inactive",
				},
			},
			want:       errors.New("update failed"),
			wantErr:    true,
			mockStatus: "inactive",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo.On("UpdateStatus", tt.args.status, tt.args.account).Return(tt.want)
				err := tt.usecase.UpdateStatus(tt.args.status, tt.args.account)
				if (err != nil) != tt.wantErr {
					t.Errorf("useCase.UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(err, tt.want) {
					t.Errorf("useCase.UpdateStatus() = %v, want %v", err, tt.want)
				}
				mockRepo.AssertExpectations(t)
			})
		})
	}
}

func Test_useCase_Delete(t *testing.T) {
	type args struct {
		account *entities.Account
	}
	mockRepo := &mocks.AccountRepositoryInterface{}
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
				accRepo: mockRepo,
			},
			args: args{
				account: &entities.Account{
					Username: "testuser",
					Password: "123",
					Role:     "admin",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "failed: user not found",
			usecase: &useCase{
				accRepo: mockRepo,
			},
			args: args{
				account: &entities.Account{},
			},
			want:    errors.New("error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Delete", tt.args.account).Return(tt.want)
			err := tt.usecase.Delete(tt.args.account)
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
