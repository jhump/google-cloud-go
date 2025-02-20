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

// [START cloudresourcemanager_generated_resourcemanager_apiv2_FoldersClient_SearchFolders]

package main

import (
	"context"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv2"
	"google.golang.org/api/iterator"
	resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v2"
)

func main() {
	// import resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v2"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := resourcemanager.NewFoldersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &resourcemanagerpb.SearchFoldersRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchFolders(ctx, req)
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

// [END cloudresourcemanager_generated_resourcemanager_apiv2_FoldersClient_SearchFolders]
