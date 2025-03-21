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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/alpha/clouddeploy_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/alpha"
)

// TargetServer implements the gRPC interface for Target.
type TargetServer struct{}

// ProtoToTargetExecutionConfigsUsagesEnum converts a TargetExecutionConfigsUsagesEnum enum from its proto representation.
func ProtoToClouddeployAlphaTargetExecutionConfigsUsagesEnum(e alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum) *alpha.TargetExecutionConfigsUsagesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum_name[int32(e)]; ok {
		e := alpha.TargetExecutionConfigsUsagesEnum(n[len("ClouddeployAlphaTargetExecutionConfigsUsagesEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetGke converts a TargetGke object from its proto representation.
func ProtoToClouddeployAlphaTargetGke(p *alphapb.ClouddeployAlphaTargetGke) *alpha.TargetGke {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetGke{
		Cluster:     dcl.StringOrNil(p.GetCluster()),
		InternalIP:  dcl.Bool(p.GetInternalIp()),
		ProxyUrl:    dcl.StringOrNil(p.GetProxyUrl()),
		DnsEndpoint: dcl.Bool(p.GetDnsEndpoint()),
	}
	return obj
}

// ProtoToTargetAnthosCluster converts a TargetAnthosCluster object from its proto representation.
func ProtoToClouddeployAlphaTargetAnthosCluster(p *alphapb.ClouddeployAlphaTargetAnthosCluster) *alpha.TargetAnthosCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetAnthosCluster{
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToTargetExecutionConfigs converts a TargetExecutionConfigs object from its proto representation.
func ProtoToClouddeployAlphaTargetExecutionConfigs(p *alphapb.ClouddeployAlphaTargetExecutionConfigs) *alpha.TargetExecutionConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetExecutionConfigs{
		WorkerPool:       dcl.StringOrNil(p.GetWorkerPool()),
		ServiceAccount:   dcl.StringOrNil(p.GetServiceAccount()),
		ArtifactStorage:  dcl.StringOrNil(p.GetArtifactStorage()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
		Verbose:          dcl.Bool(p.GetVerbose()),
	}
	for _, r := range p.GetUsages() {
		obj.Usages = append(obj.Usages, *ProtoToClouddeployAlphaTargetExecutionConfigsUsagesEnum(r))
	}
	return obj
}

// ProtoToTargetRun converts a TargetRun object from its proto representation.
func ProtoToClouddeployAlphaTargetRun(p *alphapb.ClouddeployAlphaTargetRun) *alpha.TargetRun {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetRun{
		Location: dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToTargetMultiTarget converts a TargetMultiTarget object from its proto representation.
func ProtoToClouddeployAlphaTargetMultiTarget(p *alphapb.ClouddeployAlphaTargetMultiTarget) *alpha.TargetMultiTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetMultiTarget{}
	for _, r := range p.GetTargetIds() {
		obj.TargetIds = append(obj.TargetIds, r)
	}
	return obj
}

// ProtoToTargetCustomTarget converts a TargetCustomTarget object from its proto representation.
func ProtoToClouddeployAlphaTargetCustomTarget(p *alphapb.ClouddeployAlphaTargetCustomTarget) *alpha.TargetCustomTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetCustomTarget{
		CustomTargetType: dcl.StringOrNil(p.GetCustomTargetType()),
	}
	return obj
}

// ProtoToTargetAssociatedEntities converts a TargetAssociatedEntities object from its proto representation.
func ProtoToClouddeployAlphaTargetAssociatedEntities(p *alphapb.ClouddeployAlphaTargetAssociatedEntities) *alpha.TargetAssociatedEntities {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetAssociatedEntities{}
	for _, r := range p.GetGkeClusters() {
		obj.GkeClusters = append(obj.GkeClusters, *ProtoToClouddeployAlphaTargetAssociatedEntitiesGkeClusters(r))
	}
	for _, r := range p.GetAnthosClusters() {
		obj.AnthosClusters = append(obj.AnthosClusters, *ProtoToClouddeployAlphaTargetAssociatedEntitiesAnthosClusters(r))
	}
	return obj
}

// ProtoToTargetAssociatedEntitiesGkeClusters converts a TargetAssociatedEntitiesGkeClusters object from its proto representation.
func ProtoToClouddeployAlphaTargetAssociatedEntitiesGkeClusters(p *alphapb.ClouddeployAlphaTargetAssociatedEntitiesGkeClusters) *alpha.TargetAssociatedEntitiesGkeClusters {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetAssociatedEntitiesGkeClusters{
		Cluster:    dcl.StringOrNil(p.GetCluster()),
		InternalIP: dcl.Bool(p.GetInternalIp()),
		ProxyUrl:   dcl.StringOrNil(p.GetProxyUrl()),
	}
	return obj
}

// ProtoToTargetAssociatedEntitiesAnthosClusters converts a TargetAssociatedEntitiesAnthosClusters object from its proto representation.
func ProtoToClouddeployAlphaTargetAssociatedEntitiesAnthosClusters(p *alphapb.ClouddeployAlphaTargetAssociatedEntitiesAnthosClusters) *alpha.TargetAssociatedEntitiesAnthosClusters {
	if p == nil {
		return nil
	}
	obj := &alpha.TargetAssociatedEntitiesAnthosClusters{
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToTarget converts a Target resource from its proto representation.
func ProtoToTarget(p *alphapb.ClouddeployAlphaTarget) *alpha.Target {
	obj := &alpha.Target{
		Name:            dcl.StringOrNil(p.GetName()),
		TargetId:        dcl.StringOrNil(p.GetTargetId()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		RequireApproval: dcl.Bool(p.GetRequireApproval()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Gke:             ProtoToClouddeployAlphaTargetGke(p.GetGke()),
		AnthosCluster:   ProtoToClouddeployAlphaTargetAnthosCluster(p.GetAnthosCluster()),
		Etag:            dcl.StringOrNil(p.GetEtag()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		Run:             ProtoToClouddeployAlphaTargetRun(p.GetRun()),
		MultiTarget:     ProtoToClouddeployAlphaTargetMultiTarget(p.GetMultiTarget()),
		CustomTarget:    ProtoToClouddeployAlphaTargetCustomTarget(p.GetCustomTarget()),
	}
	for _, r := range p.GetExecutionConfigs() {
		obj.ExecutionConfigs = append(obj.ExecutionConfigs, *ProtoToClouddeployAlphaTargetExecutionConfigs(r))
	}
	return obj
}

// TargetExecutionConfigsUsagesEnumToProto converts a TargetExecutionConfigsUsagesEnum enum to its proto representation.
func ClouddeployAlphaTargetExecutionConfigsUsagesEnumToProto(e *alpha.TargetExecutionConfigsUsagesEnum) alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum {
	if e == nil {
		return alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum(0)
	}
	if v, ok := alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum_value["TargetExecutionConfigsUsagesEnum"+string(*e)]; ok {
		return alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum(v)
	}
	return alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum(0)
}

// TargetGkeToProto converts a TargetGke object to its proto representation.
func ClouddeployAlphaTargetGkeToProto(o *alpha.TargetGke) *alphapb.ClouddeployAlphaTargetGke {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetGke{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetInternalIp(dcl.ValueOrEmptyBool(o.InternalIP))
	p.SetProxyUrl(dcl.ValueOrEmptyString(o.ProxyUrl))
	p.SetDnsEndpoint(dcl.ValueOrEmptyBool(o.DnsEndpoint))
	return p
}

// TargetAnthosClusterToProto converts a TargetAnthosCluster object to its proto representation.
func ClouddeployAlphaTargetAnthosClusterToProto(o *alpha.TargetAnthosCluster) *alphapb.ClouddeployAlphaTargetAnthosCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetAnthosCluster{}
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// TargetExecutionConfigsToProto converts a TargetExecutionConfigs object to its proto representation.
func ClouddeployAlphaTargetExecutionConfigsToProto(o *alpha.TargetExecutionConfigs) *alphapb.ClouddeployAlphaTargetExecutionConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetExecutionConfigs{}
	p.SetWorkerPool(dcl.ValueOrEmptyString(o.WorkerPool))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetArtifactStorage(dcl.ValueOrEmptyString(o.ArtifactStorage))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	p.SetVerbose(dcl.ValueOrEmptyBool(o.Verbose))
	sUsages := make([]alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum, len(o.Usages))
	for i, r := range o.Usages {
		sUsages[i] = alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum(alphapb.ClouddeployAlphaTargetExecutionConfigsUsagesEnum_value[string(r)])
	}
	p.SetUsages(sUsages)
	return p
}

// TargetRunToProto converts a TargetRun object to its proto representation.
func ClouddeployAlphaTargetRunToProto(o *alpha.TargetRun) *alphapb.ClouddeployAlphaTargetRun {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetRun{}
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// TargetMultiTargetToProto converts a TargetMultiTarget object to its proto representation.
func ClouddeployAlphaTargetMultiTargetToProto(o *alpha.TargetMultiTarget) *alphapb.ClouddeployAlphaTargetMultiTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetMultiTarget{}
	sTargetIds := make([]string, len(o.TargetIds))
	for i, r := range o.TargetIds {
		sTargetIds[i] = r
	}
	p.SetTargetIds(sTargetIds)
	return p
}

// TargetCustomTargetToProto converts a TargetCustomTarget object to its proto representation.
func ClouddeployAlphaTargetCustomTargetToProto(o *alpha.TargetCustomTarget) *alphapb.ClouddeployAlphaTargetCustomTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetCustomTarget{}
	p.SetCustomTargetType(dcl.ValueOrEmptyString(o.CustomTargetType))
	return p
}

// TargetAssociatedEntitiesToProto converts a TargetAssociatedEntities object to its proto representation.
func ClouddeployAlphaTargetAssociatedEntitiesToProto(o *alpha.TargetAssociatedEntities) *alphapb.ClouddeployAlphaTargetAssociatedEntities {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetAssociatedEntities{}
	sGkeClusters := make([]*alphapb.ClouddeployAlphaTargetAssociatedEntitiesGkeClusters, len(o.GkeClusters))
	for i, r := range o.GkeClusters {
		sGkeClusters[i] = ClouddeployAlphaTargetAssociatedEntitiesGkeClustersToProto(&r)
	}
	p.SetGkeClusters(sGkeClusters)
	sAnthosClusters := make([]*alphapb.ClouddeployAlphaTargetAssociatedEntitiesAnthosClusters, len(o.AnthosClusters))
	for i, r := range o.AnthosClusters {
		sAnthosClusters[i] = ClouddeployAlphaTargetAssociatedEntitiesAnthosClustersToProto(&r)
	}
	p.SetAnthosClusters(sAnthosClusters)
	return p
}

// TargetAssociatedEntitiesGkeClustersToProto converts a TargetAssociatedEntitiesGkeClusters object to its proto representation.
func ClouddeployAlphaTargetAssociatedEntitiesGkeClustersToProto(o *alpha.TargetAssociatedEntitiesGkeClusters) *alphapb.ClouddeployAlphaTargetAssociatedEntitiesGkeClusters {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetAssociatedEntitiesGkeClusters{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetInternalIp(dcl.ValueOrEmptyBool(o.InternalIP))
	p.SetProxyUrl(dcl.ValueOrEmptyString(o.ProxyUrl))
	return p
}

// TargetAssociatedEntitiesAnthosClustersToProto converts a TargetAssociatedEntitiesAnthosClusters object to its proto representation.
func ClouddeployAlphaTargetAssociatedEntitiesAnthosClustersToProto(o *alpha.TargetAssociatedEntitiesAnthosClusters) *alphapb.ClouddeployAlphaTargetAssociatedEntitiesAnthosClusters {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaTargetAssociatedEntitiesAnthosClusters{}
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// TargetToProto converts a Target resource to its proto representation.
func TargetToProto(resource *alpha.Target) *alphapb.ClouddeployAlphaTarget {
	p := &alphapb.ClouddeployAlphaTarget{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTargetId(dcl.ValueOrEmptyString(resource.TargetId))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetRequireApproval(dcl.ValueOrEmptyBool(resource.RequireApproval))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGke(ClouddeployAlphaTargetGkeToProto(resource.Gke))
	p.SetAnthosCluster(ClouddeployAlphaTargetAnthosClusterToProto(resource.AnthosCluster))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetRun(ClouddeployAlphaTargetRunToProto(resource.Run))
	p.SetMultiTarget(ClouddeployAlphaTargetMultiTargetToProto(resource.MultiTarget))
	p.SetCustomTarget(ClouddeployAlphaTargetCustomTargetToProto(resource.CustomTarget))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sExecutionConfigs := make([]*alphapb.ClouddeployAlphaTargetExecutionConfigs, len(resource.ExecutionConfigs))
	for i, r := range resource.ExecutionConfigs {
		sExecutionConfigs[i] = ClouddeployAlphaTargetExecutionConfigsToProto(&r)
	}
	p.SetExecutionConfigs(sExecutionConfigs)
	mDeployParameters := make(map[string]string, len(resource.DeployParameters))
	for k, r := range resource.DeployParameters {
		mDeployParameters[k] = r
	}
	p.SetDeployParameters(mDeployParameters)
	mAssociatedEntities := make(map[string]*alphapb.ClouddeployAlphaTargetAssociatedEntities, len(resource.AssociatedEntities))
	for k, r := range resource.AssociatedEntities {
		mAssociatedEntities[k] = ClouddeployAlphaTargetAssociatedEntitiesToProto(&r)
	}
	p.SetAssociatedEntities(mAssociatedEntities)

	return p
}

// applyTarget handles the gRPC request by passing it to the underlying Target Apply() method.
func (s *TargetServer) applyTarget(ctx context.Context, c *alpha.Client, request *alphapb.ApplyClouddeployAlphaTargetRequest) (*alphapb.ClouddeployAlphaTarget, error) {
	p := ProtoToTarget(request.GetResource())
	res, err := c.ApplyTarget(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetToProto(res)
	return r, nil
}

// applyClouddeployAlphaTarget handles the gRPC request by passing it to the underlying Target Apply() method.
func (s *TargetServer) ApplyClouddeployAlphaTarget(ctx context.Context, request *alphapb.ApplyClouddeployAlphaTargetRequest) (*alphapb.ClouddeployAlphaTarget, error) {
	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTarget(ctx, cl, request)
}

// DeleteTarget handles the gRPC request by passing it to the underlying Target Delete() method.
func (s *TargetServer) DeleteClouddeployAlphaTarget(ctx context.Context, request *alphapb.DeleteClouddeployAlphaTargetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTarget(ctx, ProtoToTarget(request.GetResource()))

}

// ListClouddeployAlphaTarget handles the gRPC request by passing it to the underlying TargetList() method.
func (s *TargetServer) ListClouddeployAlphaTarget(ctx context.Context, request *alphapb.ListClouddeployAlphaTargetRequest) (*alphapb.ListClouddeployAlphaTargetResponse, error) {
	cl, err := createConfigTarget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTarget(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ClouddeployAlphaTarget
	for _, r := range resources.Items {
		rp := TargetToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListClouddeployAlphaTargetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTarget(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
