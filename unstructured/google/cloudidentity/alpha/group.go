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
package cloudidentity

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Group struct{}

func GroupToUnstructured(r *dclService.Group) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudidentity",
			Version: "alpha",
			Type:    "Group",
		},
		Object: make(map[string]interface{}),
	}
	var rAdditionalGroupKeys []interface{}
	for _, rAdditionalGroupKeysVal := range r.AdditionalGroupKeys {
		rAdditionalGroupKeys = append(rAdditionalGroupKeys, GroupGoogleappscloudidentitygroupsvxentitykeyToUnstructured(&rAdditionalGroupKeysVal))
	}
	u.Object["additionalGroupKeys"] = rAdditionalGroupKeys
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.DynamicGroupMetadata != nil && r.DynamicGroupMetadata != dclService.EmptyGroupDynamicGroupMetadata {
		rDynamicGroupMetadata := make(map[string]interface{})
		var rDynamicGroupMetadataQueries []interface{}
		for _, rDynamicGroupMetadataQueriesVal := range r.DynamicGroupMetadata.Queries {
			rDynamicGroupMetadataQueriesObject := make(map[string]interface{})
			if rDynamicGroupMetadataQueriesVal.Query != nil {
				rDynamicGroupMetadataQueriesObject["query"] = *rDynamicGroupMetadataQueriesVal.Query
			}
			if rDynamicGroupMetadataQueriesVal.ResourceType != nil {
				rDynamicGroupMetadataQueriesObject["resourceType"] = string(*rDynamicGroupMetadataQueriesVal.ResourceType)
			}
			rDynamicGroupMetadataQueries = append(rDynamicGroupMetadataQueries, rDynamicGroupMetadataQueriesObject)
		}
		rDynamicGroupMetadata["queries"] = rDynamicGroupMetadataQueries
		if r.DynamicGroupMetadata.Status != nil && r.DynamicGroupMetadata.Status != dclService.EmptyGroupDynamicGroupMetadataStatus {
			rDynamicGroupMetadataStatus := make(map[string]interface{})
			if r.DynamicGroupMetadata.Status.Status != nil {
				rDynamicGroupMetadataStatus["status"] = string(*r.DynamicGroupMetadata.Status.Status)
			}
			if r.DynamicGroupMetadata.Status.StatusTime != nil {
				rDynamicGroupMetadataStatus["statusTime"] = *r.DynamicGroupMetadata.Status.StatusTime
			}
			rDynamicGroupMetadata["status"] = rDynamicGroupMetadataStatus
		}
		u.Object["dynamicGroupMetadata"] = rDynamicGroupMetadata
	}
	if r.GroupKey != nil {
		u.Object["groupKey"] = GroupGoogleappscloudidentitygroupsvxentitykeyToUnstructured(r.GroupKey)
	}
	if r.InitialGroupConfig != nil {
		u.Object["initialGroupConfig"] = string(*r.InitialGroupConfig)
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	var rPosixGroups []interface{}
	for _, rPosixGroupsVal := range r.PosixGroups {
		rPosixGroupsObject := make(map[string]interface{})
		if rPosixGroupsVal.Gid != nil {
			rPosixGroupsObject["gid"] = *rPosixGroupsVal.Gid
		}
		if rPosixGroupsVal.Name != nil {
			rPosixGroupsObject["name"] = *rPosixGroupsVal.Name
		}
		if rPosixGroupsVal.SystemId != nil {
			rPosixGroupsObject["systemId"] = *rPosixGroupsVal.SystemId
		}
		rPosixGroups = append(rPosixGroups, rPosixGroupsObject)
	}
	u.Object["posixGroups"] = rPosixGroups
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func GroupGoogleappscloudidentitygroupsvxentitykeyToUnstructured(r *dclService.GroupGoogleappscloudidentitygroupsvxentitykey) map[string]interface{} {
	result := make(map[string]interface{})
	if r.Id != nil {
		result["id"] = *r.Id
	}
	if r.Namespace != nil {
		result["namespace"] = *r.Namespace
	}
	return result
}

func UnstructuredToGroup(u *unstructured.Resource) (*dclService.Group, error) {
	r := &dclService.Group{}
	if _, ok := u.Object["additionalGroupKeys"]; ok {
		if s, ok := u.Object["additionalGroupKeys"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					unstructuredObjval, err := UnstructuredToGroupGoogleappscloudidentitygroupsvxentitykey(objval)
					if err != nil {
						return nil, err
					}
					r.AdditionalGroupKeys = append(r.AdditionalGroupKeys, *unstructuredObjval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AdditionalGroupKeys: expected []interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["dynamicGroupMetadata"]; ok {
		if rDynamicGroupMetadata, ok := u.Object["dynamicGroupMetadata"].(map[string]interface{}); ok {
			r.DynamicGroupMetadata = &dclService.GroupDynamicGroupMetadata{}
			if _, ok := rDynamicGroupMetadata["queries"]; ok {
				if s, ok := rDynamicGroupMetadata["queries"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rDynamicGroupMetadataQueries dclService.GroupDynamicGroupMetadataQueries
							if _, ok := objval["query"]; ok {
								if s, ok := objval["query"].(string); ok {
									rDynamicGroupMetadataQueries.Query = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDynamicGroupMetadataQueries.Query: expected string")
								}
							}
							if _, ok := objval["resourceType"]; ok {
								if s, ok := objval["resourceType"].(string); ok {
									rDynamicGroupMetadataQueries.ResourceType = dclService.GroupDynamicGroupMetadataQueriesResourceTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("rDynamicGroupMetadataQueries.ResourceType: expected string")
								}
							}
							r.DynamicGroupMetadata.Queries = append(r.DynamicGroupMetadata.Queries, rDynamicGroupMetadataQueries)
						}
					}
				} else {
					return nil, fmt.Errorf("r.DynamicGroupMetadata.Queries: expected []interface{}")
				}
			}
			if _, ok := rDynamicGroupMetadata["status"]; ok {
				if rDynamicGroupMetadataStatus, ok := rDynamicGroupMetadata["status"].(map[string]interface{}); ok {
					r.DynamicGroupMetadata.Status = &dclService.GroupDynamicGroupMetadataStatus{}
					if _, ok := rDynamicGroupMetadataStatus["status"]; ok {
						if s, ok := rDynamicGroupMetadataStatus["status"].(string); ok {
							r.DynamicGroupMetadata.Status.Status = dclService.GroupDynamicGroupMetadataStatusStatusEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.DynamicGroupMetadata.Status.Status: expected string")
						}
					}
					if _, ok := rDynamicGroupMetadataStatus["statusTime"]; ok {
						if s, ok := rDynamicGroupMetadataStatus["statusTime"].(string); ok {
							r.DynamicGroupMetadata.Status.StatusTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.DynamicGroupMetadata.Status.StatusTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.DynamicGroupMetadata.Status: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DynamicGroupMetadata: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["groupKey"]; ok {
		if rGroupKey, ok := u.Object["groupKey"].(map[string]interface{}); ok {
			var err error
			r.GroupKey, err = UnstructuredToGroupGoogleappscloudidentitygroupsvxentitykey(rGroupKey)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("r.GroupKey: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["initialGroupConfig"]; ok {
		if s, ok := u.Object["initialGroupConfig"].(string); ok {
			r.InitialGroupConfig = dclService.GroupInitialGroupConfigEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.InitialGroupConfig: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
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
	if _, ok := u.Object["posixGroups"]; ok {
		if s, ok := u.Object["posixGroups"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rPosixGroups dclService.GroupPosixGroups
					if _, ok := objval["gid"]; ok {
						if s, ok := objval["gid"].(string); ok {
							rPosixGroups.Gid = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rPosixGroups.Gid: expected string")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rPosixGroups.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rPosixGroups.Name: expected string")
						}
					}
					if _, ok := objval["systemId"]; ok {
						if s, ok := objval["systemId"].(string); ok {
							rPosixGroups.SystemId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rPosixGroups.SystemId: expected string")
						}
					}
					r.PosixGroups = append(r.PosixGroups, rPosixGroups)
				}
			}
		} else {
			return nil, fmt.Errorf("r.PosixGroups: expected []interface{}")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func UnstructuredToGroupGoogleappscloudidentitygroupsvxentitykey(obj map[string]interface{}) (*dclService.GroupGoogleappscloudidentitygroupsvxentitykey, error) {
	r := &dclService.GroupGoogleappscloudidentitygroupsvxentitykey{}
	if _, ok := obj["id"]; ok {
		if s, ok := obj["id"].(string); ok {
			r.Id = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Id: expected string")
		}
	}
	if _, ok := obj["namespace"]; ok {
		if s, ok := obj["namespace"].(string); ok {
			r.Namespace = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Namespace: expected string")
		}
	}
	return r, nil
}

func GetGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetGroup(ctx, r)
	if err != nil {
		return nil, err
	}
	return GroupToUnstructured(r), nil
}

func ListGroup(ctx context.Context, config *dcl.Config, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListGroup(ctx, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, GroupToUnstructured(r))
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

func ApplyGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGroup(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyGroup(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return GroupToUnstructured(r), nil
}

func GroupHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGroup(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyGroup(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteGroup(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return err
	}
	return c.DeleteGroup(ctx, r)
}

func GroupID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToGroup(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Group) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudidentity",
		"Group",
		"alpha",
	}
}

func (r *Group) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Group) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetGroup(ctx, config, resource)
}

func (r *Group) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyGroup(ctx, config, resource, opts...)
}

func (r *Group) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return GroupHasDiff(ctx, config, resource, opts...)
}

func (r *Group) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteGroup(ctx, config, resource)
}

func (r *Group) ID(resource *unstructured.Resource) (string, error) {
	return GroupID(resource)
}

func init() {
	unstructured.Register(&Group{})
}
