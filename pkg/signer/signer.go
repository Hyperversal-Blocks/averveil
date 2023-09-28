package signer

import (
	"crypto/cipher"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer interface {
	// EthereumAddress returns the ethereum address this signer uses.
	EthereumAddress() common.Address
	SignTx(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	GetSharedKey(their ecdsa.PublicKey) [32]byte
	GenNonce() []byte
	EncryptAndGetHash(key [32]byte, nonce []byte, message []byte) ([32]byte, []byte, error)
	DecryptMessage(sharedKey [32]byte, cipherText []byte, nonce []byte) (string, error)
	getCipherMode(key []byte) (cipher.AEAD, error)
	VerifySignature(publicKey ecdsa.PublicKey, signature, messageHash []byte) bool
	Sign(hash [32]byte) ([]byte, error)
	GetPublicKey() *ecdsa.PublicKey
	PublicKeyFromBytes(pbKey []byte) (*ecdsa.PublicKey, error)
	BytesFromPublicKey(key *ecdsa.PublicKey) []byte
}

type signer struct {
	publicKey  *ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

func (c *signer) GetPublicKey() *ecdsa.PublicKey {
	return c.publicKey
}

func (c *signer) EthereumAddress() common.Address {
	return crypto.PubkeyToAddress(*c.publicKey)
}

func New(nodePrivateKey string) (Signer, error) {
	if len(strings.TrimSpace(nodePrivateKey)) == 0 {
		return nil, fmt.Errorf("node private key cannot be empty, please generate a new key pair or provide key in config")
	}

	privateKey, err := crypto.HexToECDSA(nodePrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error generating private key from hex: %w ", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("unable to generate public key")
	}

	return &signer{
		publicKey:  publicKeyECDSA,
		privateKey: privateKey,
	}, nil
}

func NewKey() (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("unable to generate a new private key: %w", err)
	}

	return hexutil.Encode(crypto.FromECDSA(privateKey))[2:], nil
}

// SignTx signs an ethereum transaction.
func (c *signer) SignTx(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	txSigner := types.NewLondonSigner(chainID)

	signedTx, err := types.SignTx(transaction, txSigner, c.privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// TODO: Implement BBS signatures for msg arrays:
//  ref: https://identity.foundation/bbs-signature/draft-irtf-cfrg-bbs-signatures.html
//  https://asecuritysite.com/golang/bbs_sign <---- instead use similar logic but funcs from ethereum
