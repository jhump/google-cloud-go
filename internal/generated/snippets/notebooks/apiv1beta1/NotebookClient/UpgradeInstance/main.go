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

// [START notebooks_generated_notebooks_apiv1beta1_NotebookClient_UpgradeInstance]

package main

import (
	"context"

	notebooks "cloud.google.com/go/notebooks/apiv1beta1"
	notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"
)

func main() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.UpgradeInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpgradeInstance(ctx, req)
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

// [END notebooks_generated_notebooks_apiv1beta1_NotebookClient_UpgradeInstance]
