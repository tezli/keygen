/*
Copyright © 2022 Robert Tezli robert.tezli+github@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// GenerateKeys provides rsa Wrappers
func GenerateKeys(bits int, secret []byte, createPublicKey bool) ([]byte, []byte, error) {
	if bits < 2048 {
		return nil, nil, errors.New("Minimal RSA key length is 2048 bits")
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	PKCS8PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: PKCS8PrivateKey}
	privateKeyPEM := pem.EncodeToMemory(&privateKeyBlock)

	var publicKeyPEM []byte

	if createPublicKey {
		publicKey := privateKey.PublicKey
		publicKeyPKCS1 := x509.MarshalPKCS1PublicKey(&publicKey)
		publicKeyBlock := pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyPKCS1}
		publicKeyPEM = pem.EncodeToMemory(&publicKeyBlock)
	}

	return privateKeyPEM, publicKeyPEM, nil
}
