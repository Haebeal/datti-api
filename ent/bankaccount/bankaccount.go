// Code generated by ent, DO NOT EDIT.

package bankaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bankaccount type in the database.
	Label = "bank_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountCode holds the string denoting the account_code field in the database.
	FieldAccountCode = "account_code"
	// FieldBankCode holds the string denoting the bank_code field in the database.
	FieldBankCode = "bank_code"
	// FieldBranchCode holds the string denoting the branch_code field in the database.
	FieldBranchCode = "branch_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the bankaccount in the database.
	Table = "bank_accounts"
)

// Columns holds all SQL columns for bankaccount fields.
var Columns = []string{
	FieldID,
	FieldAccountCode,
	FieldBankCode,
	FieldBranchCode,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AccountCodeValidator is a validator for the "account_code" field. It is called by the builders before save.
	AccountCodeValidator func(string) error
	// BankCodeValidator is a validator for the "bank_code" field. It is called by the builders before save.
	BankCodeValidator func(string) error
	// BranchCodeValidator is a validator for the "branch_code" field. It is called by the builders before save.
	BranchCodeValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the BankAccount queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountCode orders the results by the account_code field.
func ByAccountCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountCode, opts...).ToFunc()
}

// ByBankCode orders the results by the bank_code field.
func ByBankCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankCode, opts...).ToFunc()
}

// ByBranchCode orders the results by the branch_code field.
func ByBranchCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBranchCode, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}