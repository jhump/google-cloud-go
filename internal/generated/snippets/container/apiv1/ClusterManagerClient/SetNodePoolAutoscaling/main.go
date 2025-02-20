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

// [START container_generated_container_apiv1_ClusterManagerClient_SetNodePoolAutoscaling]

package main

import (
	"context"

	container "cloud.google.com/go/container/apiv1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func main() {
	// import containerpb "google.golang.org/genproto/googleapis/container/v1"

	ctx := context.Background()
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &containerpb.SetNodePoolAutoscalingRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetNodePoolAutoscaling(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END container_generated_container_apiv1_ClusterManagerClient_SetNodePoolAutoscaling]
