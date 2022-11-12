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
	"testing"

	"github.com/tezli/keygen/pkg/io"
)

func TestHandler(t *testing.T) {
	writer := io.FakeWriter{}
	handler := NewHandler(writer)
	err := handler.Invoke(2048, []byte("secret"), "", false)
	if err != nil {
		t.Fail()
	}
}

func TestHandlerWithFile(t *testing.T) {
	writer := io.FakeWriter{}
	handler := NewHandler(writer)
	err := handler.Invoke(2048, []byte("secret"), "foo", false)
	if err != nil {
		t.Fail()
	}
}

func TestHandlerPublic(t *testing.T) {
	writer := io.FakeWriter{}
	handler := NewHandler(writer)
	err := handler.Invoke(2048, []byte("secret"), "", true)
	if err != nil {
		t.Fail()
	}
}

func TestHandlerWithFilePublic(t *testing.T) {
	writer := io.FakeWriter{}
	handler := NewHandler(writer)
	err := handler.Invoke(2048, []byte("secret"), "foo", true)
	if err != nil {
		t.Fail()
	}
}

func TestHandlerInvalidBits(t *testing.T) {
	writer := io.FakeWriter{}
	handler := NewHandler(writer)
	err := handler.Invoke(1024, []byte("secret"), "", false)
	if err == nil {
		t.Fail()
	}
}
