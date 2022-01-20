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
package orgpolicy

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Policy struct{}

func PolicyToUnstructured(r *dclService.Policy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "orgpolicy",
			Version: "ga",
			Type:    "Policy",
		},
		Object: make(map[string]interface{}),
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.Spec != nil && r.Spec != dclService.EmptyPolicySpec {
		rSpec := make(map[string]interface{})
		if r.Spec.Etag != nil {
			rSpec["etag"] = *r.Spec.Etag
		}
		if r.Spec.InheritFromParent != nil {
			rSpec["inheritFromParent"] = *r.Spec.InheritFromParent
		}
		if r.Spec.Reset != nil {
			rSpec["reset"] = *r.Spec.Reset
		}
		var rSpecRules []interface{}
		for _, rSpecRulesVal := range r.Spec.Rules {
			rSpecRulesObject := make(map[string]interface{})
			if rSpecRulesVal.AllowAll != nil {
				rSpecRulesObject["allowAll"] = *rSpecRulesVal.AllowAll
			}
			if rSpecRulesVal.Condition != nil && rSpecRulesVal.Condition != dclService.EmptyPolicySpecRulesCondition {
				rSpecRulesValCondition := make(map[string]interface{})
				if rSpecRulesVal.Condition.Description != nil {
					rSpecRulesValCondition["description"] = *rSpecRulesVal.Condition.Description
				}
				if rSpecRulesVal.Condition.Expression != nil {
					rSpecRulesValCondition["expression"] = *rSpecRulesVal.Condition.Expression
				}
				if rSpecRulesVal.Condition.Location != nil {
					rSpecRulesValCondition["location"] = *rSpecRulesVal.Condition.Location
				}
				if rSpecRulesVal.Condition.Title != nil {
					rSpecRulesValCondition["title"] = *rSpecRulesVal.Condition.Title
				}
				rSpecRulesObject["condition"] = rSpecRulesValCondition
			}
			if rSpecRulesVal.DenyAll != nil {
				rSpecRulesObject["denyAll"] = *rSpecRulesVal.DenyAll
			}
			if rSpecRulesVal.Enforce != nil {
				rSpecRulesObject["enforce"] = *rSpecRulesVal.Enforce
			}
			if rSpecRulesVal.Values != nil && rSpecRulesVal.Values != dclService.EmptyPolicySpecRulesValues {
				rSpecRulesValValues := make(map[string]interface{})
				var rSpecRulesValValuesAllowedValues []interface{}
				for _, rSpecRulesValValuesAllowedValuesVal := range rSpecRulesVal.Values.AllowedValues {
					rSpecRulesValValuesAllowedValues = append(rSpecRulesValValuesAllowedValues, rSpecRulesValValuesAllowedValuesVal)
				}
				rSpecRulesValValues["allowedValues"] = rSpecRulesValValuesAllowedValues
				var rSpecRulesValValuesDeniedValues []interface{}
				for _, rSpecRulesValValuesDeniedValuesVal := range rSpecRulesVal.Values.DeniedValues {
					rSpecRulesValValuesDeniedValues = append(rSpecRulesValValuesDeniedValues, rSpecRulesValValuesDeniedValuesVal)
				}
				rSpecRulesValValues["deniedValues"] = rSpecRulesValValuesDeniedValues
				rSpecRulesObject["values"] = rSpecRulesValValues
			}
			rSpecRules = append(rSpecRules, rSpecRulesObject)
		}
		rSpec["rules"] = rSpecRules
		if r.Spec.UpdateTime != nil {
			rSpec["updateTime"] = *r.Spec.UpdateTime
		}
		u.Object["spec"] = rSpec
	}
	return u
}

