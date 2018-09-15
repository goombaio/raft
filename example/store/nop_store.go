// Copyright 2018, gossiper project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the Licenses at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package store

// NopStore is an store.Store that does nothing
type NopStore struct {
	entries map[string]string
}

// NewNopStore ...
func NewNopStore() *NopStore {
	store := &NopStore{
		entries: make(map[string]string, 0),
	}

	return store
}

// Set ...
func (t *NopStore) Set(key string, value string) error {
	return nil
}

// Get ...
func (t *NopStore) Get(key string) (string, error) {
	return "", nil
}

// Delete ...
func (t *NopStore) Delete(key string) error {
	return nil
}
