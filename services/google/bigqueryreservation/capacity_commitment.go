// Copyright 2021 Google LLC. All Rights Reserved.
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
package bigqueryreservation

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type CapacityCommitment struct {
	Name                *string                            `json:"name"`
	SlotCount           *int64                             `json:"slotCount"`
	Plan                *CapacityCommitmentPlanEnum        `json:"plan"`
	State               *CapacityCommitmentStateEnum       `json:"state"`
	CommitmentStartTime *string                            `json:"commitmentStartTime"`
	CommitmentEndTime   *string                            `json:"commitmentEndTime"`
	FailureStatus       *CapacityCommitmentFailureStatus   `json:"failureStatus"`
	RenewalPlan         *CapacityCommitmentRenewalPlanEnum `json:"renewalPlan"`
	Project             *string                            `json:"project"`
	Location            *string                            `json:"location"`
}

func (r *CapacityCommitment) String() string {
	return dcl.SprintResource(r)
}

// The enum CapacityCommitmentPlanEnum.
type CapacityCommitmentPlanEnum string

// CapacityCommitmentPlanEnumRef returns a *CapacityCommitmentPlanEnum with the value of string s
// If the empty string is provided, nil is returned.
func CapacityCommitmentPlanEnumRef(s string) *CapacityCommitmentPlanEnum {
	if s == "" {
		return nil
	}

	v := CapacityCommitmentPlanEnum(s)
	return &v
}

func (v CapacityCommitmentPlanEnum) Validate() error {
	for _, s := range []string{"COMMITMENT_PLAN_UNSPECIFIED", "FLEX", "TRIAL", "MONTHLY", "ANNUAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CapacityCommitmentPlanEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CapacityCommitmentStateEnum.
type CapacityCommitmentStateEnum string

// CapacityCommitmentStateEnumRef returns a *CapacityCommitmentStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func CapacityCommitmentStateEnumRef(s string) *CapacityCommitmentStateEnum {
	if s == "" {
		return nil
	}

	v := CapacityCommitmentStateEnum(s)
	return &v
}

func (v CapacityCommitmentStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "PENDING", "ACTIVE", "FAILED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CapacityCommitmentStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CapacityCommitmentRenewalPlanEnum.
type CapacityCommitmentRenewalPlanEnum string

// CapacityCommitmentRenewalPlanEnumRef returns a *CapacityCommitmentRenewalPlanEnum with the value of string s
// If the empty string is provided, nil is returned.
func CapacityCommitmentRenewalPlanEnumRef(s string) *CapacityCommitmentRenewalPlanEnum {
	if s == "" {
		return nil
	}

	v := CapacityCommitmentRenewalPlanEnum(s)
	return &v
}

func (v CapacityCommitmentRenewalPlanEnum) Validate() error {
	for _, s := range []string{"COMMITMENT_PLAN_UNSPECIFIED", "FLEX", "TRIAL", "MONTHLY", "ANNUAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CapacityCommitmentRenewalPlanEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type CapacityCommitmentFailureStatus struct {
	empty   bool                                     `json:"-"`
	Code    *int64                                   `json:"code"`
	Message *string                                  `json:"message"`
	Details []CapacityCommitmentFailureStatusDetails `json:"details"`
}

type jsonCapacityCommitmentFailureStatus CapacityCommitmentFailureStatus

func (r *CapacityCommitmentFailureStatus) UnmarshalJSON(data []byte) error {
	var res jsonCapacityCommitmentFailureStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCapacityCommitmentFailureStatus
	} else {

		r.Code = res.Code

		r.Message = res.Message

		r.Details = res.Details

	}
	return nil
}

// This object is used to assert a desired state where this CapacityCommitmentFailureStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCapacityCommitmentFailureStatus *CapacityCommitmentFailureStatus = &CapacityCommitmentFailureStatus{empty: true}

func (r *CapacityCommitmentFailureStatus) Empty() bool {
	return r.empty
}

func (r *CapacityCommitmentFailureStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *CapacityCommitmentFailureStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CapacityCommitmentFailureStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

type jsonCapacityCommitmentFailureStatusDetails CapacityCommitmentFailureStatusDetails

func (r *CapacityCommitmentFailureStatusDetails) UnmarshalJSON(data []byte) error {
	var res jsonCapacityCommitmentFailureStatusDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCapacityCommitmentFailureStatusDetails
	} else {

		r.TypeUrl = res.TypeUrl

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CapacityCommitmentFailureStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCapacityCommitmentFailureStatusDetails *CapacityCommitmentFailureStatusDetails = &CapacityCommitmentFailureStatusDetails{empty: true}

func (r *CapacityCommitmentFailureStatusDetails) Empty() bool {
	return r.empty
}

func (r *CapacityCommitmentFailureStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *CapacityCommitmentFailureStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *CapacityCommitment) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery_reservation",
		Type:    "CapacityCommitment",
		Version: "bigqueryreservation",
	}
}

func (r *CapacityCommitment) ID() (string, error) {
	if err := extractCapacityCommitmentFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                dcl.ValueOrEmptyString(nr.Name),
		"slotCount":           dcl.ValueOrEmptyString(nr.SlotCount),
		"plan":                dcl.ValueOrEmptyString(nr.Plan),
		"state":               dcl.ValueOrEmptyString(nr.State),
		"commitmentStartTime": dcl.ValueOrEmptyString(nr.CommitmentStartTime),
		"commitmentEndTime":   dcl.ValueOrEmptyString(nr.CommitmentEndTime),
		"failureStatus":       dcl.ValueOrEmptyString(nr.FailureStatus),
		"renewalPlan":         dcl.ValueOrEmptyString(nr.RenewalPlan),
		"project":             dcl.ValueOrEmptyString(nr.Project),
		"location":            dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}", params), nil
}

const CapacityCommitmentMaxPage = -1

type CapacityCommitmentList struct {
	Items []*CapacityCommitment

	nextToken string

	pageSize int32

	resource *CapacityCommitment
}

func (l *CapacityCommitmentList) HasNext() bool {
	return l.nextToken != ""
}

func (l *CapacityCommitmentList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listCapacityCommitment(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListCapacityCommitment(ctx context.Context, r *CapacityCommitment) (*CapacityCommitmentList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		429: dcl.Retryability{
			Retryable: false,
			Pattern:   "",
			Timeout:   0,
		},
	})))
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListCapacityCommitmentWithMaxResults(ctx, r, CapacityCommitmentMaxPage)

}

