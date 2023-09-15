package signer

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSigner(t *testing.T) {
	t.Parallel()

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	privateKeyToString := hexutil.Encode(crypto.FromECDSA(privateKey))[2:]

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		signer, err := New(privateKeyToString)
		if err != nil {
			t.Fatal("error generating signer: ", err)
		}

		if signer == nil {
			t.Fatal("expected signer to not be nil: ", err)
		}

		ethAddress := signer.EthereumAddress()
		if &ethAddress == nil {
			t.Fatal("expected eth address to not be nil: ", err)
		}
	})

	t.Run("incorrect private key", func(t *testing.T) {
		t.Parallel()

		signer, err := New("someKey")
		if err == nil {
			t.Fatal("expected error: ", err)
		}

		if signer != nil {
			t.Fatal("expected signer to not be nil: ", err)
		}
	})

	t.Run("incorrect private key", func(t *testing.T) {
		t.Parallel()

		signer, err := New("someKey")
		if err == nil {
			t.Fatal("expected error: ", err)
		}

		if signer != nil {
			t.Fatal("expected signer to not be nil: ", err)
		}
	})
}
