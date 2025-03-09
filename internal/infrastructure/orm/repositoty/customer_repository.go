package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"time"
)

type CustomerRepository struct {
	db    *sql.DB
	table string
}

type sqlCustomer struct {
	UUID        string `db:"uuid"`
	ProfileUUID string `db:"profile_uuid"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewCustomerRepository(db *sql.DB) interfaces.CustomerRepository {
	return &CustomerRepository{
		db:    db,
		table: "customer",
	}
}

func (r *CustomerRepository) FindByProfileUUID(profileUUID valueobject.UUID) (*entity.Customer, error) {
	var customer sqlCustomer

	row := r.db.QueryRow(`SELECT * FROM customer WHERE profile_uuid = $1`, profileUUID.String())

	err := row.Scan(
		&customer.UUID,
		&customer.ProfileUUID,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainCustomer(customer)
}

func (r *CustomerRepository) FindByTelegramID(telegramID entity.TelegramID) (*entity.Customer, error) {
	panic("implement me")
}

func (r *CustomerRepository) Create(customer entity.Customer) error {
	panic("implement me")
}

func toDomainCustomer(c sqlCustomer) (*entity.Customer, error) {
	customerUUID, err := valueobject.ParseUUID(c.UUID)
	if err != nil {
		return nil, err
	}

	profileUUID, err := valueobject.ParseUUID(c.ProfileUUID)
	if err != nil {
		return nil, err
	}

	return &entity.Customer{
		UUID:        customerUUID,
		ProfileUUID: profileUUID,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}, nil
}
