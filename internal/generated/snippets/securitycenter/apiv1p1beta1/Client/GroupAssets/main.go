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

// [START securitycenter_generated_securitycenter_apiv1p1beta1_Client_GroupAssets]

package main

import (
	"context"

	securitycenter "cloud.google.com/go/securitycenter/apiv1p1beta1"
	"google.golang.org/api/iterator"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
)

func main() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.GroupAssetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.GroupAssets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

// [END securitycenter_generated_securitycenter_apiv1p1beta1_Client_GroupAssets]
