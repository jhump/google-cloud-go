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

// [START longrunning_generated_longrunning_autogen_OperationsClient_GetOperation]

package main

import (
	"context"

	longrunning "cloud.google.com/go/longrunning/autogen"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func main() {
	// import longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	ctx := context.Background()
	c, err := longrunning.NewOperationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END longrunning_generated_longrunning_autogen_OperationsClient_GetOperation]
