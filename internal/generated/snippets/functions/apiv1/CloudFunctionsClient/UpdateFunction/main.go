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

// [START cloudfunctions_generated_functions_apiv1_CloudFunctionsClient_UpdateFunction]

package main

import (
	"context"

	functions "cloud.google.com/go/functions/apiv1"
	functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"
)

func main() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.UpdateFunctionRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateFunction(ctx, req)
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

// [END cloudfunctions_generated_functions_apiv1_CloudFunctionsClient_UpdateFunction]
