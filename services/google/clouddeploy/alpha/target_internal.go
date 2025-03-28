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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Target) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Gke", "AnthosCluster", "Run", "MultiTarget", "CustomTarget"}, r.Gke, r.AnthosCluster, r.Run, r.MultiTarget, r.CustomTarget); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Gke) {
		if err := r.Gke.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AnthosCluster) {
		if err := r.AnthosCluster.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Run) {
		if err := r.Run.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MultiTarget) {
		if err := r.MultiTarget.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CustomTarget) {
		if err := r.CustomTarget.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TargetGke) validate() error {
	return nil
}
func (r *TargetAnthosCluster) validate() error {
	return nil
}
func (r *TargetExecutionConfigs) validate() error {
	if err := dcl.Required(r, "usages"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string(nil)); err != nil {
		return err
	}
	return nil
}
func (r *TargetRun) validate() error {
	if err := dcl.Required(r, "location"); err != nil {
		return err
	}
	return nil
}
func (r *TargetMultiTarget) validate() error {
	if err := dcl.Required(r, "targetIds"); err != nil {
		return err
	}
	return nil
}
func (r *TargetCustomTarget) validate() error {
	if err := dcl.Required(r, "customTargetType"); err != nil {
		return err
	}
	return nil
}
func (r *TargetAssociatedEntities) validate() error {
	return nil
}
func (r *TargetAssociatedEntitiesGkeClusters) validate() error {
	return nil
}
func (r *TargetAssociatedEntitiesAnthosClusters) validate() error {
	return nil
}
func (r *Target) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://clouddeploy.googleapis.com/v1/", params)
}

func (r *Target) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/targets/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Target) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/targets", nr.basePath(), userBasePath, params), nil

}

func (r *Target) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/targets?targetId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Target) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/targets/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Target) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *Target) SetPolicyVerb() string {
	return ""
}

func (r *Target) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *Target) IAMPolicyVersion() int {
	return 3
}

// targetApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type targetApiOperation interface {
	do(context.Context, *Target, *Client) error
}

