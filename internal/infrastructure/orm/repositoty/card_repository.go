package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type CardRepository struct {
	db    *sql.DB
	table string
}

type sqlCard struct {
	UUID         string    `db:"uuid"`
	ExternalUUID string    `db:"external_uuid"`
	AccountUUID  string    `db:"account_uuid"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{
		db:    db,
		table: "card",
	}
}

func (r *CardRepository) GetByAccount(account uuid.UUID) (*entity.Card, error) {
	var card sqlCard
	row := r.db.QueryRow(
		fmt.Sprintf(`SELECT * FROM %s WHERE account_uuid = $1`, r.table),
		account,
	)

	err := row.Scan(
		&card.UUID,
		&card.ExternalUUID,
		&card.AccountUUID,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainCard(card)
}

func toDomainCard(c sqlCard) (*entity.Card, error) {
	cardUUID, err := uuid.Parse(c.UUID)
	if err != nil {
		return nil, errors.Join(ErrInvalidUUID, err)
	}

	externalUUID, err := uuid.Parse(c.ExternalUUID)
	if err != nil {
		return nil, errors.Join(ErrInvalidUUID, err)
	}

	accountUUID, err := uuid.Parse(c.AccountUUID)
	if err != nil {
		return nil, errors.Join(ErrInvalidUUID, err)
	}

	return &entity.Card{
		UUID:         cardUUID,
		ExternalUUID: externalUUID,
		AccountUUID:  accountUUID,
	}, nil
}
