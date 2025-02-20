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

// [START datastore_generated_datastore_admin_apiv1_DatastoreAdminClient_GetIndex]

package main

import (
	"context"

	admin "cloud.google.com/go/datastore/admin/apiv1"
	adminpb "google.golang.org/genproto/googleapis/datastore/admin/v1"
)

func main() {
	// import adminpb "google.golang.org/genproto/googleapis/datastore/admin/v1"

	ctx := context.Background()
	c, err := admin.NewDatastoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.GetIndexRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIndex(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END datastore_generated_datastore_admin_apiv1_DatastoreAdminClient_GetIndex]
