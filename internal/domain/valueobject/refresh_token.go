package valueobject

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type RefreshToken string

func NewRefreshToken() RefreshToken {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(randomBytes)
	randomHash := hex.EncodeToString(hash[:])

	return RefreshToken(randomHash)
}

func (rt RefreshToken) String() string {
	return string(rt)
}

func (rt RefreshToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(rt.String())
}

func (rt *RefreshToken) UnmarshalJSON(data []byte) error {
	var refreshToken string

	if err := json.Unmarshal(data, &refreshToken); err != nil {
		return err
	}

	*rt = RefreshToken(refreshToken)

	return nil
}
