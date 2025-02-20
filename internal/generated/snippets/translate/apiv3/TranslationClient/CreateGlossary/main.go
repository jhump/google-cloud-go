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

// [START translate_generated_translate_apiv3_TranslationClient_CreateGlossary]

package main

import (
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func main() {
	// import translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"

	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &translatepb.CreateGlossaryRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateGlossary(ctx, req)
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

// [END translate_generated_translate_apiv3_TranslationClient_CreateGlossary]
