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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudresourcemanager/alpha/cloudresourcemanager_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/alpha"
)

// TagValueServer implements the gRPC interface for TagValue.
type TagValueServer struct{}

// ProtoToTagValue converts a TagValue resource from its proto representation.
func ProtoToTagValue(p *alphapb.CloudresourcemanagerAlphaTagValue) *alpha.TagValue {
	obj := &alpha.TagValue{
		Name:           dcl.StringOrNil(p.GetName()),
		Parent:         dcl.StringOrNil(p.GetParent()),
		ShortName:      dcl.StringOrNil(p.GetShortName()),
		NamespacedName: dcl.StringOrNil(p.GetNamespacedName()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
	}
	return obj
}

// TagValueToProto converts a TagValue resource to its proto representation.
func TagValueToProto(resource *alpha.TagValue) *alphapb.CloudresourcemanagerAlphaTagValue {
	p := &alphapb.CloudresourcemanagerAlphaTagValue{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetShortName(dcl.ValueOrEmptyString(resource.ShortName))
	p.SetNamespacedName(dcl.ValueOrEmptyString(resource.NamespacedName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))

	return p
}

// applyTagValue handles the gRPC request by passing it to the underlying TagValue Apply() method.
func (s *TagValueServer) applyTagValue(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudresourcemanagerAlphaTagValueRequest) (*alphapb.CloudresourcemanagerAlphaTagValue, error) {
	p := ProtoToTagValue(request.GetResource())
	res, err := c.ApplyTagValue(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TagValueToProto(res)
	return r, nil
}

// applyCloudresourcemanagerAlphaTagValue handles the gRPC request by passing it to the underlying TagValue Apply() method.
func (s *TagValueServer) ApplyCloudresourcemanagerAlphaTagValue(ctx context.Context, request *alphapb.ApplyCloudresourcemanagerAlphaTagValueRequest) (*alphapb.CloudresourcemanagerAlphaTagValue, error) {
	cl, err := createConfigTagValue(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTagValue(ctx, cl, request)
}

// DeleteTagValue handles the gRPC request by passing it to the underlying TagValue Delete() method.
func (s *TagValueServer) DeleteCloudresourcemanagerAlphaTagValue(ctx context.Context, request *alphapb.DeleteCloudresourcemanagerAlphaTagValueRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTagValue(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTagValue(ctx, ProtoToTagValue(request.GetResource()))

}

// ListCloudresourcemanagerAlphaTagValue is a no-op method because TagValue has no list method.
func (s *TagValueServer) ListCloudresourcemanagerAlphaTagValue(_ context.Context, _ *alphapb.ListCloudresourcemanagerAlphaTagValueRequest) (*alphapb.ListCloudresourcemanagerAlphaTagValueResponse, error) {
	return nil, nil
}

func createConfigTagValue(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
