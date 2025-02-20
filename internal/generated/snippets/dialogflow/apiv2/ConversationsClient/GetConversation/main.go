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

// [START dialogflow_generated_dialogflow_apiv2_ConversationsClient_GetConversation]

package main

import (
	"context"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func main() {
	// import dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"

	ctx := context.Background()
	c, err := dialogflow.NewConversationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &dialogflowpb.GetConversationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetConversation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END dialogflow_generated_dialogflow_apiv2_ConversationsClient_GetConversation]