// newUpdateTargetUpdateTargetRequest creates a request for an
// Target resource's UpdateTarget update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateTargetUpdateTargetRequest(ctx context.Context, f *Target, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		req["annotations"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.RequireApproval; !dcl.IsEmptyValueIndirect(v) {
		req["requireApproval"] = v
	}
	if v, err := expandTargetGke(c, f.Gke, res); err != nil {
		return nil, fmt.Errorf("error expanding Gke into gke: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["gke"] = v
	}
	if v, err := expandTargetAnthosCluster(c, f.AnthosCluster, res); err != nil {
		return nil, fmt.Errorf("error expanding AnthosCluster into anthosCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["anthosCluster"] = v
	}
	if v, err := expandTargetExecutionConfigsSlice(c, f.ExecutionConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding ExecutionConfigs into executionConfigs: %w", err)
	} else if v != nil {
		req["executionConfigs"] = v
	}
	if v, err := expandTargetRun(c, f.Run, res); err != nil {
		return nil, fmt.Errorf("error expanding Run into run: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["run"] = v
	}
	if v, err := expandTargetMultiTarget(c, f.MultiTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding MultiTarget into multiTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["multiTarget"] = v
	}
	if v := f.DeployParameters; !dcl.IsEmptyValueIndirect(v) {
		req["deployParameters"] = v
	}
	if v, err := expandTargetCustomTarget(c, f.CustomTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding CustomTarget into customTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["customTarget"] = v
	}
	if v, err := expandTargetAssociatedEntitiesMap(c, f.AssociatedEntities, res); err != nil {
		return nil, fmt.Errorf("error expanding AssociatedEntities into associatedEntities: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["associatedEntities"] = v
	}
	b, err := c.getTargetRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	req["name"] = fmt.Sprintf("projects/%s/locations/%s/targets/%s", *f.Project, *f.Location, *f.Name)

	return req, nil
}

// marshalUpdateTargetUpdateTargetRequest converts the update into
// the final JSON request body.
func marshalUpdateTargetUpdateTargetRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateTargetUpdateTargetOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateTargetUpdateTargetOperation) do(ctx context.Context, r *Target, c *Client) error {
	_, err := c.GetTarget(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateTarget")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateTargetUpdateTargetRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateTargetUpdateTargetRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listTargetRaw(ctx context.Context, r *Target, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != TargetMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listTargetOperation struct {
	Targets []map[string]interface{} `json:"targets"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listTarget(ctx context.Context, r *Target, pageToken string, pageSize int32) ([]*Target, string, error) {
	b, err := c.listTargetRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listTargetOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Target
	for _, v := range m.Targets {
		res, err := unmarshalMapTarget(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllTarget(ctx context.Context, f func(*Target) bool, resources []*Target) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteTarget(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteTargetOperation struct{}

func (op *deleteTargetOperation) do(ctx context.Context, r *Target, c *Client) error {
	r, err := c.GetTarget(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Target not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetTarget checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetTarget(ctx, r)
		if dcl.IsNotFound(err) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createTargetOperation struct {
	response map[string]interface{}
}

func (op *createTargetOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createTargetOperation) do(ctx context.Context, r *Target, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetTarget(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getTargetRaw(ctx context.Context, r *Target) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) targetDiffsForRawDesired(ctx context.Context, rawDesired *Target, opts ...dcl.ApplyOption) (initial, desired *Target, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Target
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Target); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Target, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetTarget(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Target resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Target resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Target resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeTargetDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Target: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Target: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractTargetFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeTargetInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Target: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeTargetDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Target: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffTarget(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeTargetInitialState(rawInitial, rawDesired *Target) (*Target, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.Gke) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.AnthosCluster, rawInitial.Run, rawInitial.MultiTarget, rawInitial.CustomTarget) {
			rawInitial.Gke = EmptyTargetGke
		}
	}

	if !dcl.IsZeroValue(rawInitial.AnthosCluster) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Gke, rawInitial.Run, rawInitial.MultiTarget, rawInitial.CustomTarget) {
			rawInitial.AnthosCluster = EmptyTargetAnthosCluster
		}
	}

	if !dcl.IsZeroValue(rawInitial.Run) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Gke, rawInitial.AnthosCluster, rawInitial.MultiTarget, rawInitial.CustomTarget) {
			rawInitial.Run = EmptyTargetRun
		}
	}

	if !dcl.IsZeroValue(rawInitial.MultiTarget) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Gke, rawInitial.AnthosCluster, rawInitial.Run, rawInitial.CustomTarget) {
			rawInitial.MultiTarget = EmptyTargetMultiTarget
		}
	}

	if !dcl.IsZeroValue(rawInitial.CustomTarget) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Gke, rawInitial.AnthosCluster, rawInitial.Run, rawInitial.MultiTarget) {
			rawInitial.CustomTarget = EmptyTargetCustomTarget
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeTargetDesiredState(rawDesired, rawInitial *Target, opts ...dcl.ApplyOption) (*Target, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Gke = canonicalizeTargetGke(rawDesired.Gke, nil, opts...)
		rawDesired.AnthosCluster = canonicalizeTargetAnthosCluster(rawDesired.AnthosCluster, nil, opts...)
		rawDesired.Run = canonicalizeTargetRun(rawDesired.Run, nil, opts...)
		rawDesired.MultiTarget = canonicalizeTargetMultiTarget(rawDesired.MultiTarget, nil, opts...)
		rawDesired.CustomTarget = canonicalizeTargetCustomTarget(rawDesired.CustomTarget, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Target{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Annotations) || (dcl.IsEmptyValueIndirect(rawDesired.Annotations) && dcl.IsEmptyValueIndirect(rawInitial.Annotations)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.BoolCanonicalize(rawDesired.RequireApproval, rawInitial.RequireApproval) {
		canonicalDesired.RequireApproval = rawInitial.RequireApproval
	} else {
		canonicalDesired.RequireApproval = rawDesired.RequireApproval
	}
	canonicalDesired.Gke = canonicalizeTargetGke(rawDesired.Gke, rawInitial.Gke, opts...)
	canonicalDesired.AnthosCluster = canonicalizeTargetAnthosCluster(rawDesired.AnthosCluster, rawInitial.AnthosCluster, opts...)
	canonicalDesired.ExecutionConfigs = canonicalizeTargetExecutionConfigsSlice(rawDesired.ExecutionConfigs, rawInitial.ExecutionConfigs, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	canonicalDesired.Run = canonicalizeTargetRun(rawDesired.Run, rawInitial.Run, opts...)
	canonicalDesired.MultiTarget = canonicalizeTargetMultiTarget(rawDesired.MultiTarget, rawInitial.MultiTarget, opts...)
	if dcl.IsZeroValue(rawDesired.DeployParameters) || (dcl.IsEmptyValueIndirect(rawDesired.DeployParameters) && dcl.IsEmptyValueIndirect(rawInitial.DeployParameters)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.DeployParameters = rawInitial.DeployParameters
	} else {
		canonicalDesired.DeployParameters = rawDesired.DeployParameters
	}
	canonicalDesired.CustomTarget = canonicalizeTargetCustomTarget(rawDesired.CustomTarget, rawInitial.CustomTarget, opts...)
	if dcl.IsZeroValue(rawDesired.AssociatedEntities) || (dcl.IsEmptyValueIndirect(rawDesired.AssociatedEntities) && dcl.IsEmptyValueIndirect(rawInitial.AssociatedEntities)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.AssociatedEntities = rawInitial.AssociatedEntities
	} else {
		canonicalDesired.AssociatedEntities = rawDesired.AssociatedEntities
	}

	if canonicalDesired.Gke != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.AnthosCluster, rawDesired.Run, rawDesired.MultiTarget, rawDesired.CustomTarget) {
			canonicalDesired.Gke = EmptyTargetGke
		}
	}

	if canonicalDesired.AnthosCluster != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Gke, rawDesired.Run, rawDesired.MultiTarget, rawDesired.CustomTarget) {
			canonicalDesired.AnthosCluster = EmptyTargetAnthosCluster
		}
	}

	if canonicalDesired.Run != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Gke, rawDesired.AnthosCluster, rawDesired.MultiTarget, rawDesired.CustomTarget) {
			canonicalDesired.Run = EmptyTargetRun
		}
	}

	if canonicalDesired.MultiTarget != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Gke, rawDesired.AnthosCluster, rawDesired.Run, rawDesired.CustomTarget) {
			canonicalDesired.MultiTarget = EmptyTargetMultiTarget
		}
	}

	if canonicalDesired.CustomTarget != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Gke, rawDesired.AnthosCluster, rawDesired.Run, rawDesired.MultiTarget) {
			canonicalDesired.CustomTarget = EmptyTargetCustomTarget
		}
	}

	return canonicalDesired, nil
}

func canonicalizeTargetNewState(c *Client, rawNew, rawDesired *Target) (*Target, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.TargetId) && dcl.IsEmptyValueIndirect(rawDesired.TargetId) {
		rawNew.TargetId = rawDesired.TargetId
	} else {
		if dcl.StringCanonicalize(rawDesired.TargetId, rawNew.TargetId) {
			rawNew.TargetId = rawDesired.TargetId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Uid) && dcl.IsEmptyValueIndirect(rawDesired.Uid) {
		rawNew.Uid = rawDesired.Uid
	} else {
		if dcl.StringCanonicalize(rawDesired.Uid, rawNew.Uid) {
			rawNew.Uid = rawDesired.Uid
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.RequireApproval) && dcl.IsEmptyValueIndirect(rawDesired.RequireApproval) {
		rawNew.RequireApproval = rawDesired.RequireApproval
	} else {
		if dcl.BoolCanonicalize(rawDesired.RequireApproval, rawNew.RequireApproval) {
			rawNew.RequireApproval = rawDesired.RequireApproval
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Gke) && dcl.IsEmptyValueIndirect(rawDesired.Gke) {
		rawNew.Gke = rawDesired.Gke
	} else {
		rawNew.Gke = canonicalizeNewTargetGke(c, rawDesired.Gke, rawNew.Gke)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AnthosCluster) && dcl.IsEmptyValueIndirect(rawDesired.AnthosCluster) {
		rawNew.AnthosCluster = rawDesired.AnthosCluster
	} else {
		rawNew.AnthosCluster = canonicalizeNewTargetAnthosCluster(c, rawDesired.AnthosCluster, rawNew.AnthosCluster)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExecutionConfigs) && dcl.IsEmptyValueIndirect(rawDesired.ExecutionConfigs) {
		rawNew.ExecutionConfigs = rawDesired.ExecutionConfigs
	} else {
		rawNew.ExecutionConfigs = canonicalizeNewTargetExecutionConfigsSlice(c, rawDesired.ExecutionConfigs, rawNew.ExecutionConfigs)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	if dcl.IsEmptyValueIndirect(rawNew.Run) && dcl.IsEmptyValueIndirect(rawDesired.Run) {
		rawNew.Run = rawDesired.Run
	} else {
		rawNew.Run = canonicalizeNewTargetRun(c, rawDesired.Run, rawNew.Run)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MultiTarget) && dcl.IsEmptyValueIndirect(rawDesired.MultiTarget) {
		rawNew.MultiTarget = rawDesired.MultiTarget
	} else {
		rawNew.MultiTarget = canonicalizeNewTargetMultiTarget(c, rawDesired.MultiTarget, rawNew.MultiTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeployParameters) && dcl.IsEmptyValueIndirect(rawDesired.DeployParameters) {
		rawNew.DeployParameters = rawDesired.DeployParameters
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CustomTarget) && dcl.IsEmptyValueIndirect(rawDesired.CustomTarget) {
		rawNew.CustomTarget = rawDesired.CustomTarget
	} else {
		rawNew.CustomTarget = canonicalizeNewTargetCustomTarget(c, rawDesired.CustomTarget, rawNew.CustomTarget)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AssociatedEntities) && dcl.IsEmptyValueIndirect(rawDesired.AssociatedEntities) {
		rawNew.AssociatedEntities = rawDesired.AssociatedEntities
	} else {
	}

	return rawNew, nil
}

func canonicalizeTargetGke(des, initial *TargetGke, opts ...dcl.ApplyOption) *TargetGke {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetGke{}

	if dcl.IsZeroValue(des.Cluster) || (dcl.IsEmptyValueIndirect(des.Cluster) && dcl.IsEmptyValueIndirect(initial.Cluster)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Cluster = initial.Cluster
	} else {
		cDes.Cluster = des.Cluster
	}
	if dcl.BoolCanonicalize(des.InternalIP, initial.InternalIP) || dcl.IsZeroValue(des.InternalIP) {
		cDes.InternalIP = initial.InternalIP
	} else {
		cDes.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ProxyUrl, initial.ProxyUrl) || dcl.IsZeroValue(des.ProxyUrl) {
		cDes.ProxyUrl = initial.ProxyUrl
	} else {
		cDes.ProxyUrl = des.ProxyUrl
	}
	if dcl.BoolCanonicalize(des.DnsEndpoint, initial.DnsEndpoint) || dcl.IsZeroValue(des.DnsEndpoint) {
		cDes.DnsEndpoint = initial.DnsEndpoint
	} else {
		cDes.DnsEndpoint = des.DnsEndpoint
	}

	return cDes
}

func canonicalizeTargetGkeSlice(des, initial []TargetGke, opts ...dcl.ApplyOption) []TargetGke {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetGke, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetGke(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetGke, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetGke(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetGke(c *Client, des, nw *TargetGke) *TargetGke {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetGke while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.InternalIP, nw.InternalIP) {
		nw.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ProxyUrl, nw.ProxyUrl) {
		nw.ProxyUrl = des.ProxyUrl
	}
	if dcl.BoolCanonicalize(des.DnsEndpoint, nw.DnsEndpoint) {
		nw.DnsEndpoint = des.DnsEndpoint
	}

	return nw
}

func canonicalizeNewTargetGkeSet(c *Client, des, nw []TargetGke) []TargetGke {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetGke
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetGkeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetGke(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetGkeSlice(c *Client, des, nw []TargetGke) []TargetGke {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetGke
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetGke(c, &d, &n))
	}

	return items
}

func canonicalizeTargetAnthosCluster(des, initial *TargetAnthosCluster, opts ...dcl.ApplyOption) *TargetAnthosCluster {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetAnthosCluster{}

	if dcl.IsZeroValue(des.Membership) || (dcl.IsEmptyValueIndirect(des.Membership) && dcl.IsEmptyValueIndirect(initial.Membership)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Membership = initial.Membership
	} else {
		cDes.Membership = des.Membership
	}

	return cDes
}

func canonicalizeTargetAnthosClusterSlice(des, initial []TargetAnthosCluster, opts ...dcl.ApplyOption) []TargetAnthosCluster {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetAnthosCluster, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetAnthosCluster(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetAnthosCluster, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetAnthosCluster(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetAnthosCluster(c *Client, des, nw *TargetAnthosCluster) *TargetAnthosCluster {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetAnthosCluster while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTargetAnthosClusterSet(c *Client, des, nw []TargetAnthosCluster) []TargetAnthosCluster {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetAnthosCluster
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetAnthosClusterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetAnthosCluster(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetAnthosClusterSlice(c *Client, des, nw []TargetAnthosCluster) []TargetAnthosCluster {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetAnthosCluster
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetAnthosCluster(c, &d, &n))
	}

	return items
}

func canonicalizeTargetExecutionConfigs(des, initial *TargetExecutionConfigs, opts ...dcl.ApplyOption) *TargetExecutionConfigs {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetExecutionConfigs{}

	if dcl.IsZeroValue(des.Usages) || (dcl.IsEmptyValueIndirect(des.Usages) && dcl.IsEmptyValueIndirect(initial.Usages)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Usages = initial.Usages
	} else {
		cDes.Usages = des.Usages
	}
	if dcl.IsZeroValue(des.WorkerPool) || (dcl.IsEmptyValueIndirect(des.WorkerPool) && dcl.IsEmptyValueIndirect(initial.WorkerPool)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WorkerPool = initial.WorkerPool
	} else {
		cDes.WorkerPool = des.WorkerPool
	}
	if dcl.StringCanonicalize(des.ServiceAccount, initial.ServiceAccount) || dcl.IsZeroValue(des.ServiceAccount) {
		cDes.ServiceAccount = initial.ServiceAccount
	} else {
		cDes.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringCanonicalize(des.ArtifactStorage, initial.ArtifactStorage) || dcl.IsZeroValue(des.ArtifactStorage) {
		cDes.ArtifactStorage = initial.ArtifactStorage
	} else {
		cDes.ArtifactStorage = des.ArtifactStorage
	}
	if dcl.StringCanonicalize(des.ExecutionTimeout, initial.ExecutionTimeout) || dcl.IsZeroValue(des.ExecutionTimeout) {
		cDes.ExecutionTimeout = initial.ExecutionTimeout
	} else {
		cDes.ExecutionTimeout = des.ExecutionTimeout
	}
	if dcl.BoolCanonicalize(des.Verbose, initial.Verbose) || dcl.IsZeroValue(des.Verbose) {
		cDes.Verbose = initial.Verbose
	} else {
		cDes.Verbose = des.Verbose
	}

	return cDes
}

func canonicalizeTargetExecutionConfigsSlice(des, initial []TargetExecutionConfigs, opts ...dcl.ApplyOption) []TargetExecutionConfigs {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetExecutionConfigs, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetExecutionConfigs(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetExecutionConfigs, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetExecutionConfigs(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetExecutionConfigs(c *Client, des, nw *TargetExecutionConfigs) *TargetExecutionConfigs {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetExecutionConfigs while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ServiceAccount, nw.ServiceAccount) {
		nw.ServiceAccount = des.ServiceAccount
	}
	if dcl.StringCanonicalize(des.ArtifactStorage, nw.ArtifactStorage) {
		nw.ArtifactStorage = des.ArtifactStorage
	}
	if dcl.StringCanonicalize(des.ExecutionTimeout, nw.ExecutionTimeout) {
		nw.ExecutionTimeout = des.ExecutionTimeout
	}
	if dcl.BoolCanonicalize(des.Verbose, nw.Verbose) {
		nw.Verbose = des.Verbose
	}

	return nw
}

func canonicalizeNewTargetExecutionConfigsSet(c *Client, des, nw []TargetExecutionConfigs) []TargetExecutionConfigs {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetExecutionConfigs
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetExecutionConfigsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetExecutionConfigs(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetExecutionConfigsSlice(c *Client, des, nw []TargetExecutionConfigs) []TargetExecutionConfigs {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetExecutionConfigs
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetExecutionConfigs(c, &d, &n))
	}

	return items
}

func canonicalizeTargetRun(des, initial *TargetRun, opts ...dcl.ApplyOption) *TargetRun {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetRun{}

	if dcl.StringCanonicalize(des.Location, initial.Location) || dcl.IsZeroValue(des.Location) {
		cDes.Location = initial.Location
	} else {
		cDes.Location = des.Location
	}

	return cDes
}

func canonicalizeTargetRunSlice(des, initial []TargetRun, opts ...dcl.ApplyOption) []TargetRun {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetRun, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetRun(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetRun, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetRun(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetRun(c *Client, des, nw *TargetRun) *TargetRun {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetRun while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Location, nw.Location) {
		nw.Location = des.Location
	}

	return nw
}

func canonicalizeNewTargetRunSet(c *Client, des, nw []TargetRun) []TargetRun {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetRun
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetRunNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetRun(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetRunSlice(c *Client, des, nw []TargetRun) []TargetRun {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetRun
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetRun(c, &d, &n))
	}

	return items
}

func canonicalizeTargetMultiTarget(des, initial *TargetMultiTarget, opts ...dcl.ApplyOption) *TargetMultiTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetMultiTarget{}

	if dcl.StringArrayCanonicalize(des.TargetIds, initial.TargetIds) {
		cDes.TargetIds = initial.TargetIds
	} else {
		cDes.TargetIds = des.TargetIds
	}

	return cDes
}

func canonicalizeTargetMultiTargetSlice(des, initial []TargetMultiTarget, opts ...dcl.ApplyOption) []TargetMultiTarget {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetMultiTarget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetMultiTarget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetMultiTarget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetMultiTarget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetMultiTarget(c *Client, des, nw *TargetMultiTarget) *TargetMultiTarget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetMultiTarget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.TargetIds, nw.TargetIds) {
		nw.TargetIds = des.TargetIds
	}

	return nw
}

func canonicalizeNewTargetMultiTargetSet(c *Client, des, nw []TargetMultiTarget) []TargetMultiTarget {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetMultiTarget
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetMultiTargetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetMultiTarget(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetMultiTargetSlice(c *Client, des, nw []TargetMultiTarget) []TargetMultiTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetMultiTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetMultiTarget(c, &d, &n))
	}

	return items
}

func canonicalizeTargetCustomTarget(des, initial *TargetCustomTarget, opts ...dcl.ApplyOption) *TargetCustomTarget {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetCustomTarget{}

	if dcl.IsZeroValue(des.CustomTargetType) || (dcl.IsEmptyValueIndirect(des.CustomTargetType) && dcl.IsEmptyValueIndirect(initial.CustomTargetType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CustomTargetType = initial.CustomTargetType
	} else {
		cDes.CustomTargetType = des.CustomTargetType
	}

	return cDes
}

func canonicalizeTargetCustomTargetSlice(des, initial []TargetCustomTarget, opts ...dcl.ApplyOption) []TargetCustomTarget {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetCustomTarget, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetCustomTarget(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetCustomTarget, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetCustomTarget(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetCustomTarget(c *Client, des, nw *TargetCustomTarget) *TargetCustomTarget {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetCustomTarget while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTargetCustomTargetSet(c *Client, des, nw []TargetCustomTarget) []TargetCustomTarget {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetCustomTarget
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetCustomTargetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetCustomTarget(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetCustomTargetSlice(c *Client, des, nw []TargetCustomTarget) []TargetCustomTarget {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetCustomTarget
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetCustomTarget(c, &d, &n))
	}

	return items
}

func canonicalizeTargetAssociatedEntities(des, initial *TargetAssociatedEntities, opts ...dcl.ApplyOption) *TargetAssociatedEntities {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetAssociatedEntities{}

	cDes.GkeClusters = canonicalizeTargetAssociatedEntitiesGkeClustersSlice(des.GkeClusters, initial.GkeClusters, opts...)
	cDes.AnthosClusters = canonicalizeTargetAssociatedEntitiesAnthosClustersSlice(des.AnthosClusters, initial.AnthosClusters, opts...)

	return cDes
}

func canonicalizeTargetAssociatedEntitiesSlice(des, initial []TargetAssociatedEntities, opts ...dcl.ApplyOption) []TargetAssociatedEntities {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetAssociatedEntities, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetAssociatedEntities(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetAssociatedEntities, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetAssociatedEntities(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetAssociatedEntities(c *Client, des, nw *TargetAssociatedEntities) *TargetAssociatedEntities {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetAssociatedEntities while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.GkeClusters = canonicalizeNewTargetAssociatedEntitiesGkeClustersSlice(c, des.GkeClusters, nw.GkeClusters)
	nw.AnthosClusters = canonicalizeNewTargetAssociatedEntitiesAnthosClustersSlice(c, des.AnthosClusters, nw.AnthosClusters)

	return nw
}

func canonicalizeNewTargetAssociatedEntitiesSet(c *Client, des, nw []TargetAssociatedEntities) []TargetAssociatedEntities {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetAssociatedEntities
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetAssociatedEntitiesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetAssociatedEntities(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetAssociatedEntitiesSlice(c *Client, des, nw []TargetAssociatedEntities) []TargetAssociatedEntities {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetAssociatedEntities
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetAssociatedEntities(c, &d, &n))
	}

	return items
}

func canonicalizeTargetAssociatedEntitiesGkeClusters(des, initial *TargetAssociatedEntitiesGkeClusters, opts ...dcl.ApplyOption) *TargetAssociatedEntitiesGkeClusters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetAssociatedEntitiesGkeClusters{}

	if dcl.IsZeroValue(des.Cluster) || (dcl.IsEmptyValueIndirect(des.Cluster) && dcl.IsEmptyValueIndirect(initial.Cluster)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Cluster = initial.Cluster
	} else {
		cDes.Cluster = des.Cluster
	}
	if dcl.BoolCanonicalize(des.InternalIP, initial.InternalIP) || dcl.IsZeroValue(des.InternalIP) {
		cDes.InternalIP = initial.InternalIP
	} else {
		cDes.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ProxyUrl, initial.ProxyUrl) || dcl.IsZeroValue(des.ProxyUrl) {
		cDes.ProxyUrl = initial.ProxyUrl
	} else {
		cDes.ProxyUrl = des.ProxyUrl
	}

	return cDes
}

func canonicalizeTargetAssociatedEntitiesGkeClustersSlice(des, initial []TargetAssociatedEntitiesGkeClusters, opts ...dcl.ApplyOption) []TargetAssociatedEntitiesGkeClusters {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetAssociatedEntitiesGkeClusters, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetAssociatedEntitiesGkeClusters(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetAssociatedEntitiesGkeClusters, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetAssociatedEntitiesGkeClusters(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetAssociatedEntitiesGkeClusters(c *Client, des, nw *TargetAssociatedEntitiesGkeClusters) *TargetAssociatedEntitiesGkeClusters {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetAssociatedEntitiesGkeClusters while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.InternalIP, nw.InternalIP) {
		nw.InternalIP = des.InternalIP
	}
	if dcl.StringCanonicalize(des.ProxyUrl, nw.ProxyUrl) {
		nw.ProxyUrl = des.ProxyUrl
	}

	return nw
}

func canonicalizeNewTargetAssociatedEntitiesGkeClustersSet(c *Client, des, nw []TargetAssociatedEntitiesGkeClusters) []TargetAssociatedEntitiesGkeClusters {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetAssociatedEntitiesGkeClusters
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetAssociatedEntitiesGkeClustersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetAssociatedEntitiesGkeClusters(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetAssociatedEntitiesGkeClustersSlice(c *Client, des, nw []TargetAssociatedEntitiesGkeClusters) []TargetAssociatedEntitiesGkeClusters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetAssociatedEntitiesGkeClusters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetAssociatedEntitiesGkeClusters(c, &d, &n))
	}

	return items
}

func canonicalizeTargetAssociatedEntitiesAnthosClusters(des, initial *TargetAssociatedEntitiesAnthosClusters, opts ...dcl.ApplyOption) *TargetAssociatedEntitiesAnthosClusters {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &TargetAssociatedEntitiesAnthosClusters{}

	if dcl.IsZeroValue(des.Membership) || (dcl.IsEmptyValueIndirect(des.Membership) && dcl.IsEmptyValueIndirect(initial.Membership)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Membership = initial.Membership
	} else {
		cDes.Membership = des.Membership
	}

	return cDes
}

func canonicalizeTargetAssociatedEntitiesAnthosClustersSlice(des, initial []TargetAssociatedEntitiesAnthosClusters, opts ...dcl.ApplyOption) []TargetAssociatedEntitiesAnthosClusters {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]TargetAssociatedEntitiesAnthosClusters, 0, len(des))
		for _, d := range des {
			cd := canonicalizeTargetAssociatedEntitiesAnthosClusters(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]TargetAssociatedEntitiesAnthosClusters, 0, len(des))
	for i, d := range des {
		cd := canonicalizeTargetAssociatedEntitiesAnthosClusters(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewTargetAssociatedEntitiesAnthosClusters(c *Client, des, nw *TargetAssociatedEntitiesAnthosClusters) *TargetAssociatedEntitiesAnthosClusters {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for TargetAssociatedEntitiesAnthosClusters while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewTargetAssociatedEntitiesAnthosClustersSet(c *Client, des, nw []TargetAssociatedEntitiesAnthosClusters) []TargetAssociatedEntitiesAnthosClusters {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []TargetAssociatedEntitiesAnthosClusters
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareTargetAssociatedEntitiesAnthosClustersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewTargetAssociatedEntitiesAnthosClusters(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewTargetAssociatedEntitiesAnthosClustersSlice(c *Client, des, nw []TargetAssociatedEntitiesAnthosClusters) []TargetAssociatedEntitiesAnthosClusters {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []TargetAssociatedEntitiesAnthosClusters
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewTargetAssociatedEntitiesAnthosClusters(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffTarget(c *Client, desired, actual *Target, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TargetId, actual.TargetId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TargetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uid, actual.Uid, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireApproval, actual.RequireApproval, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("RequireApproval")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gke, actual.Gke, dcl.DiffInfo{ObjectFunction: compareTargetGkeNewStyle, EmptyObject: EmptyTargetGke, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Gke")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnthosCluster, actual.AnthosCluster, dcl.DiffInfo{ObjectFunction: compareTargetAnthosClusterNewStyle, EmptyObject: EmptyTargetAnthosCluster, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("AnthosCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionConfigs, actual.ExecutionConfigs, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareTargetExecutionConfigsNewStyle, EmptyObject: EmptyTargetExecutionConfigs, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("ExecutionConfigs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Run, actual.Run, dcl.DiffInfo{ObjectFunction: compareTargetRunNewStyle, EmptyObject: EmptyTargetRun, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Run")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MultiTarget, actual.MultiTarget, dcl.DiffInfo{ObjectFunction: compareTargetMultiTargetNewStyle, EmptyObject: EmptyTargetMultiTarget, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("MultiTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeployParameters, actual.DeployParameters, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("DeployParameters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomTarget, actual.CustomTarget, dcl.DiffInfo{ObjectFunction: compareTargetCustomTargetNewStyle, EmptyObject: EmptyTargetCustomTarget, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("CustomTarget")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AssociatedEntities, actual.AssociatedEntities, dcl.DiffInfo{ObjectFunction: compareTargetAssociatedEntitiesNewStyle, EmptyObject: EmptyTargetAssociatedEntities, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("AssociatedEntities")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareTargetGkeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetGke)
	if !ok {
		desiredNotPointer, ok := d.(TargetGke)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetGke or *TargetGke", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetGke)
	if !ok {
		actualNotPointer, ok := a.(TargetGke)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetGke", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Cluster, actual.Cluster, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Cluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalIP, actual.InternalIP, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("InternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyUrl, actual.ProxyUrl, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("ProxyUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DnsEndpoint, actual.DnsEndpoint, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("DnsEndpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetAnthosClusterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetAnthosCluster)
	if !ok {
		desiredNotPointer, ok := d.(TargetAnthosCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAnthosCluster or *TargetAnthosCluster", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetAnthosCluster)
	if !ok {
		actualNotPointer, ok := a.(TargetAnthosCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAnthosCluster", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Membership, actual.Membership, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Membership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetExecutionConfigsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetExecutionConfigs)
	if !ok {
		desiredNotPointer, ok := d.(TargetExecutionConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetExecutionConfigs or *TargetExecutionConfigs", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetExecutionConfigs)
	if !ok {
		actualNotPointer, ok := a.(TargetExecutionConfigs)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetExecutionConfigs", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Usages, actual.Usages, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Usages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkerPool, actual.WorkerPool, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("WorkerPool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccount, actual.ServiceAccount, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("ServiceAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ArtifactStorage, actual.ArtifactStorage, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("ArtifactStorage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExecutionTimeout, actual.ExecutionTimeout, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("ExecutionTimeout")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Verbose, actual.Verbose, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Verbose")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetRunNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetRun)
	if !ok {
		desiredNotPointer, ok := d.(TargetRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetRun or *TargetRun", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetRun)
	if !ok {
		actualNotPointer, ok := a.(TargetRun)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetRun", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetMultiTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetMultiTarget)
	if !ok {
		desiredNotPointer, ok := d.(TargetMultiTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetMultiTarget or *TargetMultiTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetMultiTarget)
	if !ok {
		actualNotPointer, ok := a.(TargetMultiTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetMultiTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TargetIds, actual.TargetIds, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("TargetIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetCustomTargetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetCustomTarget)
	if !ok {
		desiredNotPointer, ok := d.(TargetCustomTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetCustomTarget or *TargetCustomTarget", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetCustomTarget)
	if !ok {
		actualNotPointer, ok := a.(TargetCustomTarget)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetCustomTarget", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CustomTargetType, actual.CustomTargetType, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("CustomTargetType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetAssociatedEntitiesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetAssociatedEntities)
	if !ok {
		desiredNotPointer, ok := d.(TargetAssociatedEntities)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAssociatedEntities or *TargetAssociatedEntities", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetAssociatedEntities)
	if !ok {
		actualNotPointer, ok := a.(TargetAssociatedEntities)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAssociatedEntities", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GkeClusters, actual.GkeClusters, dcl.DiffInfo{ObjectFunction: compareTargetAssociatedEntitiesGkeClustersNewStyle, EmptyObject: EmptyTargetAssociatedEntitiesGkeClusters, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("GkeClusters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AnthosClusters, actual.AnthosClusters, dcl.DiffInfo{ObjectFunction: compareTargetAssociatedEntitiesAnthosClustersNewStyle, EmptyObject: EmptyTargetAssociatedEntitiesAnthosClusters, OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("AnthosClusters")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetAssociatedEntitiesGkeClustersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetAssociatedEntitiesGkeClusters)
	if !ok {
		desiredNotPointer, ok := d.(TargetAssociatedEntitiesGkeClusters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAssociatedEntitiesGkeClusters or *TargetAssociatedEntitiesGkeClusters", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetAssociatedEntitiesGkeClusters)
	if !ok {
		actualNotPointer, ok := a.(TargetAssociatedEntitiesGkeClusters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAssociatedEntitiesGkeClusters", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Cluster, actual.Cluster, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Cluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InternalIP, actual.InternalIP, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("InternalIp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProxyUrl, actual.ProxyUrl, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("ProxyUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareTargetAssociatedEntitiesAnthosClustersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*TargetAssociatedEntitiesAnthosClusters)
	if !ok {
		desiredNotPointer, ok := d.(TargetAssociatedEntitiesAnthosClusters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAssociatedEntitiesAnthosClusters or *TargetAssociatedEntitiesAnthosClusters", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*TargetAssociatedEntitiesAnthosClusters)
	if !ok {
		actualNotPointer, ok := a.(TargetAssociatedEntitiesAnthosClusters)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a TargetAssociatedEntitiesAnthosClusters", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Membership, actual.Membership, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateTargetUpdateTargetOperation")}, fn.AddNest("Membership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Target) urlNormalized() *Target {
	normalized := dcl.Copy(*r).(Target)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.TargetId = dcl.SelfLinkToName(r.TargetId)
	normalized.Uid = dcl.SelfLinkToName(r.Uid)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Target) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateTarget" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/targets/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Target resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Target) marshal(c *Client) ([]byte, error) {
	m, err := expandTarget(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Target: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalTarget decodes JSON responses into the Target resource schema.
func unmarshalTarget(b []byte, c *Client, res *Target) (*Target, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapTarget(m, c, res)
}

func unmarshalMapTarget(m map[string]interface{}, c *Client, res *Target) (*Target, error) {

	flattened := flattenTarget(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandTarget expands Target into a JSON request object.
func expandTarget(c *Client, f *Target) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Annotations; dcl.ValueShouldBeSent(v) {
		m["annotations"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.RequireApproval; dcl.ValueShouldBeSent(v) {
		m["requireApproval"] = v
	}
	if v, err := expandTargetGke(c, f.Gke, res); err != nil {
		return nil, fmt.Errorf("error expanding Gke into gke: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gke"] = v
	}
	if v, err := expandTargetAnthosCluster(c, f.AnthosCluster, res); err != nil {
		return nil, fmt.Errorf("error expanding AnthosCluster into anthosCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["anthosCluster"] = v
	}
	if v, err := expandTargetExecutionConfigsSlice(c, f.ExecutionConfigs, res); err != nil {
		return nil, fmt.Errorf("error expanding ExecutionConfigs into executionConfigs: %w", err)
	} else if v != nil {
		m["executionConfigs"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}
	if v, err := expandTargetRun(c, f.Run, res); err != nil {
		return nil, fmt.Errorf("error expanding Run into run: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["run"] = v
	}
	if v, err := expandTargetMultiTarget(c, f.MultiTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding MultiTarget into multiTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["multiTarget"] = v
	}
	if v := f.DeployParameters; dcl.ValueShouldBeSent(v) {
		m["deployParameters"] = v
	}
	if v, err := expandTargetCustomTarget(c, f.CustomTarget, res); err != nil {
		return nil, fmt.Errorf("error expanding CustomTarget into customTarget: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["customTarget"] = v
	}
	if v, err := expandTargetAssociatedEntitiesMap(c, f.AssociatedEntities, res); err != nil {
		return nil, fmt.Errorf("error expanding AssociatedEntities into associatedEntities: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["associatedEntities"] = v
	}

	return m, nil
}

// flattenTarget flattens Target from a JSON request object into the
// Target type.
func flattenTarget(c *Client, i interface{}, res *Target) *Target {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Target{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.TargetId = dcl.FlattenString(m["targetId"])
	resultRes.Uid = dcl.FlattenString(m["uid"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.RequireApproval = dcl.FlattenBool(m["requireApproval"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Gke = flattenTargetGke(c, m["gke"], res)
	resultRes.AnthosCluster = flattenTargetAnthosCluster(c, m["anthosCluster"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.ExecutionConfigs = flattenTargetExecutionConfigsSlice(c, m["executionConfigs"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Run = flattenTargetRun(c, m["run"], res)
	resultRes.MultiTarget = flattenTargetMultiTarget(c, m["multiTarget"], res)
	resultRes.DeployParameters = dcl.FlattenKeyValuePairs(m["deployParameters"])
	resultRes.CustomTarget = flattenTargetCustomTarget(c, m["customTarget"], res)
	resultRes.AssociatedEntities = flattenTargetAssociatedEntitiesMap(c, m["associatedEntities"], res)

	return resultRes
}

// expandTargetGkeMap expands the contents of TargetGke into a JSON
// request object.
func expandTargetGkeMap(c *Client, f map[string]TargetGke, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetGke(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetGkeSlice expands the contents of TargetGke into a JSON
// request object.
func expandTargetGkeSlice(c *Client, f []TargetGke, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetGke(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetGkeMap flattens the contents of TargetGke from a JSON
// response object.
func flattenTargetGkeMap(c *Client, i interface{}, res *Target) map[string]TargetGke {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetGke{}
	}

	if len(a) == 0 {
		return map[string]TargetGke{}
	}

	items := make(map[string]TargetGke)
	for k, item := range a {
		items[k] = *flattenTargetGke(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetGkeSlice flattens the contents of TargetGke from a JSON
// response object.
func flattenTargetGkeSlice(c *Client, i interface{}, res *Target) []TargetGke {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetGke{}
	}

	if len(a) == 0 {
		return []TargetGke{}
	}

	items := make([]TargetGke, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetGke(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetGke expands an instance of TargetGke into a JSON
// request object.
func expandTargetGke(c *Client, f *TargetGke, res *Target) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Cluster; !dcl.IsEmptyValueIndirect(v) {
		m["cluster"] = v
	}
	if v := f.InternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["internalIp"] = v
	}
	if v := f.ProxyUrl; !dcl.IsEmptyValueIndirect(v) {
		m["proxyUrl"] = v
	}
	if v := f.DnsEndpoint; !dcl.IsEmptyValueIndirect(v) {
		m["dnsEndpoint"] = v
	}

	return m, nil
}

// flattenTargetGke flattens an instance of TargetGke from a JSON
// response object.
func flattenTargetGke(c *Client, i interface{}, res *Target) *TargetGke {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetGke{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetGke
	}
	r.Cluster = dcl.FlattenString(m["cluster"])
	r.InternalIP = dcl.FlattenBool(m["internalIp"])
	r.ProxyUrl = dcl.FlattenString(m["proxyUrl"])
	r.DnsEndpoint = dcl.FlattenBool(m["dnsEndpoint"])

	return r
}

// expandTargetAnthosClusterMap expands the contents of TargetAnthosCluster into a JSON
// request object.
func expandTargetAnthosClusterMap(c *Client, f map[string]TargetAnthosCluster, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetAnthosCluster(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetAnthosClusterSlice expands the contents of TargetAnthosCluster into a JSON
// request object.
func expandTargetAnthosClusterSlice(c *Client, f []TargetAnthosCluster, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetAnthosCluster(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetAnthosClusterMap flattens the contents of TargetAnthosCluster from a JSON
// response object.
func flattenTargetAnthosClusterMap(c *Client, i interface{}, res *Target) map[string]TargetAnthosCluster {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetAnthosCluster{}
	}

	if len(a) == 0 {
		return map[string]TargetAnthosCluster{}
	}

	items := make(map[string]TargetAnthosCluster)
	for k, item := range a {
		items[k] = *flattenTargetAnthosCluster(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetAnthosClusterSlice flattens the contents of TargetAnthosCluster from a JSON
// response object.
func flattenTargetAnthosClusterSlice(c *Client, i interface{}, res *Target) []TargetAnthosCluster {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetAnthosCluster{}
	}

	if len(a) == 0 {
		return []TargetAnthosCluster{}
	}

	items := make([]TargetAnthosCluster, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetAnthosCluster(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetAnthosCluster expands an instance of TargetAnthosCluster into a JSON
// request object.
func expandTargetAnthosCluster(c *Client, f *TargetAnthosCluster, res *Target) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Membership; !dcl.IsEmptyValueIndirect(v) {
		m["membership"] = v
	}

	return m, nil
}

// flattenTargetAnthosCluster flattens an instance of TargetAnthosCluster from a JSON
// response object.
func flattenTargetAnthosCluster(c *Client, i interface{}, res *Target) *TargetAnthosCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetAnthosCluster{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetAnthosCluster
	}
	r.Membership = dcl.FlattenString(m["membership"])

	return r
}

// expandTargetExecutionConfigsMap expands the contents of TargetExecutionConfigs into a JSON
// request object.
func expandTargetExecutionConfigsMap(c *Client, f map[string]TargetExecutionConfigs, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetExecutionConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetExecutionConfigsSlice expands the contents of TargetExecutionConfigs into a JSON
// request object.
func expandTargetExecutionConfigsSlice(c *Client, f []TargetExecutionConfigs, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetExecutionConfigs(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetExecutionConfigsMap flattens the contents of TargetExecutionConfigs from a JSON
// response object.
func flattenTargetExecutionConfigsMap(c *Client, i interface{}, res *Target) map[string]TargetExecutionConfigs {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetExecutionConfigs{}
	}

	if len(a) == 0 {
		return map[string]TargetExecutionConfigs{}
	}

	items := make(map[string]TargetExecutionConfigs)
	for k, item := range a {
		items[k] = *flattenTargetExecutionConfigs(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetExecutionConfigsSlice flattens the contents of TargetExecutionConfigs from a JSON
// response object.
func flattenTargetExecutionConfigsSlice(c *Client, i interface{}, res *Target) []TargetExecutionConfigs {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetExecutionConfigs{}
	}

	if len(a) == 0 {
		return []TargetExecutionConfigs{}
	}

	items := make([]TargetExecutionConfigs, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetExecutionConfigs(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetExecutionConfigs expands an instance of TargetExecutionConfigs into a JSON
// request object.
func expandTargetExecutionConfigs(c *Client, f *TargetExecutionConfigs, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Usages; v != nil {
		m["usages"] = v
	}
	if v := f.WorkerPool; !dcl.IsEmptyValueIndirect(v) {
		m["workerPool"] = v
	}
	if v := f.ServiceAccount; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccount"] = v
	}
	if v := f.ArtifactStorage; !dcl.IsEmptyValueIndirect(v) {
		m["artifactStorage"] = v
	}
	if v := f.ExecutionTimeout; !dcl.IsEmptyValueIndirect(v) {
		m["executionTimeout"] = v
	}
	if v := f.Verbose; !dcl.IsEmptyValueIndirect(v) {
		m["verbose"] = v
	}

	return m, nil
}

// flattenTargetExecutionConfigs flattens an instance of TargetExecutionConfigs from a JSON
// response object.
func flattenTargetExecutionConfigs(c *Client, i interface{}, res *Target) *TargetExecutionConfigs {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetExecutionConfigs{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetExecutionConfigs
	}
	r.Usages = flattenTargetExecutionConfigsUsagesEnumSlice(c, m["usages"], res)
	r.WorkerPool = dcl.FlattenString(m["workerPool"])
	r.ServiceAccount = dcl.FlattenString(m["serviceAccount"])
	r.ArtifactStorage = dcl.FlattenString(m["artifactStorage"])
	r.ExecutionTimeout = dcl.FlattenString(m["executionTimeout"])
	r.Verbose = dcl.FlattenBool(m["verbose"])

	return r
}

// expandTargetRunMap expands the contents of TargetRun into a JSON
// request object.
func expandTargetRunMap(c *Client, f map[string]TargetRun, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetRun(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetRunSlice expands the contents of TargetRun into a JSON
// request object.
func expandTargetRunSlice(c *Client, f []TargetRun, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetRun(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetRunMap flattens the contents of TargetRun from a JSON
// response object.
func flattenTargetRunMap(c *Client, i interface{}, res *Target) map[string]TargetRun {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetRun{}
	}

	if len(a) == 0 {
		return map[string]TargetRun{}
	}

	items := make(map[string]TargetRun)
	for k, item := range a {
		items[k] = *flattenTargetRun(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetRunSlice flattens the contents of TargetRun from a JSON
// response object.
func flattenTargetRunSlice(c *Client, i interface{}, res *Target) []TargetRun {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetRun{}
	}

	if len(a) == 0 {
		return []TargetRun{}
	}

	items := make([]TargetRun, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetRun(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetRun expands an instance of TargetRun into a JSON
// request object.
func expandTargetRun(c *Client, f *TargetRun, res *Target) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Location; !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenTargetRun flattens an instance of TargetRun from a JSON
// response object.
func flattenTargetRun(c *Client, i interface{}, res *Target) *TargetRun {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetRun{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetRun
	}
	r.Location = dcl.FlattenString(m["location"])

	return r
}

// expandTargetMultiTargetMap expands the contents of TargetMultiTarget into a JSON
// request object.
func expandTargetMultiTargetMap(c *Client, f map[string]TargetMultiTarget, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetMultiTarget(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetMultiTargetSlice expands the contents of TargetMultiTarget into a JSON
// request object.
func expandTargetMultiTargetSlice(c *Client, f []TargetMultiTarget, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetMultiTarget(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetMultiTargetMap flattens the contents of TargetMultiTarget from a JSON
// response object.
func flattenTargetMultiTargetMap(c *Client, i interface{}, res *Target) map[string]TargetMultiTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetMultiTarget{}
	}

	if len(a) == 0 {
		return map[string]TargetMultiTarget{}
	}

	items := make(map[string]TargetMultiTarget)
	for k, item := range a {
		items[k] = *flattenTargetMultiTarget(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetMultiTargetSlice flattens the contents of TargetMultiTarget from a JSON
// response object.
func flattenTargetMultiTargetSlice(c *Client, i interface{}, res *Target) []TargetMultiTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetMultiTarget{}
	}

	if len(a) == 0 {
		return []TargetMultiTarget{}
	}

	items := make([]TargetMultiTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetMultiTarget(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetMultiTarget expands an instance of TargetMultiTarget into a JSON
// request object.
func expandTargetMultiTarget(c *Client, f *TargetMultiTarget, res *Target) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TargetIds; v != nil {
		m["targetIds"] = v
	}

	return m, nil
}

// flattenTargetMultiTarget flattens an instance of TargetMultiTarget from a JSON
// response object.
func flattenTargetMultiTarget(c *Client, i interface{}, res *Target) *TargetMultiTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetMultiTarget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetMultiTarget
	}
	r.TargetIds = dcl.FlattenStringSlice(m["targetIds"])

	return r
}

// expandTargetCustomTargetMap expands the contents of TargetCustomTarget into a JSON
// request object.
func expandTargetCustomTargetMap(c *Client, f map[string]TargetCustomTarget, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetCustomTarget(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetCustomTargetSlice expands the contents of TargetCustomTarget into a JSON
// request object.
func expandTargetCustomTargetSlice(c *Client, f []TargetCustomTarget, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetCustomTarget(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetCustomTargetMap flattens the contents of TargetCustomTarget from a JSON
// response object.
func flattenTargetCustomTargetMap(c *Client, i interface{}, res *Target) map[string]TargetCustomTarget {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetCustomTarget{}
	}

	if len(a) == 0 {
		return map[string]TargetCustomTarget{}
	}

	items := make(map[string]TargetCustomTarget)
	for k, item := range a {
		items[k] = *flattenTargetCustomTarget(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetCustomTargetSlice flattens the contents of TargetCustomTarget from a JSON
// response object.
func flattenTargetCustomTargetSlice(c *Client, i interface{}, res *Target) []TargetCustomTarget {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetCustomTarget{}
	}

	if len(a) == 0 {
		return []TargetCustomTarget{}
	}

	items := make([]TargetCustomTarget, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetCustomTarget(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetCustomTarget expands an instance of TargetCustomTarget into a JSON
// request object.
func expandTargetCustomTarget(c *Client, f *TargetCustomTarget, res *Target) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CustomTargetType; !dcl.IsEmptyValueIndirect(v) {
		m["customTargetType"] = v
	}

	return m, nil
}

// flattenTargetCustomTarget flattens an instance of TargetCustomTarget from a JSON
// response object.
func flattenTargetCustomTarget(c *Client, i interface{}, res *Target) *TargetCustomTarget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetCustomTarget{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetCustomTarget
	}
	r.CustomTargetType = dcl.FlattenString(m["customTargetType"])

	return r
}

// expandTargetAssociatedEntitiesMap expands the contents of TargetAssociatedEntities into a JSON
// request object.
func expandTargetAssociatedEntitiesMap(c *Client, f map[string]TargetAssociatedEntities, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetAssociatedEntities(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetAssociatedEntitiesSlice expands the contents of TargetAssociatedEntities into a JSON
// request object.
func expandTargetAssociatedEntitiesSlice(c *Client, f []TargetAssociatedEntities, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetAssociatedEntities(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetAssociatedEntitiesMap flattens the contents of TargetAssociatedEntities from a JSON
// response object.
func flattenTargetAssociatedEntitiesMap(c *Client, i interface{}, res *Target) map[string]TargetAssociatedEntities {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetAssociatedEntities{}
	}

	if len(a) == 0 {
		return map[string]TargetAssociatedEntities{}
	}

	items := make(map[string]TargetAssociatedEntities)
	for k, item := range a {
		items[k] = *flattenTargetAssociatedEntities(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetAssociatedEntitiesSlice flattens the contents of TargetAssociatedEntities from a JSON
// response object.
func flattenTargetAssociatedEntitiesSlice(c *Client, i interface{}, res *Target) []TargetAssociatedEntities {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetAssociatedEntities{}
	}

	if len(a) == 0 {
		return []TargetAssociatedEntities{}
	}

	items := make([]TargetAssociatedEntities, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetAssociatedEntities(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetAssociatedEntities expands an instance of TargetAssociatedEntities into a JSON
// request object.
func expandTargetAssociatedEntities(c *Client, f *TargetAssociatedEntities, res *Target) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandTargetAssociatedEntitiesGkeClustersSlice(c, f.GkeClusters, res); err != nil {
		return nil, fmt.Errorf("error expanding GkeClusters into gkeClusters: %w", err)
	} else if v != nil {
		m["gkeClusters"] = v
	}
	if v, err := expandTargetAssociatedEntitiesAnthosClustersSlice(c, f.AnthosClusters, res); err != nil {
		return nil, fmt.Errorf("error expanding AnthosClusters into anthosClusters: %w", err)
	} else if v != nil {
		m["anthosClusters"] = v
	}

	return m, nil
}

// flattenTargetAssociatedEntities flattens an instance of TargetAssociatedEntities from a JSON
// response object.
func flattenTargetAssociatedEntities(c *Client, i interface{}, res *Target) *TargetAssociatedEntities {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetAssociatedEntities{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetAssociatedEntities
	}
	r.GkeClusters = flattenTargetAssociatedEntitiesGkeClustersSlice(c, m["gkeClusters"], res)
	r.AnthosClusters = flattenTargetAssociatedEntitiesAnthosClustersSlice(c, m["anthosClusters"], res)

	return r
}

// expandTargetAssociatedEntitiesGkeClustersMap expands the contents of TargetAssociatedEntitiesGkeClusters into a JSON
// request object.
func expandTargetAssociatedEntitiesGkeClustersMap(c *Client, f map[string]TargetAssociatedEntitiesGkeClusters, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetAssociatedEntitiesGkeClusters(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetAssociatedEntitiesGkeClustersSlice expands the contents of TargetAssociatedEntitiesGkeClusters into a JSON
// request object.
func expandTargetAssociatedEntitiesGkeClustersSlice(c *Client, f []TargetAssociatedEntitiesGkeClusters, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetAssociatedEntitiesGkeClusters(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetAssociatedEntitiesGkeClustersMap flattens the contents of TargetAssociatedEntitiesGkeClusters from a JSON
// response object.
func flattenTargetAssociatedEntitiesGkeClustersMap(c *Client, i interface{}, res *Target) map[string]TargetAssociatedEntitiesGkeClusters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetAssociatedEntitiesGkeClusters{}
	}

	if len(a) == 0 {
		return map[string]TargetAssociatedEntitiesGkeClusters{}
	}

	items := make(map[string]TargetAssociatedEntitiesGkeClusters)
	for k, item := range a {
		items[k] = *flattenTargetAssociatedEntitiesGkeClusters(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetAssociatedEntitiesGkeClustersSlice flattens the contents of TargetAssociatedEntitiesGkeClusters from a JSON
// response object.
func flattenTargetAssociatedEntitiesGkeClustersSlice(c *Client, i interface{}, res *Target) []TargetAssociatedEntitiesGkeClusters {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetAssociatedEntitiesGkeClusters{}
	}

	if len(a) == 0 {
		return []TargetAssociatedEntitiesGkeClusters{}
	}

	items := make([]TargetAssociatedEntitiesGkeClusters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetAssociatedEntitiesGkeClusters(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetAssociatedEntitiesGkeClusters expands an instance of TargetAssociatedEntitiesGkeClusters into a JSON
// request object.
func expandTargetAssociatedEntitiesGkeClusters(c *Client, f *TargetAssociatedEntitiesGkeClusters, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Cluster; !dcl.IsEmptyValueIndirect(v) {
		m["cluster"] = v
	}
	if v := f.InternalIP; !dcl.IsEmptyValueIndirect(v) {
		m["internalIp"] = v
	}
	if v := f.ProxyUrl; !dcl.IsEmptyValueIndirect(v) {
		m["proxyUrl"] = v
	}

	return m, nil
}

// flattenTargetAssociatedEntitiesGkeClusters flattens an instance of TargetAssociatedEntitiesGkeClusters from a JSON
// response object.
func flattenTargetAssociatedEntitiesGkeClusters(c *Client, i interface{}, res *Target) *TargetAssociatedEntitiesGkeClusters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetAssociatedEntitiesGkeClusters{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetAssociatedEntitiesGkeClusters
	}
	r.Cluster = dcl.FlattenString(m["cluster"])
	r.InternalIP = dcl.FlattenBool(m["internalIp"])
	r.ProxyUrl = dcl.FlattenString(m["proxyUrl"])

	return r
}

// expandTargetAssociatedEntitiesAnthosClustersMap expands the contents of TargetAssociatedEntitiesAnthosClusters into a JSON
// request object.
func expandTargetAssociatedEntitiesAnthosClustersMap(c *Client, f map[string]TargetAssociatedEntitiesAnthosClusters, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandTargetAssociatedEntitiesAnthosClusters(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandTargetAssociatedEntitiesAnthosClustersSlice expands the contents of TargetAssociatedEntitiesAnthosClusters into a JSON
// request object.
func expandTargetAssociatedEntitiesAnthosClustersSlice(c *Client, f []TargetAssociatedEntitiesAnthosClusters, res *Target) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandTargetAssociatedEntitiesAnthosClusters(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenTargetAssociatedEntitiesAnthosClustersMap flattens the contents of TargetAssociatedEntitiesAnthosClusters from a JSON
// response object.
func flattenTargetAssociatedEntitiesAnthosClustersMap(c *Client, i interface{}, res *Target) map[string]TargetAssociatedEntitiesAnthosClusters {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetAssociatedEntitiesAnthosClusters{}
	}

	if len(a) == 0 {
		return map[string]TargetAssociatedEntitiesAnthosClusters{}
	}

	items := make(map[string]TargetAssociatedEntitiesAnthosClusters)
	for k, item := range a {
		items[k] = *flattenTargetAssociatedEntitiesAnthosClusters(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenTargetAssociatedEntitiesAnthosClustersSlice flattens the contents of TargetAssociatedEntitiesAnthosClusters from a JSON
// response object.
func flattenTargetAssociatedEntitiesAnthosClustersSlice(c *Client, i interface{}, res *Target) []TargetAssociatedEntitiesAnthosClusters {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetAssociatedEntitiesAnthosClusters{}
	}

	if len(a) == 0 {
		return []TargetAssociatedEntitiesAnthosClusters{}
	}

	items := make([]TargetAssociatedEntitiesAnthosClusters, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetAssociatedEntitiesAnthosClusters(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandTargetAssociatedEntitiesAnthosClusters expands an instance of TargetAssociatedEntitiesAnthosClusters into a JSON
// request object.
func expandTargetAssociatedEntitiesAnthosClusters(c *Client, f *TargetAssociatedEntitiesAnthosClusters, res *Target) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Membership; !dcl.IsEmptyValueIndirect(v) {
		m["membership"] = v
	}

	return m, nil
}

// flattenTargetAssociatedEntitiesAnthosClusters flattens an instance of TargetAssociatedEntitiesAnthosClusters from a JSON
// response object.
func flattenTargetAssociatedEntitiesAnthosClusters(c *Client, i interface{}, res *Target) *TargetAssociatedEntitiesAnthosClusters {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &TargetAssociatedEntitiesAnthosClusters{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyTargetAssociatedEntitiesAnthosClusters
	}
	r.Membership = dcl.FlattenString(m["membership"])

	return r
}

// flattenTargetExecutionConfigsUsagesEnumMap flattens the contents of TargetExecutionConfigsUsagesEnum from a JSON
// response object.
func flattenTargetExecutionConfigsUsagesEnumMap(c *Client, i interface{}, res *Target) map[string]TargetExecutionConfigsUsagesEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]TargetExecutionConfigsUsagesEnum{}
	}

	if len(a) == 0 {
		return map[string]TargetExecutionConfigsUsagesEnum{}
	}

	items := make(map[string]TargetExecutionConfigsUsagesEnum)
	for k, item := range a {
		items[k] = *flattenTargetExecutionConfigsUsagesEnum(item.(interface{}))
	}

	return items
}

// flattenTargetExecutionConfigsUsagesEnumSlice flattens the contents of TargetExecutionConfigsUsagesEnum from a JSON
// response object.
func flattenTargetExecutionConfigsUsagesEnumSlice(c *Client, i interface{}, res *Target) []TargetExecutionConfigsUsagesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []TargetExecutionConfigsUsagesEnum{}
	}

	if len(a) == 0 {
		return []TargetExecutionConfigsUsagesEnum{}
	}

	items := make([]TargetExecutionConfigsUsagesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenTargetExecutionConfigsUsagesEnum(item.(interface{})))
	}

	return items
}

// flattenTargetExecutionConfigsUsagesEnum asserts that an interface is a string, and returns a
// pointer to a *TargetExecutionConfigsUsagesEnum with the same value as that string.
func flattenTargetExecutionConfigsUsagesEnum(i interface{}) *TargetExecutionConfigsUsagesEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return TargetExecutionConfigsUsagesEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Target) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalTarget(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type targetDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         targetApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToTargetDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]targetDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []targetDiff
	// For each operation name, create a targetDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := targetDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToTargetApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToTargetApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (targetApiOperation, error) {
	switch opName {

	case "updateTargetUpdateTargetOperation":
		return &updateTargetUpdateTargetOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractTargetFields(r *Target) error {
	vGke := r.Gke
	if vGke == nil {
		// note: explicitly not the empty object.
		vGke = &TargetGke{}
	}
	if err := extractTargetGkeFields(r, vGke); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGke) {
		r.Gke = vGke
	}
	vAnthosCluster := r.AnthosCluster
	if vAnthosCluster == nil {
		// note: explicitly not the empty object.
		vAnthosCluster = &TargetAnthosCluster{}
	}
	if err := extractTargetAnthosClusterFields(r, vAnthosCluster); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAnthosCluster) {
		r.AnthosCluster = vAnthosCluster
	}
	vRun := r.Run
	if vRun == nil {
		// note: explicitly not the empty object.
		vRun = &TargetRun{}
	}
	if err := extractTargetRunFields(r, vRun); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRun) {
		r.Run = vRun
	}
	vMultiTarget := r.MultiTarget
	if vMultiTarget == nil {
		// note: explicitly not the empty object.
		vMultiTarget = &TargetMultiTarget{}
	}
	if err := extractTargetMultiTargetFields(r, vMultiTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMultiTarget) {
		r.MultiTarget = vMultiTarget
	}
	vCustomTarget := r.CustomTarget
	if vCustomTarget == nil {
		// note: explicitly not the empty object.
		vCustomTarget = &TargetCustomTarget{}
	}
	if err := extractTargetCustomTargetFields(r, vCustomTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCustomTarget) {
		r.CustomTarget = vCustomTarget
	}
	return nil
}
func extractTargetGkeFields(r *Target, o *TargetGke) error {
	return nil
}
func extractTargetAnthosClusterFields(r *Target, o *TargetAnthosCluster) error {
	return nil
}
func extractTargetExecutionConfigsFields(r *Target, o *TargetExecutionConfigs) error {
	return nil
}
func extractTargetRunFields(r *Target, o *TargetRun) error {
	return nil
}
func extractTargetMultiTargetFields(r *Target, o *TargetMultiTarget) error {
	return nil
}
func extractTargetCustomTargetFields(r *Target, o *TargetCustomTarget) error {
	return nil
}
func extractTargetAssociatedEntitiesFields(r *Target, o *TargetAssociatedEntities) error {
	return nil
}
func extractTargetAssociatedEntitiesGkeClustersFields(r *Target, o *TargetAssociatedEntitiesGkeClusters) error {
	return nil
}
func extractTargetAssociatedEntitiesAnthosClustersFields(r *Target, o *TargetAssociatedEntitiesAnthosClusters) error {
	return nil
}

func postReadExtractTargetFields(r *Target) error {
	vGke := r.Gke
	if vGke == nil {
		// note: explicitly not the empty object.
		vGke = &TargetGke{}
	}
	if err := postReadExtractTargetGkeFields(r, vGke); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGke) {
		r.Gke = vGke
	}
	vAnthosCluster := r.AnthosCluster
	if vAnthosCluster == nil {
		// note: explicitly not the empty object.
		vAnthosCluster = &TargetAnthosCluster{}
	}
	if err := postReadExtractTargetAnthosClusterFields(r, vAnthosCluster); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAnthosCluster) {
		r.AnthosCluster = vAnthosCluster
	}
	vRun := r.Run
	if vRun == nil {
		// note: explicitly not the empty object.
		vRun = &TargetRun{}
	}
	if err := postReadExtractTargetRunFields(r, vRun); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRun) {
		r.Run = vRun
	}
	vMultiTarget := r.MultiTarget
	if vMultiTarget == nil {
		// note: explicitly not the empty object.
		vMultiTarget = &TargetMultiTarget{}
	}
	if err := postReadExtractTargetMultiTargetFields(r, vMultiTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMultiTarget) {
		r.MultiTarget = vMultiTarget
	}
	vCustomTarget := r.CustomTarget
	if vCustomTarget == nil {
		// note: explicitly not the empty object.
		vCustomTarget = &TargetCustomTarget{}
	}
	if err := postReadExtractTargetCustomTargetFields(r, vCustomTarget); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCustomTarget) {
		r.CustomTarget = vCustomTarget
	}
	return nil
}
func postReadExtractTargetGkeFields(r *Target, o *TargetGke) error {
	return nil
}
func postReadExtractTargetAnthosClusterFields(r *Target, o *TargetAnthosCluster) error {
	return nil
}
func postReadExtractTargetExecutionConfigsFields(r *Target, o *TargetExecutionConfigs) error {
	return nil
}
func postReadExtractTargetRunFields(r *Target, o *TargetRun) error {
	return nil
}
func postReadExtractTargetMultiTargetFields(r *Target, o *TargetMultiTarget) error {
	return nil
}
func postReadExtractTargetCustomTargetFields(r *Target, o *TargetCustomTarget) error {
	return nil
}
func postReadExtractTargetAssociatedEntitiesFields(r *Target, o *TargetAssociatedEntities) error {
	return nil
}
func postReadExtractTargetAssociatedEntitiesGkeClustersFields(r *Target, o *TargetAssociatedEntitiesGkeClusters) error {
	return nil
}
func postReadExtractTargetAssociatedEntitiesAnthosClustersFields(r *Target, o *TargetAssociatedEntitiesAnthosClusters) error {
	return nil
}
