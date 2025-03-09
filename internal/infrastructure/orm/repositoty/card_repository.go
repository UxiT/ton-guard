package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/valueobject"
	"fmt"
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

func (r *CardRepository) GetByAccount(account valueobject.UUID) (*entity.Card, error) {
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
	cardUUID, err := valueobject.ParseUUID(c.UUID)
	if err != nil {
		return nil, err
	}

	externalUUID, err := valueobject.ParseUUID(c.ExternalUUID)
	if err != nil {
		return nil, err
	}

	accountUUID, err := valueobject.ParseUUID(c.AccountUUID)
	if err != nil {
		return nil, err
	}

	return &entity.Card{
		UUID:         cardUUID,
		ExternalUUID: externalUUID,
		AccountUUID:  accountUUID,
	}, nil
}
