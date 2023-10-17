package sensors

import "github.com/hyperversal-blocks/averveil/pkg/signer"

type transcoder struct {
	signer signer.Signer
}

type Transcoder interface {
}

func InitTranscoder() Transcoder {
	return &transcoder{}
}
