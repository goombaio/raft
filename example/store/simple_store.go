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

package store

import (
	"errors"
	"sync"
)

// SimpleStore ...
type SimpleStore struct {
	mu      sync.Mutex
	entries map[string]string
}

// NewSimpleStore ...
func NewSimpleStore() *SimpleStore {
	store := &SimpleStore{
		entries: make(map[string]string, 0),
	}

	return store
}

// Get ...
func (s *SimpleStore) Get(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if value, ok := s.entries[key]; ok {
		return value, nil
	}

	return "", errors.New("Key %s not found")
}

// Set ...
func (s *SimpleStore) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.entries[key] = value

	return nil
}

// Delete ...
func (s *SimpleStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.entries[key]; ok {
		delete(s.entries, key)
		return nil
	}

	return errors.New("Key %s not found")
}
