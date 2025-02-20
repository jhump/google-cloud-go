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

// [START cloudscheduler_generated_scheduler_apiv1_CloudSchedulerClient_DeleteJob]

package main

import (
	"context"

	scheduler "cloud.google.com/go/scheduler/apiv1"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1"
)

func main() {
	ctx := context.Background()
	c, err := scheduler.NewCloudSchedulerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &schedulerpb.DeleteJobRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END cloudscheduler_generated_scheduler_apiv1_CloudSchedulerClient_DeleteJob]
