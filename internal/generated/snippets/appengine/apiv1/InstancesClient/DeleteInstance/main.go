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

// [START appengine_generated_appengine_apiv1_InstancesClient_DeleteInstance]

package main

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func main() {
	// import appenginepb "google.golang.org/genproto/googleapis/appengine/v1"

	ctx := context.Background()
	c, err := appengine.NewInstancesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &appenginepb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END appengine_generated_appengine_apiv1_InstancesClient_DeleteInstance]