func UnstructuredToPolicy(u *unstructured.Resource) (*dclService.Policy, error) {
	r := &dclService.Policy{}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["spec"]; ok {
		if rSpec, ok := u.Object["spec"].(map[string]interface{}); ok {
			r.Spec = &dclService.PolicySpec{}
			if _, ok := rSpec["etag"]; ok {
				if s, ok := rSpec["etag"].(string); ok {
					r.Spec.Etag = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Spec.Etag: expected string")
				}
			}
			if _, ok := rSpec["inheritFromParent"]; ok {
				if b, ok := rSpec["inheritFromParent"].(bool); ok {
					r.Spec.InheritFromParent = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Spec.InheritFromParent: expected bool")
				}
			}
			if _, ok := rSpec["reset"]; ok {
				if b, ok := rSpec["reset"].(bool); ok {
					r.Spec.Reset = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Spec.Reset: expected bool")
				}
			}
			if _, ok := rSpec["rules"]; ok {
				if s, ok := rSpec["rules"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rSpecRules dclService.PolicySpecRules
							if _, ok := objval["allowAll"]; ok {
								if b, ok := objval["allowAll"].(bool); ok {
									rSpecRules.AllowAll = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rSpecRules.AllowAll: expected bool")
								}
							}
							if _, ok := objval["condition"]; ok {
								if rSpecRulesCondition, ok := objval["condition"].(map[string]interface{}); ok {
									rSpecRules.Condition = &dclService.PolicySpecRulesCondition{}
									if _, ok := rSpecRulesCondition["description"]; ok {
										if s, ok := rSpecRulesCondition["description"].(string); ok {
											rSpecRules.Condition.Description = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rSpecRules.Condition.Description: expected string")
										}
									}
									if _, ok := rSpecRulesCondition["expression"]; ok {
										if s, ok := rSpecRulesCondition["expression"].(string); ok {
											rSpecRules.Condition.Expression = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rSpecRules.Condition.Expression: expected string")
										}
									}
									if _, ok := rSpecRulesCondition["location"]; ok {
										if s, ok := rSpecRulesCondition["location"].(string); ok {
											rSpecRules.Condition.Location = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rSpecRules.Condition.Location: expected string")
										}
									}
									if _, ok := rSpecRulesCondition["title"]; ok {
										if s, ok := rSpecRulesCondition["title"].(string); ok {
											rSpecRules.Condition.Title = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rSpecRules.Condition.Title: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rSpecRules.Condition: expected map[string]interface{}")
								}
							}
							if _, ok := objval["denyAll"]; ok {
								if b, ok := objval["denyAll"].(bool); ok {
									rSpecRules.DenyAll = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rSpecRules.DenyAll: expected bool")
								}
							}
							if _, ok := objval["enforce"]; ok {
								if b, ok := objval["enforce"].(bool); ok {
									rSpecRules.Enforce = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rSpecRules.Enforce: expected bool")
								}
							}
							if _, ok := objval["values"]; ok {
								if rSpecRulesValues, ok := objval["values"].(map[string]interface{}); ok {
									rSpecRules.Values = &dclService.PolicySpecRulesValues{}
									if _, ok := rSpecRulesValues["allowedValues"]; ok {
										if s, ok := rSpecRulesValues["allowedValues"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rSpecRules.Values.AllowedValues = append(rSpecRules.Values.AllowedValues, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rSpecRules.Values.AllowedValues: expected []interface{}")
										}
									}
									if _, ok := rSpecRulesValues["deniedValues"]; ok {
										if s, ok := rSpecRulesValues["deniedValues"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rSpecRules.Values.DeniedValues = append(rSpecRules.Values.DeniedValues, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rSpecRules.Values.DeniedValues: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rSpecRules.Values: expected map[string]interface{}")
								}
							}
							r.Spec.Rules = append(r.Spec.Rules, rSpecRules)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Spec.Rules: expected []interface{}")
				}
			}
			if _, ok := rSpec["updateTime"]; ok {
				if s, ok := rSpec["updateTime"].(string); ok {
					r.Spec.UpdateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Spec.UpdateTime: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Spec: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return PolicyToUnstructured(r), nil
}

func ListPolicy(ctx context.Context, config *dcl.Config, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListPolicy(ctx, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, PolicyToUnstructured(r))
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

func ApplyPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return PolicyToUnstructured(r), nil
}

func PolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeletePolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return err
	}
	return c.DeletePolicy(ctx, r)
}

func PolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Policy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"orgpolicy",
		"Policy",
		"ga",
	}
}

func (r *Policy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Policy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Policy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicy(ctx, config, resource)
}

func (r *Policy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPolicy(ctx, config, resource, opts...)
}

func (r *Policy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PolicyHasDiff(ctx, config, resource, opts...)
}

func (r *Policy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePolicy(ctx, config, resource)
}

func (r *Policy) ID(resource *unstructured.Resource) (string, error) {
	return PolicyID(resource)
}

func init() {
	unstructured.Register(&Policy{})
}
