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

// [START logging_generated_logging_ContextFunc]

package main

import (
	"context"

	"cloud.google.com/go/logging"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, "my-project")
	if err != nil {
		// TODO: Handle error.
	}
	lg := client.Logger("logID", logging.ContextFunc(func() (context.Context, func()) {
		ctx, span := trace.StartSpan(context.Background(), "this span will not be exported",
			trace.WithSampler(trace.NeverSample()))
		return ctx, span.End
	}))
	_ = lg // TODO: Use lg
}

// [END logging_generated_logging_ContextFunc]
