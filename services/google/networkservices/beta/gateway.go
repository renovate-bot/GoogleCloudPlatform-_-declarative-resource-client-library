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
package beta

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Gateway struct {
	Name            *string           `json:"name"`
	CreateTime      *string           `json:"createTime"`
	UpdateTime      *string           `json:"updateTime"`
	Labels          map[string]string `json:"labels"`
	Description     *string           `json:"description"`
	Type            *GatewayTypeEnum  `json:"type"`
	Addresses       []string          `json:"addresses"`
	Ports           []int64           `json:"ports"`
	Scope           *string           `json:"scope"`
	ServerTlsPolicy *string           `json:"serverTlsPolicy"`
	Project         *string           `json:"project"`
	Location        *string           `json:"location"`
	SelfLink        *string           `json:"selfLink"`
}

func (r *Gateway) String() string {
	return dcl.SprintResource(r)
}

// The enum GatewayTypeEnum.
type GatewayTypeEnum string

// GatewayTypeEnumRef returns a *GatewayTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GatewayTypeEnumRef(s string) *GatewayTypeEnum {
	v := GatewayTypeEnum(s)
	return &v
}

func (v GatewayTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TYPE_UNSPECIFIED", "OPEN_MESH", "SECURE_WEB_GATEWAY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GatewayTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Gateway) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "Gateway",
		Version: "beta",
	}
}

func (r *Gateway) ID() (string, error) {
	if err := extractGatewayFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":              dcl.ValueOrEmptyString(nr.Name),
		"create_time":       dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":       dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":            dcl.ValueOrEmptyString(nr.Labels),
		"description":       dcl.ValueOrEmptyString(nr.Description),
		"type":              dcl.ValueOrEmptyString(nr.Type),
		"addresses":         dcl.ValueOrEmptyString(nr.Addresses),
		"ports":             dcl.ValueOrEmptyString(nr.Ports),
		"scope":             dcl.ValueOrEmptyString(nr.Scope),
		"server_tls_policy": dcl.ValueOrEmptyString(nr.ServerTlsPolicy),
		"project":           dcl.ValueOrEmptyString(nr.Project),
		"location":          dcl.ValueOrEmptyString(nr.Location),
		"self_link":         dcl.ValueOrEmptyString(nr.SelfLink),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/gateways/{{name}}", params), nil
}

const GatewayMaxPage = -1

type GatewayList struct {
	Items []*Gateway

	nextToken string

	pageSize int32

	resource *Gateway
}

func (l *GatewayList) HasNext() bool {
	return l.nextToken != ""
}

func (l *GatewayList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listGateway(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListGateway(ctx context.Context, project, location string) (*GatewayList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListGatewayWithMaxResults(ctx, project, location, GatewayMaxPage)

}

func (c *Client) ListGatewayWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*GatewayList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Gateway{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listGateway(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &GatewayList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetGateway(ctx context.Context, r *Gateway) (*Gateway, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractGatewayFields(r)

	b, err := c.getGatewayRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalGateway(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeGatewayNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractGatewayFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteGateway(ctx context.Context, r *Gateway) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Gateway resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Gateway...")
	deleteOp := deleteGatewayOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllGateway deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllGateway(ctx context.Context, project, location string, filter func(*Gateway) bool) error {
	listObj, err := c.ListGateway(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllGateway(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllGateway(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyGateway(ctx context.Context, rawDesired *Gateway, opts ...dcl.ApplyOption) (*Gateway, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Gateway
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyGatewayHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyGatewayHelper(c *Client, ctx context.Context, rawDesired *Gateway, opts ...dcl.ApplyOption) (*Gateway, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyGateway...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractGatewayFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.gatewayDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToGatewayDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []gatewayApiOperation
	if create {
		ops = append(ops, &createGatewayOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyGatewayDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyGatewayDiff(c *Client, ctx context.Context, desired *Gateway, rawDesired *Gateway, ops []gatewayApiOperation, opts ...dcl.ApplyOption) (*Gateway, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetGateway(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createGatewayOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapGateway(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeGatewayNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeGatewayNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeGatewayDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractGatewayFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractGatewayFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffGateway(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
