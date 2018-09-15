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

package server_test

import (
	"testing"

	"github.com/goombaio/raft/example/server"
	"github.com/goombaio/raft/example/store"
)

type TestServer struct {
	*server.Service
}

// Test_NewServer tests that a server can perform all basic operations.
func Test_NewServer(t *testing.T) {
	simpleStore := store.NewSimpleStore()

	address := ":8000"

	service := &TestServer{server.NewService(address, simpleStore)}
	if service == nil {
		t.Fatalf("Can't the HTTP service")
	}

	if err := service.Start(); err != nil {
		t.Fatalf("failed to start HTTP service: %s", err)
	}
}
