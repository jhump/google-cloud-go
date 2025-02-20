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

// [START language_generated_language_apiv1_Client_AnalyzeSyntax]

package main

import (
	"context"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	// import languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"

	ctx := context.Background()
	c, err := language.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &languagepb.AnalyzeSyntaxRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AnalyzeSyntax(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END language_generated_language_apiv1_Client_AnalyzeSyntax]
