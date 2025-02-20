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

// [START osconfig_generated_osconfig_apiv1beta_Client_GetGuestPolicy]

package main

import (
	"context"

	osconfig "cloud.google.com/go/osconfig/apiv1beta"
	osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1beta"
)

func main() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1beta"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.GetGuestPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGuestPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END osconfig_generated_osconfig_apiv1beta_Client_GetGuestPolicy]
