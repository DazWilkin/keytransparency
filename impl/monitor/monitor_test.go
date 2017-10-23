// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package monitor implements the monitor service. A monitor repeatedly polls a
// key-transparency server's Mutations API and signs Map Roots if it could
// reconstruct
// clients can query.
package monitor

import (
	"context"
	"testing"

	"github.com/google/keytransparency/core/monitor/storage"
)

func TestGetSignedMapRoot(t *testing.T) {
	srv := New(storage.New())
	_, err := srv.GetSignedMapRoot(context.TODO(), nil)
	if got, want := err, ErrNothingProcessed; got != want {
		t.Errorf("GetSignedMapRoot(_, _): %v, want %v", got, want)
	}
}