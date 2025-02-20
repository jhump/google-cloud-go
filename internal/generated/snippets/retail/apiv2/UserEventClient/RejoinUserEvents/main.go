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

// [START retail_generated_retail_apiv2_UserEventClient_RejoinUserEvents]

package main

import (
	"context"

	retail "cloud.google.com/go/retail/apiv2"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
)

func main() {
	// import retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"

	ctx := context.Background()
	c, err := retail.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &retailpb.RejoinUserEventsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RejoinUserEvents(ctx, req)
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

// [END retail_generated_retail_apiv2_UserEventClient_RejoinUserEvents]
