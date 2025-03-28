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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/pubsub/beta/pubsub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/beta"
)

// TopicServer implements the gRPC interface for Topic.
type TopicServer struct{}

// ProtoToTopicMessageStoragePolicy converts a TopicMessageStoragePolicy object from its proto representation.
func ProtoToPubsubBetaTopicMessageStoragePolicy(p *betapb.PubsubBetaTopicMessageStoragePolicy) *beta.TopicMessageStoragePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.TopicMessageStoragePolicy{}
	for _, r := range p.GetAllowedPersistenceRegions() {
		obj.AllowedPersistenceRegions = append(obj.AllowedPersistenceRegions, r)
	}
	return obj
}

// ProtoToTopic converts a Topic resource from its proto representation.
func ProtoToTopic(p *betapb.PubsubBetaTopic) *beta.Topic {
	obj := &beta.Topic{
		Name:                 dcl.StringOrNil(p.GetName()),
		KmsKeyName:           dcl.StringOrNil(p.GetKmsKeyName()),
		MessageStoragePolicy: ProtoToPubsubBetaTopicMessageStoragePolicy(p.GetMessageStoragePolicy()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// TopicMessageStoragePolicyToProto converts a TopicMessageStoragePolicy object to its proto representation.
func PubsubBetaTopicMessageStoragePolicyToProto(o *beta.TopicMessageStoragePolicy) *betapb.PubsubBetaTopicMessageStoragePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.PubsubBetaTopicMessageStoragePolicy{}
	sAllowedPersistenceRegions := make([]string, len(o.AllowedPersistenceRegions))
	for i, r := range o.AllowedPersistenceRegions {
		sAllowedPersistenceRegions[i] = r
	}
	p.SetAllowedPersistenceRegions(sAllowedPersistenceRegions)
	return p
}

// TopicToProto converts a Topic resource to its proto representation.
func TopicToProto(resource *beta.Topic) *betapb.PubsubBetaTopic {
	p := &betapb.PubsubBetaTopic{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetKmsKeyName(dcl.ValueOrEmptyString(resource.KmsKeyName))
	p.SetMessageStoragePolicy(PubsubBetaTopicMessageStoragePolicyToProto(resource.MessageStoragePolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) applyTopic(ctx context.Context, c *beta.Client, request *betapb.ApplyPubsubBetaTopicRequest) (*betapb.PubsubBetaTopic, error) {
	p := ProtoToTopic(request.GetResource())
	res, err := c.ApplyTopic(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TopicToProto(res)
	return r, nil
}

// applyPubsubBetaTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) ApplyPubsubBetaTopic(ctx context.Context, request *betapb.ApplyPubsubBetaTopicRequest) (*betapb.PubsubBetaTopic, error) {
	cl, err := createConfigTopic(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTopic(ctx, cl, request)
}

// DeleteTopic handles the gRPC request by passing it to the underlying Topic Delete() method.
func (s *TopicServer) DeletePubsubBetaTopic(ctx context.Context, request *betapb.DeletePubsubBetaTopicRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTopic(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTopic(ctx, ProtoToTopic(request.GetResource()))

}

// ListPubsubBetaTopic handles the gRPC request by passing it to the underlying TopicList() method.
func (s *TopicServer) ListPubsubBetaTopic(ctx context.Context, request *betapb.ListPubsubBetaTopicRequest) (*betapb.ListPubsubBetaTopicResponse, error) {
	cl, err := createConfigTopic(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTopic(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.PubsubBetaTopic
	for _, r := range resources.Items {
		rp := TopicToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListPubsubBetaTopicResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTopic(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
