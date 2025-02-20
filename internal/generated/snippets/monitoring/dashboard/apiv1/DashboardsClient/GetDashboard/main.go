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

// [START monitoring_generated_monitoring_dashboard_apiv1_DashboardsClient_GetDashboard]

package main

import (
	"context"

	dashboard "cloud.google.com/go/monitoring/dashboard/apiv1"
	dashboardpb "google.golang.org/genproto/googleapis/monitoring/dashboard/v1"
)

func main() {
	// import dashboardpb "google.golang.org/genproto/googleapis/monitoring/dashboard/v1"

	ctx := context.Background()
	c, err := dashboard.NewDashboardsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &dashboardpb.GetDashboardRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDashboard(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END monitoring_generated_monitoring_dashboard_apiv1_DashboardsClient_GetDashboard]
