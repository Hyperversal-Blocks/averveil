package auth

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/hyperversal-blocks/averveil/pkg/jwt"
	"github.com/hyperversal-blocks/averveil/pkg/signer"
	"github.com/hyperversal-blocks/averveil/pkg/store"
	"github.com/hyperversal-blocks/averveil/pkg/user"
)

type auth struct {
	Signer      signer.Signer
	Store       store.Store
	JWT         jwt.JWT
	userService user.Service
}

func New(signer signer.Signer, store store.Store, JWT jwt.JWT, userService user.Service) Auth {
	return &auth{Signer: signer, Store: store, JWT: JWT, userService: userService}
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

type Auth interface {
	GetChallenge(publicKey []byte) (*Challenge, error)
	VerifyChallenge(deciphered string, pubKey []byte) (bool, error)
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

	// TODO: add ttl
	obj, err := json.Marshal(&StoredChallenge{
		Ciphered: hex.EncodeToString(cipheredText),
		Nonce:    hex.EncodeToString(nonce),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to marshal challenge: %w", err)
	}

	err = a.Store.Put(hex.EncodeToString(publicKey), obj)
	if err != nil {
		return nil, fmt.Errorf("unable to store challenge: %w", err)
	}

	return &Challenge{
		Key:     hex.EncodeToString(a.Signer.BytesFromPublicKey(a.Signer.GetPublicKey())),
		Message: hex.EncodeToString(cipheredText),
		Nonce:   hex.EncodeToString(nonce),
	}, nil
}

func (a *auth) VerifyChallenge(deciphered string, pubKey []byte) (bool, error) {
	obj, err := a.Store.Get(hex.EncodeToString(pubKey))
	if err != nil {
		return false, fmt.Errorf("unable to store challenge: %w", err)
	}

	challenge := &StoredChallenge{}

	err = json.Unmarshal(obj, challenge)
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
