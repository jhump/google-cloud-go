// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START pubsublite_generated_pubsublite_apiv1_CursorClient_ListPartitionCursors]

package main

import (
	"context"

	pubsublite "cloud.google.com/go/pubsublite/apiv1"
	"google.golang.org/api/iterator"
	pubsublitepb "google.golang.org/genproto/googleapis/cloud/pubsublite/v1"
)

func main() {
	// import pubsublitepb "google.golang.org/genproto/googleapis/cloud/pubsublite/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := pubsublite.NewCursorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsublitepb.ListPartitionCursorsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPartitionCursors(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

// [END pubsublite_generated_pubsublite_apiv1_CursorClient_ListPartitionCursors]
