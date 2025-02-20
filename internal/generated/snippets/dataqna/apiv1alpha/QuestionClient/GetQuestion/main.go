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

// [START dataqna_generated_dataqna_apiv1alpha_QuestionClient_GetQuestion]

package main

import (
	"context"

	dataqna "cloud.google.com/go/dataqna/apiv1alpha"
	dataqnapb "google.golang.org/genproto/googleapis/cloud/dataqna/v1alpha"
)

func main() {
	// import dataqnapb "google.golang.org/genproto/googleapis/cloud/dataqna/v1alpha"

	ctx := context.Background()
	c, err := dataqna.NewQuestionClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &dataqnapb.GetQuestionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetQuestion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END dataqna_generated_dataqna_apiv1alpha_QuestionClient_GetQuestion]
