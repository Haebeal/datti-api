package repository

import (
	"context"

	"github.com/datti-api/pkg/domain/model"
)

type BankAccountRepository interface {
	UpsertBankAccount(c context.Context, uid string, accountCode string, bankCode string, branchCode string) (string, error)
	GetBankAccountById(c context.Context, uid string) (*model.BankAccount, error)
	DeleteBankAccount(c context.Context, uid string) (*model.BankAccount, error)
}
