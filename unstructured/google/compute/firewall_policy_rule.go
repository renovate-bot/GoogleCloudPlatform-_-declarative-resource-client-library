// Copyright 2022 Google LLC. All Rights Reserved.
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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type FirewallPolicyRule struct{}

func FirewallPolicyRuleToUnstructured(r *dclService.FirewallPolicyRule) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "ga",
			Type:    "FirewallPolicyRule",
		},
		Object: make(map[string]interface{}),
	}
	if r.Action != nil {
		u.Object["action"] = *r.Action
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Direction != nil {
		u.Object["direction"] = string(*r.Direction)
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.EnableLogging != nil {
		u.Object["enableLogging"] = *r.EnableLogging
	}
	if r.FirewallPolicy != nil {
		u.Object["firewallPolicy"] = *r.FirewallPolicy
	}
	if r.Kind != nil {
		u.Object["kind"] = *r.Kind
	}
	if r.Match != nil && r.Match != dclService.EmptyFirewallPolicyRuleMatch {
		rMatch := make(map[string]interface{})
		var rMatchDestIPRanges []interface{}
		for _, rMatchDestIPRangesVal := range r.Match.DestIPRanges {
			rMatchDestIPRanges = append(rMatchDestIPRanges, rMatchDestIPRangesVal)
		}
		rMatch["destIPRanges"] = rMatchDestIPRanges
		var rMatchLayer4Configs []interface{}
		for _, rMatchLayer4ConfigsVal := range r.Match.Layer4Configs {
			rMatchLayer4ConfigsObject := make(map[string]interface{})
			if rMatchLayer4ConfigsVal.IPProtocol != nil {
				rMatchLayer4ConfigsObject["ipProtocol"] = *rMatchLayer4ConfigsVal.IPProtocol
			}
			var rMatchLayer4ConfigsValPorts []interface{}
			for _, rMatchLayer4ConfigsValPortsVal := range rMatchLayer4ConfigsVal.Ports {
				rMatchLayer4ConfigsValPorts = append(rMatchLayer4ConfigsValPorts, rMatchLayer4ConfigsValPortsVal)
			}
			rMatchLayer4ConfigsObject["ports"] = rMatchLayer4ConfigsValPorts
			rMatchLayer4Configs = append(rMatchLayer4Configs, rMatchLayer4ConfigsObject)
		}
		rMatch["layer4Configs"] = rMatchLayer4Configs
		var rMatchSrcIPRanges []interface{}
		for _, rMatchSrcIPRangesVal := range r.Match.SrcIPRanges {
			rMatchSrcIPRanges = append(rMatchSrcIPRanges, rMatchSrcIPRangesVal)
		}
		rMatch["srcIPRanges"] = rMatchSrcIPRanges
		u.Object["match"] = rMatch
	}
	if r.Priority != nil {
		u.Object["priority"] = *r.Priority
	}
	if r.RuleTupleCount != nil {
		u.Object["ruleTupleCount"] = *r.RuleTupleCount
	}
	var rTargetResources []interface{}
	for _, rTargetResourcesVal := range r.TargetResources {
		rTargetResources = append(rTargetResources, rTargetResourcesVal)
	}
	u.Object["targetResources"] = rTargetResources
	var rTargetServiceAccounts []interface{}
	for _, rTargetServiceAccountsVal := range r.TargetServiceAccounts {
		rTargetServiceAccounts = append(rTargetServiceAccounts, rTargetServiceAccountsVal)
	}
	u.Object["targetServiceAccounts"] = rTargetServiceAccounts
	return u
}

func UnstructuredToFirewallPolicyRule(u *unstructured.Resource) (*dclService.FirewallPolicyRule, error) {
	r := &dclService.FirewallPolicyRule{}
	if _, ok := u.Object["action"]; ok {
		if s, ok := u.Object["action"].(string); ok {
			r.Action = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Action: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["direction"]; ok {
		if s, ok := u.Object["direction"].(string); ok {
			r.Direction = dclService.FirewallPolicyRuleDirectionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Direction: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["enableLogging"]; ok {
		if b, ok := u.Object["enableLogging"].(bool); ok {
			r.EnableLogging = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableLogging: expected bool")
		}
	}
	if _, ok := u.Object["firewallPolicy"]; ok {
		if s, ok := u.Object["firewallPolicy"].(string); ok {
			r.FirewallPolicy = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.FirewallPolicy: expected string")
		}
	}
	if _, ok := u.Object["kind"]; ok {
		if s, ok := u.Object["kind"].(string); ok {
			r.Kind = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Kind: expected string")
		}
	}
	if _, ok := u.Object["match"]; ok {
		if rMatch, ok := u.Object["match"].(map[string]interface{}); ok {
			r.Match = &dclService.FirewallPolicyRuleMatch{}
			if _, ok := rMatch["destIPRanges"]; ok {
				if s, ok := rMatch["destIPRanges"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.DestIPRanges = append(r.Match.DestIPRanges, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.DestIPRanges: expected []interface{}")
				}
			}
			if _, ok := rMatch["layer4Configs"]; ok {
				if s, ok := rMatch["layer4Configs"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMatchLayer4Configs dclService.FirewallPolicyRuleMatchLayer4Configs
							if _, ok := objval["ipProtocol"]; ok {
								if s, ok := objval["ipProtocol"].(string); ok {
									rMatchLayer4Configs.IPProtocol = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMatchLayer4Configs.IPProtocol: expected string")
								}
							}
							if _, ok := objval["ports"]; ok {
								if s, ok := objval["ports"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rMatchLayer4Configs.Ports = append(rMatchLayer4Configs.Ports, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rMatchLayer4Configs.Ports: expected []interface{}")
								}
							}
							r.Match.Layer4Configs = append(r.Match.Layer4Configs, rMatchLayer4Configs)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.Layer4Configs: expected []interface{}")
				}
			}
			if _, ok := rMatch["srcIPRanges"]; ok {
				if s, ok := rMatch["srcIPRanges"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Match.SrcIPRanges = append(r.Match.SrcIPRanges, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Match.SrcIPRanges: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Match: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["priority"]; ok {
		if i, ok := u.Object["priority"].(int64); ok {
			r.Priority = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Priority: expected int64")
		}
	}
	if _, ok := u.Object["ruleTupleCount"]; ok {
		if i, ok := u.Object["ruleTupleCount"].(int64); ok {
			r.RuleTupleCount = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.RuleTupleCount: expected int64")
		}
	}
	if _, ok := u.Object["targetResources"]; ok {
		if s, ok := u.Object["targetResources"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.TargetResources = append(r.TargetResources, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TargetResources: expected []interface{}")
		}
	}
	if _, ok := u.Object["targetServiceAccounts"]; ok {
		if s, ok := u.Object["targetServiceAccounts"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.TargetServiceAccounts = append(r.TargetServiceAccounts, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TargetServiceAccounts: expected []interface{}")
		}
	}
	return r, nil
}

func GetFirewallPolicyRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyRule(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFirewallPolicyRule(ctx, r)
	if err != nil {
		return nil, err
	}
	return FirewallPolicyRuleToUnstructured(r), nil
}

func ListFirewallPolicyRule(ctx context.Context, config *dcl.Config, firewallpolicy string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFirewallPolicyRule(ctx, firewallpolicy)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FirewallPolicyRuleToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyFirewallPolicyRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyRule(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirewallPolicyRule(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFirewallPolicyRule(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FirewallPolicyRuleToUnstructured(r), nil
}

func FirewallPolicyRuleHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyRule(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirewallPolicyRule(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFirewallPolicyRule(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFirewallPolicyRule(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyRule(u)
	if err != nil {
		return err
	}
	return c.DeleteFirewallPolicyRule(ctx, r)
}

func FirewallPolicyRuleID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFirewallPolicyRule(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *FirewallPolicyRule) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"FirewallPolicyRule",
		"ga",
	}
}

func (r *FirewallPolicyRule) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyRule) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyRule) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyRule) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyRule) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFirewallPolicyRule(ctx, config, resource)
}

func (r *FirewallPolicyRule) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFirewallPolicyRule(ctx, config, resource, opts...)
}

func (r *FirewallPolicyRule) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FirewallPolicyRuleHasDiff(ctx, config, resource, opts...)
}

func (r *FirewallPolicyRule) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFirewallPolicyRule(ctx, config, resource)
}

func (r *FirewallPolicyRule) ID(resource *unstructured.Resource) (string, error) {
	return FirewallPolicyRuleID(resource)
}

func init() {
	unstructured.Register(&FirewallPolicyRule{})
}