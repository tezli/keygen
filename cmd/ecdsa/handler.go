package ecdsa

import (
	"github.com/tezli/keygen/pkg/ecdsa"
	"github.com/tezli/keygen/pkg/io"
)

type ecdsaHandler struct {
	writer io.Writer
}

// NewHandler returns a new ECDSA handler
func NewHandler(writer io.Writer) Handler {
	return ecdsaHandler{writer}
}

func (h ecdsaHandler) Invoke(curve string, secret []byte, file string, public bool) error {

	privateKey, publicKey, err := ecdsa.GenerateKeys(curve, []byte(secret), public)
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
