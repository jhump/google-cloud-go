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

// [START cloudchannel_generated_channel_apiv1_CloudChannelClient_DeleteCustomer]

package main

import (
	"context"

	channel "cloud.google.com/go/channel/apiv1"
	channelpb "google.golang.org/genproto/googleapis/cloud/channel/v1"
)

func main() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &channelpb.DeleteCustomerRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteCustomer(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END cloudchannel_generated_channel_apiv1_CloudChannelClient_DeleteCustomer]
