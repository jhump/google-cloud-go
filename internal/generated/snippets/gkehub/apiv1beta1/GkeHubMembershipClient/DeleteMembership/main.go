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

// [START gkehub_generated_gkehub_apiv1beta1_GkeHubMembershipClient_DeleteMembership]

package main

import (
	"context"

	gkehub "cloud.google.com/go/gkehub/apiv1beta1"
	gkehubpb "google.golang.org/genproto/googleapis/cloud/gkehub/v1beta1"
)

func main() {
	// import gkehubpb "google.golang.org/genproto/googleapis/cloud/gkehub/v1beta1"

	ctx := context.Background()
	c, err := gkehub.NewGkeHubMembershipClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gkehubpb.DeleteMembershipRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteMembership(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END gkehub_generated_gkehub_apiv1beta1_GkeHubMembershipClient_DeleteMembership]
