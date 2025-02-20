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

// [START redis_generated_redis_apiv1beta1_CloudRedisClient_ExportInstance]

package main

import (
	"context"

	redis "cloud.google.com/go/redis/apiv1beta1"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1beta1"
)

func main() {
	// import redispb "google.golang.org/genproto/googleapis/cloud/redis/v1beta1"

	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &redispb.ExportInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportInstance(ctx, req)
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

// [END redis_generated_redis_apiv1beta1_CloudRedisClient_ExportInstance]
