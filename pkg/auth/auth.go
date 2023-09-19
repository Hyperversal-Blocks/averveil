package auth

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/hyperversalblocks/averveil/pkg/signer"
)

type auth struct {
	Signer signer.Signer
}

type Challenge struct {
	Key     string
	Message string
	Nonce   string
}

type Auth interface {
	GetChallenge(string) (*Challenge, error)
}

// GetChallenge by generating a shared key from public key of user and private key of node
func (a *auth) GetChallenge(publicKey []byte) (*Challenge, error) {
	pubKey, err := a.Signer.PublicKeyFromBytes(publicKey)
	if err != nil {
		return nil, fmt.Errorf("unable to generate a challenge: %w", err)
	}

	sharedKey := a.Signer.GetSharedKey(*pubKey)
	nonce := a.Signer.GenNonce()

	_, cipheredText, err := a.Signer.EncryptAndGetHash(sharedKey, nonce, []byte(uuid.NewString()))
	if err != nil {
		return nil, fmt.Errorf("unable to encrypt and get hash: %w", err)
	}

	return &Challenge{
		Key:     hex.EncodeToString(a.Signer.BytesFromPublicKey(a.Signer.GetPublicKey())),
		Message: hex.EncodeToString(cipheredText),
		Nonce:   hex.EncodeToString(nonce),
	}, nil
}

func (a *auth) VerifyChallenge(challenge *Challenge) (bool, error) {
	sharedKey, err := hex.DecodeString(challenge.Key)
	if err != nil {
		return false, fmt.Errorf("unable to decode: %w", err)
	}

	cipheredText, err := hex.DecodeString(challenge.Message)
	if err != nil {
		return false, fmt.Errorf("unable to decode text: %w", err)
	}

	nonce, err := hex.DecodeString(challenge.Nonce)
	if err != nil {
		return false, fmt.Errorf("unable to decode nonce: %w", err)
	}

	message, err := a.Signer.DecryptMessage([32]byte(sharedKey), cipheredText, nonce)
	if err != nil {
		return false, fmt.Errorf("unable to decrypt message with key: %w", err)
	}

	if strings.Compare(challenge.Message, message) != 0 {
		return false, nil
	}

	return true, nil
}
