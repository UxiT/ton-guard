package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"time"
)

type ProfileRepository struct {
	db    *sql.DB
	table string
}

type sqlProfile struct {
	UUID         string    `db:"uuid"`
	TelegramID   int       `db:"telegram_id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func NewProfileRepository(db *sql.DB) interfaces.ProfileRepository {
	return &ProfileRepository{
		db:    db,
		table: "profile",
	}
}

func (r *ProfileRepository) FindByUUID(profileUUID valueobject.UUID) (*entity.Profile, error) {
	var profile sqlProfile

	row := r.db.QueryRow("SELECT * FROM profile WHERE uuid = $1", profileUUID.String())
	err := row.Scan(
		&profile.UUID,
		&profile.TelegramID,
		&profile.Email,
		&profile.PasswordHash,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainProfile(profile)
}

func (r *ProfileRepository) FindByTelegramID(telegramID entity.TelegramID) (*entity.Profile, error) {
	var profile sqlProfile

	row := r.db.QueryRow("SELECT * FROM profile WHERE telegram_id = $1", telegramID.Int())
	err := row.Scan(
		&profile.UUID,
		&profile.TelegramID,
		&profile.Email,
		&profile.PasswordHash,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainProfile(profile)
}

func (r *ProfileRepository) Create(profile entity.Profile) error {
	_, err := r.db.Exec(
		"INSERT INTO profile (uuid, telegram_id, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		profile.UUID,
		profile.TelegramID,
		profile.Email,
		profile.PasswordHash,
		time.Now(),
		time.Now(),
	)

	return err
}

func toDomainProfile(c sqlProfile) (*entity.Profile, error) {
	profileUUID, err := valueobject.ParseUUID(c.UUID)
	if err != nil {
		return nil, err
	}

	telegramID, err := entity.NewTelegramID(c.TelegramID)
	if err != nil {
		return nil, err
	}

	email, err := entity.NewEmail(c.Email)
	if err != nil {
		return nil, err
	}

	profile := entity.Profile{
		UUID:         profileUUID,
		TelegramID:   telegramID,
		Email:        email,
		PasswordHash: c.PasswordHash,
	}

	return &profile, nil
}
