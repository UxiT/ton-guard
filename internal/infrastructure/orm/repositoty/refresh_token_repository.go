package repositoty

import (
	"database/sql"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"fmt"
	"time"
)

type RefreshTokenRepository struct {
	db              *sql.DB
	table           string
	refreshTokenTTL time.Duration
}

type sqlRefreshToken struct {
	UUID        string
	ProfileUUID string

	ExpiresAt time.Time
	DeletedAt sql.NullTime
	CreatedAt time.Time
}

func NewRefreshTokenRepository(db *sql.DB, tokenTTL time.Duration) interfaces.RefreshTokenRepository {
	return RefreshTokenRepository{
		db:              db,
		table:           "refresh_token",
		refreshTokenTTL: tokenTTL,
	}
}

func (r RefreshTokenRepository) FindByToken(token string) (*entity.RefreshToken, error) {
	var rToken sqlRefreshToken

	row := r.db.QueryRow("select * from refresh_token where uuid = $1", token)
	err := row.Scan(
		rToken.UUID,
		rToken.ProfileUUID,
		rToken.ExpiresAt,
		rToken.DeletedAt,
		&rToken.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainRefreshToken(rToken)
}

func (r RefreshTokenRepository) GetLastForProfile(profileUUID valueobject.UUID) (*entity.RefreshToken, error) {
	var rToken sqlRefreshToken

	row := r.db.QueryRow("select * from refresh_token where profile_uuid = $1 and deleted_at is null", profileUUID.String())
	err := row.Scan(
		rToken.UUID,
		rToken.ProfileUUID,
		rToken.ExpiresAt,
		rToken.DeletedAt,
		&rToken.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return toDomainRefreshToken(rToken)
}

func (r RefreshTokenRepository) Delete(token valueobject.RefreshToken) error {
	_, err := r.db.Exec("update refresh_token set deleted_at = now() where uuid = $1", token.String())

	return err
}

func (r RefreshTokenRepository) Create(profileUUID valueobject.UUID) error {
	timeNow := time.Now()
	expiresAt := timeNow.Add(r.refreshTokenTTL)
	refreshToken := valueobject.NewRefreshToken()

	_, err := r.db.Exec(
		fmt.Sprintf("insert into %s (uuid, profile_uuid, expires_at, created_at) values ($1, $2, $3, $4)", r.table),
		refreshToken.String(),
		profileUUID.String(),
		expiresAt,
		timeNow,
	)

	return err
}

func toDomainRefreshToken(t sqlRefreshToken) (*entity.RefreshToken, error) {
	profileUUID, err := valueobject.ParseUUID(t.ProfileUUID)
	if err != nil {
		return nil, err
	}

	return &entity.RefreshToken{
		UUID:        valueobject.RefreshToken(t.UUID),
		ProfileUUID: profileUUID,
	}, nil
}
