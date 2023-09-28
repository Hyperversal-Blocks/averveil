package sensors

import "github.com/hyperversalblocks/averveil/pkg/signer"

type transcoder struct {
	signer signer.Signer
}

type Transcoder interface {
}

func InitTranscoder() Transcoder {
	return &transcoder{}
}
