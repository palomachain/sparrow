package client

import (
	"github.com/vizualni/whoops"
)

const (
	ErrSignatureVerificationFailed              = whoops.String("signature verification failed")
	ErrSignatureDoesNotMatchItsRespectiveSigner = whoops.String("signature does not match its respective signer")
	ErrTooLittleOrTooManySignaturesProvided     = whoops.String("too many or too little signatures provided")
)
