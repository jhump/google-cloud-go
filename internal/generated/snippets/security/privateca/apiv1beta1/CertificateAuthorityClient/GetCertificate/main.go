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

// [START privateca_generated_security_privateca_apiv1beta1_CertificateAuthorityClient_GetCertificate]

package main

import (
	"context"

	privateca "cloud.google.com/go/security/privateca/apiv1beta1"
	privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1beta1"
)

func main() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1beta1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.GetCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END privateca_generated_security_privateca_apiv1beta1_CertificateAuthorityClient_GetCertificate]
