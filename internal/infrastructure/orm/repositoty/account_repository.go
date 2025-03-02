package repositoty

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type AccountRepository struct {
	db    *sql.DB
	table string
}

type sqlAccount struct {
	UUID         string          `db:"uuid"`
	ExternalUUID string          `db:"external_uuid"`
	CustomerUUID string          `db:"customer_uuid"`
	Currency     string          `db:"currency"`
	Status       string          `db:"status"`
	Balance      decimal.Decimal `db:"balance"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    time.Time       `db:"updated_at"`
}

//func NewAccountRepository(db *sql.DB) *AccountRepository {
//	return &AccountRepository{
//		db:    db,
//		table: "account",
//	}
//}
//
//func (r *AccountRepository) GetByCustomer(customer uuid.UUID) (*aggregate.Account, error) {
//	var account sqlAccount
//
//	row := r.db.QueryRow(
//		fmt.Sprintf("select * from %s where customer_uuid = $1", r.table),
//		customer,
//	)
//
//	err := row.Scan(
//		&account.UUID,
//		&account.ExternalUUID,
//		&account.CustomerUUID,
//		&account.Currency,
//		&account.Status,
//		&account.Balance,
//		&account.CreatedAt,
//		&account.UpdatedAt,
//	)
//
//	if err != nil {
//		return nil, err
//	}
//}
//
//func toDomainAccount(a sqlAccount) (*aggregate.Account, error) {
//	accountUUID, err := uuid.Parse(a.UUID)
//	if err != nil {
//		return nil, errors.Join(ErrInvalidUUID, err)
//	}
//
//	externalUUID, err := uuid.Parse(a.ExternalUUID)
//	if err != nil {
//		return nil, errors.Join(ErrInvalidUUID, err)
//	}
//
//	customerUUID, err := uuid.Parse(a.CustomerUUID)
//	if err != nil {
//		return nil, errors.Join(ErrInvalidUUID, err)
//	}
//
//	return
//}
