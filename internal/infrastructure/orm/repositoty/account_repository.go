package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	domaintype "decard/internal/domain/type"
	"decard/internal/domain/valueobject"
	"fmt"
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
	Currency     string          `db:"currency"`
	Status       string          `db:"status"`
	Balance      decimal.Decimal `db:"balance"`
	CustomerUUID string          `db:"customer_uuid"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewAccountRepository(db *sql.DB) interfaces.AccountRepository {
	return &AccountRepository{
		db:    db,
		table: "account",
	}
}

func (r *AccountRepository) GetByCustomer(customer valueobject.UUID) (*entity.Account, error) {
	var account sqlAccount

	row := r.db.QueryRow(
		fmt.Sprintf("select * from %s where customer_uuid = $1", r.table),
		customer.String(),
	)

	err := row.Scan(
		&account.UUID,
		&account.ExternalUUID,
		&account.Currency,
		&account.Status,
		&account.Balance,
		&account.CustomerUUID,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainAccount(account)
}

func (r *AccountRepository) Create(account entity.Account) error {
	_, err := r.db.Exec(
		"INSERT INTO "+r.table+" (uuid, external_uuid, currency, status, balance, customer_uuid, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		account.UUID.String(),
		account.ExternalUUID.String(),
		account.Currency,
		account.Status,
		account.Balance,
		account.CustomerUUID,
		account.CreatedAt,
		account.UpdatedAt,
	)

	return err
}

func toDomainAccount(a sqlAccount) (*entity.Account, error) {
	accountUUID, err := valueobject.ParseUUID(a.UUID)
	if err != nil {
		return nil, err
	}

	externalUUID, err := valueobject.ParseUUID(a.ExternalUUID)
	if err != nil {
		return nil, err
	}

	balance, err := entity.NewBalance(a.Balance.String())
	if err != nil {
		return nil, err
	}

	return &entity.Account{
		UUID:         accountUUID,
		ExternalUUID: externalUUID,
		Currency:     domaintype.Currency(a.Currency),
		Status:       domaintype.AccountStatus(a.Status),
		Balance:      *balance,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}, nil
}
