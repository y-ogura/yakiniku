package usecase_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/y-ogura/yakiniku/account"
	"github.com/y-ogura/yakiniku/account/repository/mocks"
	"github.com/y-ogura/yakiniku/account/usecase"
)

func TestList(t *testing.T) {
	mockAccounts := []*account.Account{
		&account.Account{ID: uuid.NewV4(), Name: "mockAccount", Email: "test@example.com", Password: "asdf0987"},
	}
	accountRepo := new(mocks.AccountRepository)
	accountRepo.On("List").Return(mockAccounts, nil)
	defer accountRepo.AssertCalled(t, "List")

	expect := []*account.Client{
		&account.Client{ID: mockAccounts[0].ID, Name: "mockAccount", Email: "test@example.com"},
	}

	u := usecase.NewAccountUsecase(accountRepo)
	res, err := u.List()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expect, res)
}
