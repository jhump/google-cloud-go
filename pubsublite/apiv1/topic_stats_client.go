// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package pubsublite

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	pubsublitepb "google.golang.org/genproto/googleapis/cloud/pubsublite/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newTopicStatsClientHook clientHook

// TopicStatsCallOptions contains the retry settings for each method of TopicStatsClient.
type TopicStatsCallOptions struct {
	ComputeMessageStats []gax.CallOption
	ComputeHeadCursor   []gax.CallOption
}

func defaultTopicStatsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("pubsublite.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("pubsublite.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://pubsublite.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTopicStatsCallOptions() *TopicStatsCallOptions {
	return &TopicStatsCallOptions{
		ComputeMessageStats: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.Aborted,
					codes.Internal,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ComputeHeadCursor: []gax.CallOption{},
	}
}

// TopicStatsClient is a client for interacting with Pub/Sub Lite API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type TopicStatsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	topicStatsClient pubsublitepb.TopicStatsServiceClient

	// The call options for this service.
	CallOptions *TopicStatsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTopicStatsClient creates a new topic stats service client.
//
// This service allows users to get stats about messages in their topic.
func NewTopicStatsClient(ctx context.Context, opts ...option.ClientOption) (*TopicStatsClient, error) {
	clientOpts := defaultTopicStatsClientOptions()

	if newTopicStatsClientHook != nil {
		hookOpts, err := newTopicStatsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &TopicStatsClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultTopicStatsCallOptions(),

		topicStatsClient: pubsublitepb.NewTopicStatsServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TopicStatsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TopicStatsClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TopicStatsClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ComputeMessageStats compute statistics about a range of messages in a given topic and
// partition.
func (c *TopicStatsClient) ComputeMessageStats(ctx context.Context, req *pubsublitepb.ComputeMessageStatsRequest, opts ...gax.CallOption) (*pubsublitepb.ComputeMessageStatsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "topic", url.QueryEscape(req.GetTopic())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ComputeMessageStats[0:len(c.CallOptions.ComputeMessageStats):len(c.CallOptions.ComputeMessageStats)], opts...)
	var resp *pubsublitepb.ComputeMessageStatsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.topicStatsClient.ComputeMessageStats(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ComputeHeadCursor compute the head cursor for the partition.
// The head cursor’s offset is guaranteed to be less than or equal to all
// messages which have not yet been acknowledged as published, and
// greater than the offset of any message whose publish has already
// been acknowledged. It is zero if there have never been messages in the
// partition.
func (c *TopicStatsClient) ComputeHeadCursor(ctx context.Context, req *pubsublitepb.ComputeHeadCursorRequest, opts ...gax.CallOption) (*pubsublitepb.ComputeHeadCursorResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "topic", url.QueryEscape(req.GetTopic())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ComputeHeadCursor[0:len(c.CallOptions.ComputeHeadCursor):len(c.CallOptions.ComputeHeadCursor)], opts...)
	var resp *pubsublitepb.ComputeHeadCursorResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.topicStatsClient.ComputeHeadCursor(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
