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

// [START analyticsdata_generated_analytics_data_apiv1alpha_AlphaAnalyticsDataClient_GetMetadata]

package main

import (
	"context"

	data "cloud.google.com/go/analytics/data/apiv1alpha"
	datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"
)

func main() {
	// import datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"

	ctx := context.Background()
	c, err := data.NewAlphaAnalyticsDataClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datapb.GetMetadataRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetMetadata(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END analyticsdata_generated_analytics_data_apiv1alpha_AlphaAnalyticsDataClient_GetMetadata]
