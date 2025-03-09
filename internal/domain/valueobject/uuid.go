package valueobject

import (
	"encoding/json"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func NewUUID() UUID {
	return UUID(uuid.New())
}

func ParseUUID(s string) (UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}

	return UUID(parsed), nil
}

func (u UUID) String() string {
	return uuid.UUID(u).String()
}

func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(u).String())
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	var uuidStr string
	if err := json.Unmarshal(data, &uuidStr); err != nil {
		return err
	}

	baseUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return err
	}

	*u = UUID(baseUUID)

	return nil
}
