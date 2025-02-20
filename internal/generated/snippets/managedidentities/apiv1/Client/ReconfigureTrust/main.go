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

// [START managedidentities_generated_managedidentities_apiv1_Client_ReconfigureTrust]

package main

import (
	"context"

	managedidentities "cloud.google.com/go/managedidentities/apiv1"
	managedidentitiespb "google.golang.org/genproto/googleapis/cloud/managedidentities/v1"
)

func main() {
	// import managedidentitiespb "google.golang.org/genproto/googleapis/cloud/managedidentities/v1"

	ctx := context.Background()
	c, err := managedidentities.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &managedidentitiespb.ReconfigureTrustRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ReconfigureTrust(ctx, req)
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

// [END managedidentities_generated_managedidentities_apiv1_Client_ReconfigureTrust]
