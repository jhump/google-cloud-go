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

// [START iamcredentials_generated_iam_credentials_apiv1_IamCredentialsClient_SignBlob]

package main

import (
	"context"

	credentials "cloud.google.com/go/iam/credentials/apiv1"
	credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

func main() {
	// import credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"

	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &credentialspb.SignBlobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SignBlob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END iamcredentials_generated_iam_credentials_apiv1_IamCredentialsClient_SignBlob]
