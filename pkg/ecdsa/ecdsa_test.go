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
	"testing"
)

func TestECDSAWithSecret(t *testing.T) {
	private, public, err := GenerateKeys(CURVE_P256, []byte("secret"), true)
	if private == nil || public == nil || err != nil {
		t.Fail()
	}
}

func TestECDSAWithoutSecret(t *testing.T) {
	private, public, err := GenerateKeys(CURVE_P256, nil, false)
	if private == nil || public != nil || err != nil {
		t.Fail()
	}
}

func TestECDSAWithNilCurve(t *testing.T) {
	_, _, err := GenerateKeys("FOO", nil, false)
	if err == nil {
		t.Fail()
	}
}
