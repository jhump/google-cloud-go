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

// [START spanner_generated_spanner_admin_instance_apiv1_InstanceAdminClient_DeleteInstance]

package main

import (
	"context"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	instancepb "google.golang.org/genproto/googleapis/spanner/admin/instance/v1"
)

func main() {
	ctx := context.Background()
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &instancepb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END spanner_generated_spanner_admin_instance_apiv1_InstanceAdminClient_DeleteInstance]
