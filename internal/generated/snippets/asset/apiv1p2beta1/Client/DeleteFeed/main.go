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

// [START cloudasset_generated_asset_apiv1p2beta1_Client_DeleteFeed]

package main

import (
	"context"

	asset "cloud.google.com/go/asset/apiv1p2beta1"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1p2beta1"
)

func main() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.DeleteFeedRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END cloudasset_generated_asset_apiv1p2beta1_Client_DeleteFeed]
