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

// [START dialogflow_generated_dialogflow_cx_apiv3beta1_SessionsClient_StreamingDetectIntent]

package main

import (
	"context"
	"io"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
)

func main() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewSessionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	stream, err := c.StreamingDetectIntent(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	go func() {
		reqs := []*cxpb.StreamingDetectIntentRequest{
			// TODO: Create requests.
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				// TODO: Handle error.
			}
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

// [END dialogflow_generated_dialogflow_cx_apiv3beta1_SessionsClient_StreamingDetectIntent]
