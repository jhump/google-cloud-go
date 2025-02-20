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

// [START servicecontrol_generated_servicecontrol_apiv1_ServiceControllerClient_Check]

package main

import (
	"context"

	servicecontrol "cloud.google.com/go/servicecontrol/apiv1"
	servicecontrolpb "google.golang.org/genproto/googleapis/api/servicecontrol/v1"
)

func main() {
	// import servicecontrolpb "google.golang.org/genproto/googleapis/api/servicecontrol/v1"

	ctx := context.Background()
	c, err := servicecontrol.NewServiceControllerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicecontrolpb.CheckRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Check(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END servicecontrol_generated_servicecontrol_apiv1_ServiceControllerClient_Check]
