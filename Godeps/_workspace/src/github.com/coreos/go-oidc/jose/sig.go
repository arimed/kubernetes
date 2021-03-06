package jose

import (
	"fmt"
	"strings"
)

type Verifier interface {
	ID() string
	Alg() string
	Verify(sig []byte, data []byte) error
}

type Signer interface {
	Verifier
	Sign(data []byte) (sig []byte, err error)
}

func NewVerifier(jwk JWK) (Verifier, error) {
	if strings.ToUpper(jwk.Type) != "RSA" {
		return nil, fmt.Errorf("unsupported key type %q", jwk.Type)
	}

	return NewVerifierRSA(jwk)
}