func (c *Client) ListCapacityCommitmentWithMaxResults(ctx context.Context, r *CapacityCommitment, pageSize int32) (*CapacityCommitmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listCapacityCommitment(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &CapacityCommitmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetCapacityCommitment(ctx context.Context, r *CapacityCommitment) (*CapacityCommitment, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		429: dcl.Retryability{
			Retryable: false,
			Pattern:   "",
			Timeout:   0,
		},
	})))
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractCapacityCommitmentFields(r)

	b, err := c.getCapacityCommitmentRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalCapacityCommitment(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeCapacityCommitmentNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteCapacityCommitment(ctx context.Context, r *CapacityCommitment) error {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		429: dcl.Retryability{
			Retryable: false,
			Pattern:   "",
			Timeout:   0,
		},
	})))
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("CapacityCommitment resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting CapacityCommitment...")
	deleteOp := deleteCapacityCommitmentOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllCapacityCommitment deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllCapacityCommitment(ctx context.Context, project, location string, filter func(*CapacityCommitment) bool) error {
	r := &CapacityCommitment{
		Project:  &project,
		Location: &location,
	}
	listObj, err := c.ListCapacityCommitment(ctx, r)
	if err != nil {
		return err
	}

	err = c.deleteAllCapacityCommitment(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllCapacityCommitment(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyCapacityCommitment(ctx context.Context, rawDesired *CapacityCommitment, opts ...dcl.ApplyOption) (*CapacityCommitment, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		429: dcl.Retryability{
			Retryable: false,
			Pattern:   "",
			Timeout:   0,
		},
	})))
	var resultNewState *CapacityCommitment
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyCapacityCommitmentHelper(c, ctx, rawDesired, opts...)
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

func applyCapacityCommitmentHelper(c *Client, ctx context.Context, rawDesired *CapacityCommitment, opts ...dcl.ApplyOption) (*CapacityCommitment, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyCapacityCommitment...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractCapacityCommitmentFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.capacityCommitmentDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToCapacityCommitmentDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
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
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []capacityCommitmentApiOperation
	if create {
		ops = append(ops, &createCapacityCommitmentOperation{})
	} else if recreate {
		ops = append(ops, &deleteCapacityCommitmentOperation{})
		ops = append(ops, &createCapacityCommitmentOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeCapacityCommitmentDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
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

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetCapacityCommitment(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createCapacityCommitmentOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapCapacityCommitment(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeCapacityCommitmentNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeCapacityCommitmentNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeCapacityCommitmentDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffCapacityCommitment(c, newDesired, newState)
	if err != nil {
		return nil, err
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