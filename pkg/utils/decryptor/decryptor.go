package decryptor

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/rs/zerolog/log"
	"strings"
)

type Decryptor struct {
	privateKey *rsa.PrivateKey
}

func NewDecryptor(privateKey *rsa.PrivateKey) *Decryptor {
	return &Decryptor{privateKey: privateKey}
}

func (d Decryptor) Decrypt(encryptedMessage string) (string, error) {
	message := strings.TrimPrefix(encryptedMessage, "-----BEGIN CardNumber MESSAGE-----")
	message = strings.TrimSuffix(message, "-----END CardNumber MESSAGE-----")
	message = strings.ReplaceAll(message, "\n", "")

	message = "HTNR+gNvL1z+ZfWG21j+3baRqrC3Q/ZZZNlBg+OX9UGuIwQU7t9Wyz9shWFmm7jgfZ0Rz2AMwO70fBQyCfJXR+5aomCQYkNhR0TYOZfNmSfhzaUy12z7olUrntwMlVukFDqmtnFQL+mSEDc0pNV1ubPxy/IAQbMOTiQzdOxTpA01ScrsUeReajCueyBSqMV5PQkaQtDxKSI73k8HYxddzQX01I9+Xnr2XbkuAFusUp1UC5coGgIMxgcjtn8BNQ0dgjH0uHVrxEhNCA+hkHTnLusjoBr+Ut6b9c5k81R41eBt5sEUeemckv4qMcx30gbiArolaEpxstCeGofGU3rrXomyQWnd+rdkwyyL7QkUskFGshsFWLBSHQya0/m4i5cwC4mHbJHWIw8+X6QoPUW9u+bni0r3IdADJMn3UGNAH619cf7tj/wzMYkoonIRrIKmOmzi0OOHcLwRWpQeWYrdkE9fzfrH3CVLgep+TMsaiyKXGRakFBvENYGYdOqTzqT0gfoXPlb4BG4tolaZTLWJVmQiLLT3j0kigP6KBddGI2OqdViguFT2wuUk3BIm1SMQdY226Q0fGggm3ZqDwTna2E2TyWeG2ZR58R7HEN7YD6ec/CyExO9VsInK31/4BObrKRZSuNE56gYDJHQf35TJD0N0T8p61uT45krH1I0/9TI="

	r, err := d.privateKey.Decrypt(rand.Reader, []byte(message), nil)
	if err != nil {
		panic(err)
	}

	log.Printf("%s", r)

	return string(r), nil
}
