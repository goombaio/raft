// Copyright 2018, gossiper project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package store_test

import (
	"testing"

	"github.com/goombaio/raft/example/store"
)

func TestNopStore(t *testing.T) {
	_ = store.NewNopStore()
}

func TestNopStore_Set(t *testing.T) {
	nopStore := store.NewNopStore()

	err := nopStore.Set("key", "value")
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestNopStore_Get(t *testing.T) {
	nopStore := store.NewNopStore()

	value, err := nopStore.Get("key")
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if value != "" {
		t.Fatalf("Expected value %q but got %q", "", value)
	}
}

func TestNopStore_Delete(t *testing.T) {
	nopStore := store.NewNopStore()

	err := nopStore.Delete("key")
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}
