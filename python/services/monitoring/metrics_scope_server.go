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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// MetricsScopeServer implements the gRPC interface for MetricsScope.
type MetricsScopeServer struct{}

// ProtoToMetricsScopeMonitoredProjects converts a MetricsScopeMonitoredProjects object from its proto representation.
func ProtoToMonitoringMetricsScopeMonitoredProjects(p *monitoringpb.MonitoringMetricsScopeMonitoredProjects) *monitoring.MetricsScopeMonitoredProjects {
	if p == nil {
		return nil
	}
	obj := &monitoring.MetricsScopeMonitoredProjects{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToMetricsScope converts a MetricsScope resource from its proto representation.
func ProtoToMetricsScope(p *monitoringpb.MonitoringMetricsScope) *monitoring.MetricsScope {
	obj := &monitoring.MetricsScope{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMonitoredProjects() {
		obj.MonitoredProjects = append(obj.MonitoredProjects, *ProtoToMonitoringMetricsScopeMonitoredProjects(r))
	}
	return obj
}

// MetricsScopeMonitoredProjectsToProto converts a MetricsScopeMonitoredProjects object to its proto representation.
func MonitoringMetricsScopeMonitoredProjectsToProto(o *monitoring.MetricsScopeMonitoredProjects) *monitoringpb.MonitoringMetricsScopeMonitoredProjects {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringMetricsScopeMonitoredProjects{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// MetricsScopeToProto converts a MetricsScope resource to its proto representation.
func MetricsScopeToProto(resource *monitoring.MetricsScope) *monitoringpb.MonitoringMetricsScope {
	p := &monitoringpb.MonitoringMetricsScope{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	sMonitoredProjects := make([]*monitoringpb.MonitoringMetricsScopeMonitoredProjects, len(resource.MonitoredProjects))
	for i, r := range resource.MonitoredProjects {
		sMonitoredProjects[i] = MonitoringMetricsScopeMonitoredProjectsToProto(&r)
	}
	p.SetMonitoredProjects(sMonitoredProjects)

	return p
}

// applyMetricsScope handles the gRPC request by passing it to the underlying MetricsScope Apply() method.
func (s *MetricsScopeServer) applyMetricsScope(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringMetricsScopeRequest) (*monitoringpb.MonitoringMetricsScope, error) {
	p := ProtoToMetricsScope(request.GetResource())
	res, err := c.ApplyMetricsScope(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricsScopeToProto(res)
	return r, nil
}

// applyMonitoringMetricsScope handles the gRPC request by passing it to the underlying MetricsScope Apply() method.
func (s *MetricsScopeServer) ApplyMonitoringMetricsScope(ctx context.Context, request *monitoringpb.ApplyMonitoringMetricsScopeRequest) (*monitoringpb.MonitoringMetricsScope, error) {
	cl, err := createConfigMetricsScope(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetricsScope(ctx, cl, request)
}

// DeleteMetricsScope handles the gRPC request by passing it to the underlying MetricsScope Delete() method.
func (s *MetricsScopeServer) DeleteMonitoringMetricsScope(ctx context.Context, request *monitoringpb.DeleteMonitoringMetricsScopeRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for MetricsScope")

}

// ListMonitoringMetricsScope is a no-op method because MetricsScope has no list method.
func (s *MetricsScopeServer) ListMonitoringMetricsScope(_ context.Context, _ *monitoringpb.ListMonitoringMetricsScopeRequest) (*monitoringpb.ListMonitoringMetricsScopeResponse, error) {
	return nil, nil
}

func createConfigMetricsScope(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
