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

package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const (
	CURVE_P224 = "P224"
	CURVE_P256 = "P256"
	CURVE_P384 = "P384"
	CURVE_P521 = "P521"
)

// GenerateKeys provides ecdsa wrappers
func GenerateKeys(curve string, secret []byte, createPublicKey bool) ([]byte, []byte, error) {

	var c elliptic.Curve
	switch curve {
	case CURVE_P224:
		c = elliptic.P224()
	case CURVE_P256:
		c = elliptic.P256()
	case CURVE_P384:
		c = elliptic.P384()
	case CURVE_P521:
		c = elliptic.P521()
	default:
		return nil, nil, errors.New("invalid curve value")
	}

	key, _ := ecdsa.GenerateKey(c, rand.Reader)
	PKCS8PrivateKey, _ := x509.MarshalECPrivateKey(key)
	privateKeyBlock := pem.Block{Type: "PRIVATE KEY", Bytes: PKCS8PrivateKey}
	privateKeyPEM := pem.EncodeToMemory(&privateKeyBlock)

	var publicKeyPEM []byte

	if createPublicKey {
		publicKey := key.PublicKey
		PKCS1PublicKey, _ := x509.MarshalPKIXPublicKey(&publicKey)
		publicKeyBlock := pem.Block{Type: "PUBLIC KEY", Bytes: PKCS1PublicKey}
		publicKeyPEM = pem.EncodeToMemory(&publicKeyBlock)
	}
	return privateKeyPEM, publicKeyPEM, nil
}
