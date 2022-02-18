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
package identitytoolkit

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type OAuthIdpConfig struct{}

func OAuthIdpConfigToUnstructured(r *dclService.OAuthIdpConfig) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "identitytoolkit",
			Version: "alpha",
			Type:    "OAuthIdpConfig",
		},
		Object: make(map[string]interface{}),
	}
	if r.ClientId != nil {
		u.Object["clientId"] = *r.ClientId
	}
	if r.ClientSecret != nil {
		u.Object["clientSecret"] = *r.ClientSecret
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Enabled != nil {
		u.Object["enabled"] = *r.Enabled
	}
	if r.Issuer != nil {
		u.Object["issuer"] = *r.Issuer
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ResponseType != nil && r.ResponseType != dclService.EmptyOAuthIdpConfigResponseType {
		rResponseType := make(map[string]interface{})
		if r.ResponseType.Code != nil {
			rResponseType["code"] = *r.ResponseType.Code
		}
		if r.ResponseType.IdToken != nil {
			rResponseType["idToken"] = *r.ResponseType.IdToken
		}
		if r.ResponseType.Token != nil {
			rResponseType["token"] = *r.ResponseType.Token
		}
		u.Object["responseType"] = rResponseType
	}
	return u
}

func UnstructuredToOAuthIdpConfig(u *unstructured.Resource) (*dclService.OAuthIdpConfig, error) {
	r := &dclService.OAuthIdpConfig{}
	if _, ok := u.Object["clientId"]; ok {
		if s, ok := u.Object["clientId"].(string); ok {
			r.ClientId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClientId: expected string")
		}
	}
	if _, ok := u.Object["clientSecret"]; ok {
		if s, ok := u.Object["clientSecret"].(string); ok {
			r.ClientSecret = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClientSecret: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["enabled"]; ok {
		if b, ok := u.Object["enabled"].(bool); ok {
			r.Enabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Enabled: expected bool")
		}
	}
	if _, ok := u.Object["issuer"]; ok {
		if s, ok := u.Object["issuer"].(string); ok {
			r.Issuer = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Issuer: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["responseType"]; ok {
		if rResponseType, ok := u.Object["responseType"].(map[string]interface{}); ok {
			r.ResponseType = &dclService.OAuthIdpConfigResponseType{}
			if _, ok := rResponseType["code"]; ok {
				if b, ok := rResponseType["code"].(bool); ok {
					r.ResponseType.Code = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResponseType.Code: expected bool")
				}
			}
			if _, ok := rResponseType["idToken"]; ok {
				if b, ok := rResponseType["idToken"].(bool); ok {
					r.ResponseType.IdToken = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResponseType.IdToken: expected bool")
				}
			}
			if _, ok := rResponseType["token"]; ok {
				if b, ok := rResponseType["token"].(bool); ok {
					r.ResponseType.Token = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResponseType.Token: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResponseType: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetOAuthIdpConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOAuthIdpConfig(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetOAuthIdpConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	return OAuthIdpConfigToUnstructured(r), nil
}

func ListOAuthIdpConfig(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListOAuthIdpConfig(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, OAuthIdpConfigToUnstructured(r))
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

func ApplyOAuthIdpConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOAuthIdpConfig(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToOAuthIdpConfig(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyOAuthIdpConfig(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return OAuthIdpConfigToUnstructured(r), nil
}

func OAuthIdpConfigHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOAuthIdpConfig(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToOAuthIdpConfig(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyOAuthIdpConfig(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteOAuthIdpConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOAuthIdpConfig(u)
	if err != nil {
		return err
	}
	return c.DeleteOAuthIdpConfig(ctx, r)
}

func OAuthIdpConfigID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToOAuthIdpConfig(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *OAuthIdpConfig) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"identitytoolkit",
		"OAuthIdpConfig",
		"alpha",
	}
}

func (r *OAuthIdpConfig) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OAuthIdpConfig) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OAuthIdpConfig) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OAuthIdpConfig) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OAuthIdpConfig) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetOAuthIdpConfig(ctx, config, resource)
}

func (r *OAuthIdpConfig) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyOAuthIdpConfig(ctx, config, resource, opts...)
}

func (r *OAuthIdpConfig) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return OAuthIdpConfigHasDiff(ctx, config, resource, opts...)
}

func (r *OAuthIdpConfig) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteOAuthIdpConfig(ctx, config, resource)
}

func (r *OAuthIdpConfig) ID(resource *unstructured.Resource) (string, error) {
	return OAuthIdpConfigID(resource)
}

func init() {
	unstructured.Register(&OAuthIdpConfig{})
}
