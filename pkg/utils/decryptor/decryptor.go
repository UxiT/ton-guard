package decryptor

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"decard/internal/domain"
	"encoding/base64"
	"github.com/rs/zerolog"
	"regexp"
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
	re, err := regexp.Compile(`-----BEGIN (\w+) MESSAGE-----\n([\s\S]*?)\n-----END \1 MESSAGE-----`)
	if err != nil {
		d.logger.Error().Err(err).Msg("Failed to compile regex")

		return "", err
	}

	matches := re.FindStringSubmatch(b64EncodedMessage)
	label := matches[1]

	message := strings.TrimSpace(matches[2])
	message = strings.ReplaceAll(message, "\n", "")

	decodedMessage, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		d.logger.Error().Err(err).Str("message", b64EncodedMessage).Msg("failed to decode message")

		return "", domain.ErrInternal
	}

	r, err := d.privateKey.Decrypt(rand.Reader, decodedMessage, &rsa.OAEPOptions{
		Hash:  crypto.SHA256,
		Label: []byte(label),
	})

	if err != nil {
		d.logger.Error().
			Err(err).
			Str("message", message).
			Str("label", label).
			Msg("failed to decrypt message")

		return "", err
	}

	return string(r), nil
}
