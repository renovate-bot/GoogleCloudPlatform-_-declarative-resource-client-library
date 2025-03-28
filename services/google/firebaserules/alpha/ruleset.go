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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Ruleset struct {
	Name       *string          `json:"name"`
	Source     *RulesetSource   `json:"source"`
	CreateTime *string          `json:"createTime"`
	Metadata   *RulesetMetadata `json:"metadata"`
	Project    *string          `json:"project"`
}

func (r *Ruleset) String() string {
	return dcl.SprintResource(r)
}

// The enum RulesetSourceLanguageEnum.
type RulesetSourceLanguageEnum string

// RulesetSourceLanguageEnumRef returns a *RulesetSourceLanguageEnum with the value of string s
// If the empty string is provided, nil is returned.
func RulesetSourceLanguageEnumRef(s string) *RulesetSourceLanguageEnum {
	v := RulesetSourceLanguageEnum(s)
	return &v
}

func (v RulesetSourceLanguageEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LANGUAGE_UNSPECIFIED", "FIREBASE_RULES", "EVENT_FLOW_TRIGGERS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "RulesetSourceLanguageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type RulesetSource struct {
	empty    bool                       `json:"-"`
	Files    []RulesetSourceFiles       `json:"files"`
	Language *RulesetSourceLanguageEnum `json:"language"`
}

type jsonRulesetSource RulesetSource

func (r *RulesetSource) UnmarshalJSON(data []byte) error {
	var res jsonRulesetSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRulesetSource
	} else {

		r.Files = res.Files

		r.Language = res.Language

	}
	return nil
}

// This object is used to assert a desired state where this RulesetSource is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRulesetSource *RulesetSource = &RulesetSource{empty: true}

func (r *RulesetSource) Empty() bool {
	return r.empty
}

func (r *RulesetSource) String() string {
	return dcl.SprintResource(r)
}

func (r *RulesetSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RulesetSourceFiles struct {
	empty       bool    `json:"-"`
	Content     *string `json:"content"`
	Name        *string `json:"name"`
	Fingerprint *string `json:"fingerprint"`
}

type jsonRulesetSourceFiles RulesetSourceFiles

func (r *RulesetSourceFiles) UnmarshalJSON(data []byte) error {
	var res jsonRulesetSourceFiles
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRulesetSourceFiles
	} else {

		r.Content = res.Content

		r.Name = res.Name

		r.Fingerprint = res.Fingerprint

	}
	return nil
}

// This object is used to assert a desired state where this RulesetSourceFiles is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRulesetSourceFiles *RulesetSourceFiles = &RulesetSourceFiles{empty: true}

func (r *RulesetSourceFiles) Empty() bool {
	return r.empty
}

func (r *RulesetSourceFiles) String() string {
	return dcl.SprintResource(r)
}

func (r *RulesetSourceFiles) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type RulesetMetadata struct {
	empty    bool     `json:"-"`
	Services []string `json:"services"`
}

type jsonRulesetMetadata RulesetMetadata

func (r *RulesetMetadata) UnmarshalJSON(data []byte) error {
	var res jsonRulesetMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyRulesetMetadata
	} else {

		r.Services = res.Services

	}
	return nil
}

// This object is used to assert a desired state where this RulesetMetadata is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyRulesetMetadata *RulesetMetadata = &RulesetMetadata{empty: true}

func (r *RulesetMetadata) Empty() bool {
	return r.empty
}

func (r *RulesetMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *RulesetMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Ruleset) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "firebaserules",
		Type:    "Ruleset",
		Version: "alpha",
	}
}

func (r *Ruleset) ID() (string, error) {
	if err := extractRulesetFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"source":      dcl.ValueOrEmptyString(nr.Source),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"metadata":    dcl.ValueOrEmptyString(nr.Metadata),
		"project":     dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/rulesets/{{name}}", params), nil
}

const RulesetMaxPage = -1

type RulesetList struct {
	Items []*Ruleset

	nextToken string

	pageSize int32

	resource *Ruleset
}

func (l *RulesetList) HasNext() bool {
	return l.nextToken != ""
}

func (l *RulesetList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listRuleset(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListRuleset(ctx context.Context, project string) (*RulesetList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListRulesetWithMaxResults(ctx, project, RulesetMaxPage)

}

func (c *Client) ListRulesetWithMaxResults(ctx context.Context, project string, pageSize int32) (*RulesetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Ruleset{
		Project: &project,
	}
	items, token, err := c.listRuleset(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &RulesetList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetRuleset(ctx context.Context, r *Ruleset) (*Ruleset, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractRulesetFields(r)

	b, err := c.getRulesetRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 400) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalRuleset(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeRulesetNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractRulesetFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteRuleset(ctx context.Context, r *Ruleset) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Ruleset resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Ruleset...")
	deleteOp := deleteRulesetOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllRuleset deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllRuleset(ctx context.Context, project string, filter func(*Ruleset) bool) error {
	listObj, err := c.ListRuleset(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllRuleset(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllRuleset(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyRuleset(ctx context.Context, rawDesired *Ruleset, opts ...dcl.ApplyOption) (*Ruleset, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Ruleset
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyRulesetHelper(c, ctx, rawDesired, opts...)
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

func applyRulesetHelper(c *Client, ctx context.Context, rawDesired *Ruleset, opts ...dcl.ApplyOption) (*Ruleset, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyRuleset...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractRulesetFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.rulesetDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToRulesetDiffs(c.Config, fieldDiffs, opts)
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
	var ops []rulesetApiOperation
	if create {
		ops = append(ops, &createRulesetOperation{})
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
	return applyRulesetDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyRulesetDiff(c *Client, ctx context.Context, desired *Ruleset, rawDesired *Ruleset, ops []rulesetApiOperation, opts ...dcl.ApplyOption) (*Ruleset, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetRuleset(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createRulesetOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapRuleset(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeRulesetNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeRulesetNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeRulesetDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractRulesetFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractRulesetFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffRuleset(c, newDesired, newState)
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
