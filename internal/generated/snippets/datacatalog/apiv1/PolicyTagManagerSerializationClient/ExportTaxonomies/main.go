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

// [START datacatalog_generated_datacatalog_apiv1_PolicyTagManagerSerializationClient_ExportTaxonomies]

package main

import (
	"context"

	datacatalog "cloud.google.com/go/datacatalog/apiv1"
	datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
)

func main() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewPolicyTagManagerSerializationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.ExportTaxonomiesRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ExportTaxonomies(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END datacatalog_generated_datacatalog_apiv1_PolicyTagManagerSerializationClient_ExportTaxonomies]
