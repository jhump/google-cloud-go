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

// [START cloudtasks_generated_cloudtasks_apiv2beta3_Client_DeleteTask]

package main

import (
	"context"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2beta3"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
)

func main() {
	ctx := context.Background()
	c, err := cloudtasks.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &taskspb.DeleteTaskRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTask(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END cloudtasks_generated_cloudtasks_apiv2beta3_Client_DeleteTask]
