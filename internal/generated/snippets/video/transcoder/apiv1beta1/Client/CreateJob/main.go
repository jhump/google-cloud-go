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

// [START transcoder_generated_video_transcoder_apiv1beta1_Client_CreateJob]

package main

import (
	"context"

	transcoder "cloud.google.com/go/video/transcoder/apiv1beta1"
	transcoderpb "google.golang.org/genproto/googleapis/cloud/video/transcoder/v1beta1"
)

func main() {
	// import transcoderpb "google.golang.org/genproto/googleapis/cloud/video/transcoder/v1beta1"

	ctx := context.Background()
	c, err := transcoder.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &transcoderpb.CreateJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END transcoder_generated_video_transcoder_apiv1beta1_Client_CreateJob]
