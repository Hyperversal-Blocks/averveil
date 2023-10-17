package sensors

import "averveil/pkg/signer"

type transcoder struct {
	signer signer.Signer
}

type Transcoder interface {
}

func InitTranscoder() Transcoder {
	return &transcoder{}
}
