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

// [START orgpolicy_generated_orgpolicy_apiv2_Client_GetPolicy]

package main

import (
	"context"

	orgpolicy "cloud.google.com/go/orgpolicy/apiv2"
	orgpolicypb "google.golang.org/genproto/googleapis/cloud/orgpolicy/v2"
)

func main() {
	// import orgpolicypb "google.golang.org/genproto/googleapis/cloud/orgpolicy/v2"

	ctx := context.Background()
	c, err := orgpolicy.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &orgpolicypb.GetPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END orgpolicy_generated_orgpolicy_apiv2_Client_GetPolicy]
