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
	"bytes"
	"testing"
)

type FakeHandler struct{}

func (m FakeHandler) Invoke(curve string, secret []byte, file string, public bool) error {
	return nil
}

func TestCommand(t *testing.T) {
	t.Parallel()
	handler := FakeHandler{}
	command := Command(handler)

	b := bytes.NewBufferString("")
	command.SetOut(b)

	err := command.Execute()
	if err == nil {
		t.Fail()
	}

	command.Flags().Set("curve", "P256")

	err = command.Execute()
	if err != nil {
		t.Fail()
	}
}
