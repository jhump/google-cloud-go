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

// [START appengine_generated_appengine_apiv1_VersionsClient_ListVersions]

package main

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	"google.golang.org/api/iterator"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func main() {
	// import appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := appengine.NewVersionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &appenginepb.ListVersionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListVersions(ctx, req)
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

// [END appengine_generated_appengine_apiv1_VersionsClient_ListVersions]
