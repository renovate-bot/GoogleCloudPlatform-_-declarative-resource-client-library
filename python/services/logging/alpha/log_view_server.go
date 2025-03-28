// Copyright 2025 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/alpha/logging_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/alpha"
)

// LogViewServer implements the gRPC interface for LogView.
type LogViewServer struct{}

// ProtoToLogView converts a LogView resource from its proto representation.
func ProtoToLogView(p *alphapb.LoggingAlphaLogView) *alpha.LogView {
	obj := &alpha.LogView{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Parent:      dcl.StringOrNil(p.GetParent()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		Bucket:      dcl.StringOrNil(p.GetBucket()),
	}
	return obj
}

// LogViewToProto converts a LogView resource to its proto representation.
func LogViewToProto(resource *alpha.LogView) *alphapb.LoggingAlphaLogView {
	p := &alphapb.LoggingAlphaLogView{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetFilter(dcl.ValueOrEmptyString(resource.Filter))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetBucket(dcl.ValueOrEmptyString(resource.Bucket))

	return p
}

// applyLogView handles the gRPC request by passing it to the underlying LogView Apply() method.
func (s *LogViewServer) applyLogView(ctx context.Context, c *alpha.Client, request *alphapb.ApplyLoggingAlphaLogViewRequest) (*alphapb.LoggingAlphaLogView, error) {
	p := ProtoToLogView(request.GetResource())
	res, err := c.ApplyLogView(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogViewToProto(res)
	return r, nil
}

// applyLoggingAlphaLogView handles the gRPC request by passing it to the underlying LogView Apply() method.
func (s *LogViewServer) ApplyLoggingAlphaLogView(ctx context.Context, request *alphapb.ApplyLoggingAlphaLogViewRequest) (*alphapb.LoggingAlphaLogView, error) {
	cl, err := createConfigLogView(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLogView(ctx, cl, request)
}

// DeleteLogView handles the gRPC request by passing it to the underlying LogView Delete() method.
func (s *LogViewServer) DeleteLoggingAlphaLogView(ctx context.Context, request *alphapb.DeleteLoggingAlphaLogViewRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogView(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogView(ctx, ProtoToLogView(request.GetResource()))

}

// ListLoggingAlphaLogView handles the gRPC request by passing it to the underlying LogViewList() method.
func (s *LogViewServer) ListLoggingAlphaLogView(ctx context.Context, request *alphapb.ListLoggingAlphaLogViewRequest) (*alphapb.ListLoggingAlphaLogViewResponse, error) {
	cl, err := createConfigLogView(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogView(ctx, request.GetLocation(), request.GetBucket(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.LoggingAlphaLogView
	for _, r := range resources.Items {
		rp := LogViewToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListLoggingAlphaLogViewResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLogView(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
