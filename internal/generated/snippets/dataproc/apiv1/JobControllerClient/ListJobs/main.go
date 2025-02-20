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

// [START dataproc_generated_dataproc_apiv1_JobControllerClient_ListJobs]

package main

import (
	"context"

	dataproc "cloud.google.com/go/dataproc/apiv1"
	"google.golang.org/api/iterator"
	dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
)

func main() {
	// import dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := dataproc.NewJobControllerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &dataprocpb.ListJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListJobs(ctx, req)
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

// [END dataproc_generated_dataproc_apiv1_JobControllerClient_ListJobs]
