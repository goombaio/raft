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

package server

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/goombaio/raft/example/store"
)

// Service ...
type Service struct {
	address  string
	listener net.Listener
	store    store.Storer
}

// NewService ...
func NewService(address string, store store.Storer) *Service {
	service := &Service{
		address: address,
		store:   store,
	}

	return service
}

// Start ...
func (s *Service) Start() error {
	server := http.Server{
		Handler: s,
	}

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.listener = listener

	http.Handle("/", s)

	go func() {
		err := server.Serve(s.listener)
		if err != nil {
			log.Fatalf("HTTP serve: %s", err)
		}
	}()

	return nil
}

// Close ...
func (s *Service) Close() {
	s.listener.Close()
}

// ServeHTTP ...
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// - Use a Mux & Router
	if strings.HasPrefix(r.URL.Path, "/key") {
		s.handleKeyRequest(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// handleKeyRequest ...
func (s *Service) handleKeyRequest(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// - Use a router to better handling url parsing/params
	// - Decuple this handler into 3, one for each HTTP Method
	getKey := func() string {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			return ""
		}
		return parts[2]
	}

	switch r.Method {
	case "GET":
		k := getKey()
		if k == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		v, err := s.store.Get(k)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(map[string]string{k: v})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = io.WriteString(w, string(b))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "POST":
		// Read the value from the POST body.
		m := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for k, v := range m {
			if err := s.store.Set(k, v); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	case "DELETE":
		k := getKey()
		if k == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Delete(k); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err := s.store.Delete(k)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Addr returns the address on which the Service is listening
func (s *Service) Addr() net.Addr {
	return s.listener.Addr()
}
