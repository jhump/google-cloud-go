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

// [START containeranalysis_generated_containeranalysis_apiv1beta1_GrafeasV1Beta1Client_UpdateNote]

package main

import (
	"context"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
	grafeaspb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/grafeas"
)

func main() {
	// import grafeaspb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/grafeas"

	ctx := context.Background()
	c, err := containeranalysis.NewGrafeasV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &grafeaspb.UpdateNoteRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateNote(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END containeranalysis_generated_containeranalysis_apiv1beta1_GrafeasV1Beta1Client_UpdateNote]
