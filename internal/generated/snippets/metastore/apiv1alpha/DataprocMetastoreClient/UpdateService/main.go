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

// [START metastore_generated_metastore_apiv1alpha_DataprocMetastoreClient_UpdateService]

package main

import (
	"context"

	metastore "cloud.google.com/go/metastore/apiv1alpha"
	metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"
)

func main() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.UpdateServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END metastore_generated_metastore_apiv1alpha_DataprocMetastoreClient_UpdateService]
