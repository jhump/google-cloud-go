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

// [START clouddebugger_generated_debugger_apiv2_Controller2Client_RegisterDebuggee]

package main

import (
	"context"

	debugger "cloud.google.com/go/debugger/apiv2"
	clouddebuggerpb "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2"
)

func main() {
	// import clouddebuggerpb "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2"

	ctx := context.Background()
	c, err := debugger.NewController2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &clouddebuggerpb.RegisterDebuggeeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RegisterDebuggee(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END clouddebugger_generated_debugger_apiv2_Controller2Client_RegisterDebuggee]
