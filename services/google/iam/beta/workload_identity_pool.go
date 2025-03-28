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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type WorkloadIdentityPool struct {
	Name        *string                        `json:"name"`
	DisplayName *string                        `json:"displayName"`
	Description *string                        `json:"description"`
	State       *WorkloadIdentityPoolStateEnum `json:"state"`
	Disabled    *bool                          `json:"disabled"`
	Project     *string                        `json:"project"`
	Location    *string                        `json:"location"`
}

func (r *WorkloadIdentityPool) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkloadIdentityPoolStateEnum.
type WorkloadIdentityPoolStateEnum string

// WorkloadIdentityPoolStateEnumRef returns a *WorkloadIdentityPoolStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadIdentityPoolStateEnumRef(s string) *WorkloadIdentityPoolStateEnum {
	v := WorkloadIdentityPoolStateEnum(s)
	return &v
}

func (v WorkloadIdentityPoolStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "ACTIVE", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadIdentityPoolStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *WorkloadIdentityPool) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "iam",
		Type:    "WorkloadIdentityPool",
		Version: "beta",
	}
}

func (r *WorkloadIdentityPool) ID() (string, error) {
	if err := extractWorkloadIdentityPoolFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":         dcl.ValueOrEmptyString(nr.Name),
		"display_name": dcl.ValueOrEmptyString(nr.DisplayName),
		"description":  dcl.ValueOrEmptyString(nr.Description),
		"state":        dcl.ValueOrEmptyString(nr.State),
		"disabled":     dcl.ValueOrEmptyString(nr.Disabled),
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"location":     dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{name}}", params), nil
}

const WorkloadIdentityPoolMaxPage = -1

type WorkloadIdentityPoolList struct {
	Items []*WorkloadIdentityPool

	nextToken string

	pageSize int32

	resource *WorkloadIdentityPool
}

func (l *WorkloadIdentityPoolList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkloadIdentityPoolList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkloadIdentityPool(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkloadIdentityPool(ctx context.Context, project, location string) (*WorkloadIdentityPoolList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListWorkloadIdentityPoolWithMaxResults(ctx, project, location, WorkloadIdentityPoolMaxPage)

}

func (c *Client) ListWorkloadIdentityPoolWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*WorkloadIdentityPoolList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &WorkloadIdentityPool{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listWorkloadIdentityPool(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkloadIdentityPoolList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) DeleteWorkloadIdentityPool(ctx context.Context, r *WorkloadIdentityPool) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("WorkloadIdentityPool resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting WorkloadIdentityPool...")
	deleteOp := deleteWorkloadIdentityPoolOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkloadIdentityPool deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkloadIdentityPool(ctx context.Context, project, location string, filter func(*WorkloadIdentityPool) bool) error {
	listObj, err := c.ListWorkloadIdentityPool(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkloadIdentityPool(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkloadIdentityPool(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkloadIdentityPool(ctx context.Context, rawDesired *WorkloadIdentityPool, opts ...dcl.ApplyOption) (*WorkloadIdentityPool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *WorkloadIdentityPool
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyWorkloadIdentityPoolHelper(c, ctx, rawDesired, opts...)
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

func applyWorkloadIdentityPoolHelper(c *Client, ctx context.Context, rawDesired *WorkloadIdentityPool, opts ...dcl.ApplyOption) (*WorkloadIdentityPool, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyWorkloadIdentityPool...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractWorkloadIdentityPoolFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.workloadIdentityPoolDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToWorkloadIdentityPoolDiffs(c.Config, fieldDiffs, opts)
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
	var ops []workloadIdentityPoolApiOperation
	if create {
		ops = append(ops, &createWorkloadIdentityPoolOperation{})
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
	return applyWorkloadIdentityPoolDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyWorkloadIdentityPoolDiff(c *Client, ctx context.Context, desired *WorkloadIdentityPool, rawDesired *WorkloadIdentityPool, ops []workloadIdentityPoolApiOperation, opts ...dcl.ApplyOption) (*WorkloadIdentityPool, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetWorkloadIdentityPool(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createWorkloadIdentityPoolOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapWorkloadIdentityPool(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeWorkloadIdentityPoolNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkloadIdentityPoolNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkloadIdentityPoolDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractWorkloadIdentityPoolFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractWorkloadIdentityPoolFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkloadIdentityPool(c, newDesired, newState)
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
