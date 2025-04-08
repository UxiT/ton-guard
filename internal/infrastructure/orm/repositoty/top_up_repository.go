package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	domaintype "decard/internal/domain/type"
	"decard/internal/domain/valueobject"
	"github.com/shopspring/decimal"
	"time"
)

type TopUpRepository struct {
	db    *sql.DB
	table string
}

type sqlTopUp struct {
	UUID        string          `db:"uuid"`
	ProfileUUID string          `db:"profile_uuid"`
	Amount      decimal.Decimal `db:"amount"`
	Network     string          `db:"network"`
	Status      string          `db:"status"`
	IsClosed    bool            `db:"is_closed"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewTopUpRepository(db *sql.DB) interfaces.TopUpRepository {
	return &TopUpRepository{
		db:    db,
		table: "account",
	}
}

func (r *TopUpRepository) Create(topUp entity.TopUp) error {
	_, err := r.db.Exec(
		"INSERT INTO "+r.table+" (uuid, profile_uuid, amount, network, status, is_closed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		topUp.UUID.String(),
		topUp.Profile.String(),
		topUp.Amount,
		topUp.Network,
		topUp.Status,
		topUp.IsClosed,
		topUp.CreatedAt,
		topUp.UpdatedAt,
	)

	return err
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

	balance, err := entity.NewBalance(a.Balance.String())
	if err != nil {
		return nil, err
	}

	return &entity.TopUp{
		UUID:      accountUUID,
		Profile:   profileUUID,
		Amount:    domaintype.Currency(a.Currency),
		Status:    domaintype.AccountStatus(a.Status),
		Balance:   *balance,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}, nil
}
