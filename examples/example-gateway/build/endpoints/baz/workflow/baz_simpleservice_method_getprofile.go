// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"context"
	"net/textproto"

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsIDlClientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/baz/baz"
	endpointsIDlEndpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints-idl/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceGetProfileWorkflow defines the interface for SimpleServiceGetProfile workflow
type SimpleServiceGetProfileWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsIDlEndpointsBazBaz.SimpleService_GetProfile_Args,
	) (context.Context, *endpointsIDlEndpointsBazBaz.GetProfileResponse, zanzibar.Header, error)
}

// NewSimpleServiceGetProfileWorkflow creates a workflow
func NewSimpleServiceGetProfileWorkflow(deps *module.Dependencies) SimpleServiceGetProfileWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.baz.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
		}
	}

	return &simpleServiceGetProfileWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
		defaultDeps:               deps.Default,
	}
}

// simpleServiceGetProfileWorkflow calls thrift client Baz.GetProfile
type simpleServiceGetProfileWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
	defaultDeps               *zanzibar.DefaultDependencies
}

// Handle calls thrift client.
func (w simpleServiceGetProfileWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsIDlEndpointsBazBaz.SimpleService_GetProfile_Args,
) (context.Context, *endpointsIDlEndpointsBazBaz.GetProfileResponse, zanzibar.Header, error) {
	clientRequest := convertToGetProfileClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	var k string

	k = textproto.CanonicalMIMEHeaderKey("x-uber-foo")
	h, ok = reqHeaders.Get(k)
	if ok {
		clientHeaders[k] = h
	}
	k = textproto.CanonicalMIMEHeaderKey("x-uber-bar")
	h, ok = reqHeaders.Get(k)
	if ok {
		clientHeaders[k] = h
	}

	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	for _, whitelistedHeader := range w.whitelistedDynamicHeaders {
		headerVal, ok := reqHeaders.Get(whitelistedHeader)
		if ok {
			clientHeaders[whitelistedHeader] = headerVal
		}
	}

	//when maxRetry is 0, timeout per client level is used & one attempt is made, and timoutPerAttempt is not used
	var timeoutAndRetryConfig = zanzibar.TimeoutAndRetryOptions{}

	//when endpoint level timeout information is available, override it with client level config
	if w.defaultDeps.Config.ContainsKey("endpoints.baz.getProfile.timeoutPerAttempt") {
		scaleFactor := w.defaultDeps.Config.MustGetFloat("endpoints.baz.getProfile.scaleFactor")
		maxRetry := int(w.defaultDeps.Config.MustGetInt("endpoints.baz.getProfile.retryCount"))

		backOffTimeAcrossRetriesCfg := int(w.defaultDeps.Config.MustGetInt("endpoints.baz.getProfile.backOffTimeAcrossRetries"))
		timeoutPerAttemptConf := int(w.defaultDeps.Config.MustGetInt("endpoints.baz.getProfile.timeoutPerAttempt"))

		timeoutAndRetryConfig = zanzibar.BuildTimeoutAndRetryConfig(int(timeoutPerAttemptConf), backOffTimeAcrossRetriesCfg, maxRetry, scaleFactor)
	}

	ctx, clientRespBody, cliRespHeaders, err := w.Clients.Baz.GetProfile(
		ctx, clientHeaders, clientRequest, &timeoutAndRetryConfig,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsIDlClientsBazBaz.AuthErr:
			serverErr := convertGetProfileAuthErr(
				errValue,
			)

			return ctx, nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return ctx, nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceGetProfileClientResponse(clientRespBody)
	if val, ok := cliRespHeaders[zanzibar.ClientResponseDurationKey]; ok {
		resHeaders.Set(zanzibar.ClientResponseDurationKey, val)
	}

	resHeaders.Set(zanzibar.ClientTypeKey, "tchannel")
	return ctx, response, resHeaders, nil
}

func convertToGetProfileClientRequest(in *endpointsIDlEndpointsBazBaz.SimpleService_GetProfile_Args) *clientsIDlClientsBazBaz.SimpleService_GetProfile_Args {
	out := &clientsIDlClientsBazBaz.SimpleService_GetProfile_Args{}

	if in.Request != nil {
		out.Request = &clientsIDlClientsBazBaz.GetProfileRequest{}
		out.Request.Target = clientsIDlClientsBazBaz.UUID(in.Request.Target)
	} else {
		out.Request = nil
	}

	return out
}

func convertGetProfileAuthErr(
	clientError *clientsIDlClientsBazBaz.AuthErr,
) *endpointsIDlEndpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsIDlEndpointsBazBaz.AuthErr{}
	return serverError
}

func convertSimpleServiceGetProfileClientResponse(in *clientsIDlClientsBazBaz.GetProfileResponse) *endpointsIDlEndpointsBazBaz.GetProfileResponse {
	out := &endpointsIDlEndpointsBazBaz.GetProfileResponse{}

	out.Payloads = make([]*endpointsIDlEndpointsBazBaz.Profile, len(in.Payloads))
	for index1, value2 := range in.Payloads {
		if value2 != nil {
			out.Payloads[index1] = &endpointsIDlEndpointsBazBaz.Profile{}
			if in.Payloads[index1].Recur1 != nil {
				out.Payloads[index1].Recur1 = &endpointsIDlEndpointsBazBaz.Recur1{}
				out.Payloads[index1].Recur1.Field1 = make(map[endpointsIDlEndpointsBazBaz.UUID]*endpointsIDlEndpointsBazBaz.Recur2, len(in.Payloads[index1].Recur1.Field1))
				for key3, value4 := range in.Payloads[index1].Recur1.Field1 {
					if value4 != nil {
						out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)] = &endpointsIDlEndpointsBazBaz.Recur2{}
						if in.Payloads[index1].Recur1.Field1[key3].Field21 != nil {
							out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)].Field21 = &endpointsIDlEndpointsBazBaz.Recur3{}
							if in.Payloads[index1] != nil && in.Payloads[index1].Recur1 != nil && in.Payloads[index1].Recur1.Field1[key3] != nil && in.Payloads[index1].Recur1.Field1[key3].Field21 != nil {
								out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)].Field21.Field31 = endpointsIDlEndpointsBazBaz.UUID(in.Payloads[index1].Recur1.Field1[key3].Field21.Field31)
							}
						} else {
							out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)].Field21 = nil
						}
						if in.Payloads[index1].Recur1.Field1[key3].Field22 != nil {
							out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)].Field22 = &endpointsIDlEndpointsBazBaz.Recur3{}
							if in.Payloads[index1] != nil && in.Payloads[index1].Recur1 != nil && in.Payloads[index1].Recur1.Field1[key3] != nil && in.Payloads[index1].Recur1.Field1[key3].Field22 != nil {
								out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)].Field22.Field31 = endpointsIDlEndpointsBazBaz.UUID(in.Payloads[index1].Recur1.Field1[key3].Field22.Field31)
							}
						} else {
							out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)].Field22 = nil
						}
					} else {
						out.Payloads[index1].Recur1.Field1[endpointsIDlEndpointsBazBaz.UUID(key3)] = nil
					}
				}
			} else {
				out.Payloads[index1].Recur1 = nil
			}
		} else {
			out.Payloads[index1] = nil
		}
	}

	return out
}