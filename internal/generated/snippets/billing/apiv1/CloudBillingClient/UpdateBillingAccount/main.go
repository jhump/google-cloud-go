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

// [START cloudbilling_generated_billing_apiv1_CloudBillingClient_UpdateBillingAccount]

package main

import (
	"context"

	billing "cloud.google.com/go/billing/apiv1"
	billingpb "google.golang.org/genproto/googleapis/cloud/billing/v1"
)

func main() {
	// import billingpb "google.golang.org/genproto/googleapis/cloud/billing/v1"

	ctx := context.Background()
	c, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &billingpb.UpdateBillingAccountRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateBillingAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END cloudbilling_generated_billing_apiv1_CloudBillingClient_UpdateBillingAccount]
