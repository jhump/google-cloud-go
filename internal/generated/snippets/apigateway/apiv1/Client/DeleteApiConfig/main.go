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

// [START apigateway_generated_apigateway_apiv1_Client_DeleteApiConfig]

package main

import (
	"context"

	apigateway "cloud.google.com/go/apigateway/apiv1"
	apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
)

func main() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.DeleteApiConfigRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteApiConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END apigateway_generated_apigateway_apiv1_Client_DeleteApiConfig]
