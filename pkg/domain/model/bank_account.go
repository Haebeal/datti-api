package model

type BankAccount struct {
	UserID      string `json:"uid" gorm:"primarykey"`
	AccountCode string `json:"accountCode"`
	BankCode    string `json:"bankCode"`
	BranchCode  string `json:"branchCode"`
}
