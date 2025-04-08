package decryptor

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"decard/internal/domain"
	"encoding/base64"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
)

type Decryptor struct {
	logger     *zerolog.Logger
	privateKey *rsa.PrivateKey
}

func NewDecryptor(privateKey *rsa.PrivateKey, logger *zerolog.Logger) *Decryptor {
	return &Decryptor{
		privateKey: privateKey,
		logger:     logger,
	}
}

func (d Decryptor) Decrypt(b64EncodedMessage string) (string, error) {
	message := strings.ReplaceAll(b64EncodedMessage, "-----BEGIN CardNumber MESSAGE-----", "")
	message = strings.ReplaceAll(message, "-----END CardNumber MESSAGE-----", "")
	message = strings.ReplaceAll(message, "\n", "")

	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		d.logger.Error().Err(err).Str("message", b64EncodedMessage).Msg("failed to decrypt message")

		return "", domain.ErrInternal
	}

	r, err := d.privateKey.Decrypt(rand.Reader, data, &rsa.OAEPOptions{Hash: crypto.SHA256, Label: []byte("CardNumber")})

	if err != nil {
		panic(err)
	}

	log.Printf("%s", r)

	return string(r), nil
}
