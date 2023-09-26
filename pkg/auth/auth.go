package auth

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"

	"github.com/hyperversalblocks/averveil/pkg/jwt"
	"github.com/hyperversalblocks/averveil/pkg/signer"
)

type auth struct {
	Signer signer.Signer
	Client *badger.DB
	JWT    jwt.JWT
}

type Challenge struct {
	Key     string
	Message string
	Nonce   string
}

type StoredChallenge struct {
	Ciphered string `json:"ciphered"`
	Nonce    string `json:"nonce"`
}

func New(signer signer.Signer, jwt jwt.JWT) {

}

type Auth interface {
	GetChallenge(context.Context, []byte) (*Challenge, error)
	VerifyChallenge(ctx context.Context, deciphered string) (bool, error)
}

// GetChallenge by generating a shared key from public key of user and private key of node
func (a *auth) GetChallenge(ctx context.Context, publicKey []byte) (*Challenge, error) {
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

func (a *auth) VerifyChallenge(ctx context.Context, deciphered string, pubKey []byte) (bool, error) {
	obj := a.Client.Get(ctx, hex.EncodeToString(pubKey))

	challenge := &StoredChallenge{}

	err := json.Unmarshal(obj.([]byte), challenge)
	if err != nil {
		return false, fmt.Errorf("unable to unmarshal stored object: %w", err)
	}

	publicKey, err := a.Signer.PublicKeyFromBytes(pubKey)
	if err != nil {
		return false, fmt.Errorf("unable to generate a challenge: %w", err)
	}

	sharedKey := a.Signer.GetSharedKey(*publicKey)

	decodedCipher, err := hex.DecodeString(challenge.Ciphered)
	if err != nil {
		return false, fmt.Errorf("err while decoding cipher: %w", err)
	}

	decodedNonce, err := hex.DecodeString(challenge.Nonce)
	if err != nil {
		return false, fmt.Errorf("err while decoding nonce: %w", err)
	}

	message, err := a.Signer.DecryptMessage(sharedKey, decodedCipher, decodedNonce)
	if err != nil {
		return false, fmt.Errorf("unable to decrypt message with key: %w", err)
	}

	if strings.Compare(deciphered, message) != 0 {
		return false, nil
	}

	return true, nil
}
