package rsa

import (
	"github.com/tezli/keygen/pkg/io"
	"github.com/tezli/keygen/pkg/rsa"
)

type rsaHandler struct {
	writer io.Writer
}

// NewHandler returns a new RSA handler
func NewHandler(writer io.Writer) Handler {
	return &rsaHandler{writer}
}

func (h rsaHandler) Invoke(bits int, secret []byte, file string, public bool) error {

	privateKey, publicKey, err := rsa.GenerateKeys(bits, secret, public)
	if err != nil {
		return err
	}
	if file != "" {
		h.writer.WriteFile(file, privateKey)
		if publicKey != nil {
			h.writer.WriteFile(file, publicKey)
		}
	} else {
		h.writer.Print(string(privateKey))
		if public == true {
			h.writer.Print(string(publicKey))
		}
	}
	return nil
}
