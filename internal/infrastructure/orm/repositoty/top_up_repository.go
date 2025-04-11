package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"github.com/shopspring/decimal"
	"time"
)

type TopUpRepository struct {
	db    *sql.DB
	table string
}

type sqlTopUp struct {
	UUID          string          `db:"uuid"`
	ProfileUUID   string          `db:"profile_uuid"`
	Amount        decimal.Decimal `db:"amount"`
	Network       string          `db:"network"`
	Status        string          `db:"status"`
	TransactionID *string         `db:"transaction_id"`
	IsClosed      bool            `db:"is_closed"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewTopUpRepository(db *sql.DB) interfaces.TopUpRepository {
	return &TopUpRepository{
		db:    db,
		table: "account",
	}
}

func (r TopUpRepository) Create(topUp entity.TopUp) error {
	_, err := r.db.Exec(
		"INSERT INTO "+r.table+" (uuid, profile_uuid, amount, network, status, transaction_id, is_closed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		topUp.UUID.String(),
		topUp.Customer.String(),
		topUp.Amount,
		topUp.Network,
		topUp.Status,
		topUp.TransactionID,
		topUp.IsClosed,
		topUp.CreatedAt,
		topUp.UpdatedAt,
	)

	return err
}

func (r TopUpRepository) GetByUUID(uuid valueobject.UUID) (*entity.TopUp, error) {
	var topUp sqlTopUp

	row := r.db.QueryRow("SELECT * FROM top_up WHERE uuid = $1", uuid.String())
	err := row.Scan(
		&topUp.UUID,
		&topUp.ProfileUUID,
		&topUp.Amount,
		&topUp.Network,
		&topUp.Status,
		&topUp.TransactionID,
		&topUp.IsClosed,
		&topUp.CreatedAt,
		&topUp.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainTopUp(topUp)
}

func (r TopUpRepository) SetStatus(uuid valueobject.UUID, status entity.TopUpStatus) error {
	_, err := r.db.Exec(
		"UPDATE "+r.table+" SET status = $1, updated_at = $2 WHERE uuid = $3",
		status,
		time.Now(),
		uuid.String(),
	)

	return err
}

func (r TopUpRepository) AddTransactionID(uuid valueobject.UUID, transactionID string) error {
	_, err := r.db.Exec(
		"UPDATE "+r.table+" SET transaction_id = $1, updated_at = $2 WHERE uuid = $3",
		transactionID,
		time.Now(),
		uuid.String(),
	)

	return err
}

func (r TopUpRepository) GetCustomerCurrentTopUp(profileUUID valueobject.UUID) (*entity.TopUp, error) {
	var topUp sqlTopUp

	row := r.db.QueryRow("SELECT * FROM top_up WHERE uuid = $1 and is_closed = false", profileUUID.String())

	err := row.Scan(
		&topUp.UUID,
		&topUp.ProfileUUID,
		&topUp.Amount,
		&topUp.Network,
		&topUp.Status,
		&topUp.TransactionID,
		&topUp.IsClosed,
		&topUp.CreatedAt,
		&topUp.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainTopUp(topUp)
}

func toDomainTopUp(t sqlTopUp) (*entity.TopUp, error) {
	accountUUID, err := valueobject.ParseUUID(t.UUID)
	if err != nil {
		return nil, err
	}

	profileUUID, err := valueobject.ParseUUID(t.ProfileUUID)
	if err != nil {
		return nil, err
	}

	return &entity.TopUp{
		UUID:          accountUUID,
		Customer:      profileUUID,
		Amount:        t.Amount,
		Status:        entity.TopUpStatus(t.Status),
		Network:       t.Network,
		TransactionID: t.TransactionID,
		IsClosed:      t.IsClosed,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}, nil
}
