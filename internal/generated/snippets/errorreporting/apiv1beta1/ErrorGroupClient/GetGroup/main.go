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

// [START clouderrorreporting_generated_errorreporting_apiv1beta1_ErrorGroupClient_GetGroup]

package main

import (
	"context"

	errorreporting "cloud.google.com/go/errorreporting/apiv1beta1"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
)

func main() {
	// import clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"

	ctx := context.Background()
	c, err := errorreporting.NewErrorGroupClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &clouderrorreportingpb.GetGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END clouderrorreporting_generated_errorreporting_apiv1beta1_ErrorGroupClient_GetGroup]
