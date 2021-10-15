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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *LogExclusion) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string(nil)); err != nil {
		return err
	}
	if err := dcl.Required(r, "filter"); err != nil {
		return err
	}
	return nil
}
func (r *LogExclusion) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://logging.googleapis.com/v2/", params)
}

func (r *LogExclusion) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
		"name":   dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/exclusions/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *LogExclusion) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("{{parent}}/exclusions", nr.basePath(), userBasePath, params), nil

}

func (r *LogExclusion) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("{{parent}}/exclusions", nr.basePath(), userBasePath, params), nil

}

func (r *LogExclusion) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
		"name":   dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("{{parent}}/exclusions/{{name}}", nr.basePath(), userBasePath, params), nil
}

// logExclusionApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type logExclusionApiOperation interface {
	do(context.Context, *LogExclusion, *Client) error
}

// newUpdateLogExclusionUpdateExclusionRequest creates a request for an
// LogExclusion resource's UpdateExclusion update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateLogExclusionUpdateExclusionRequest(ctx context.Context, f *LogExclusion, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Filter; !dcl.IsEmptyValueIndirect(v) {
		req["filter"] = v
	}
	if v := f.Disabled; !dcl.IsEmptyValueIndirect(v) {
		req["disabled"] = v
	}
	return req, nil
}

// marshalUpdateLogExclusionUpdateExclusionRequest converts the update into
// the final JSON request body.
func marshalUpdateLogExclusionUpdateExclusionRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateLogExclusionUpdateExclusionOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateLogExclusionUpdateExclusionOperation) do(ctx context.Context, r *LogExclusion, c *Client) error {
	_, err := c.GetLogExclusion(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateExclusion")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateLogExclusionUpdateExclusionRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateLogExclusionUpdateExclusionRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listLogExclusionRaw(ctx context.Context, r *LogExclusion, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != LogExclusionMaxPage {
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

type listLogExclusionOperation struct {
	Exclusions []map[string]interface{} `json:"exclusions"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listLogExclusion(ctx context.Context, r *LogExclusion, pageToken string, pageSize int32) ([]*LogExclusion, string, error) {
	b, err := c.listLogExclusionRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listLogExclusionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*LogExclusion
	for _, v := range m.Exclusions {
		res, err := unmarshalMapLogExclusion(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllLogExclusion(ctx context.Context, f func(*LogExclusion) bool, resources []*LogExclusion) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteLogExclusion(ctx, res)
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

type deleteLogExclusionOperation struct{}

func (op *deleteLogExclusionOperation) do(ctx context.Context, r *LogExclusion, c *Client) error {
	r, err := c.GetLogExclusion(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "LogExclusion not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetLogExclusion checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete LogExclusion: %w", err)
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetLogExclusion(ctx, r)
		if !dcl.IsNotFound(err) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createLogExclusionOperation struct {
	response map[string]interface{}
}

func (op *createLogExclusionOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createLogExclusionOperation) do(ctx context.Context, r *LogExclusion, c *Client) error {
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetLogExclusion(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getLogExclusionRaw(ctx context.Context, r *LogExclusion) ([]byte, error) {

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

func (c *Client) logExclusionDiffsForRawDesired(ctx context.Context, rawDesired *LogExclusion, opts ...dcl.ApplyOption) (initial, desired *LogExclusion, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *LogExclusion
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*LogExclusion); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected LogExclusion, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetLogExclusion(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a LogExclusion resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve LogExclusion resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that LogExclusion resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeLogExclusionDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for LogExclusion: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for LogExclusion: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeLogExclusionInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for LogExclusion: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeLogExclusionDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for LogExclusion: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffLogExclusion(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeLogExclusionInitialState(rawInitial, rawDesired *LogExclusion) (*LogExclusion, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeLogExclusionDesiredState(rawDesired, rawInitial *LogExclusion, opts ...dcl.ApplyOption) (*LogExclusion, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}

	canonicalDesired := &LogExclusion{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.Filter, rawInitial.Filter) {
		canonicalDesired.Filter = rawInitial.Filter
	} else {
		canonicalDesired.Filter = rawDesired.Filter
	}
	if dcl.BoolCanonicalize(rawDesired.Disabled, rawInitial.Disabled) {
		canonicalDesired.Disabled = rawInitial.Disabled
	} else {
		canonicalDesired.Disabled = rawDesired.Disabled
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}

	return canonicalDesired, nil
}

func canonicalizeLogExclusionNewState(c *Client, rawNew, rawDesired *LogExclusion) (*LogExclusion, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Filter) && dcl.IsNotReturnedByServer(rawDesired.Filter) {
		rawNew.Filter = rawDesired.Filter
	} else {
		if dcl.StringCanonicalize(rawDesired.Filter, rawNew.Filter) {
			rawNew.Filter = rawDesired.Filter
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Disabled) && dcl.IsNotReturnedByServer(rawDesired.Disabled) {
		rawNew.Disabled = rawDesired.Disabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.Disabled, rawNew.Disabled) {
			rawNew.Disabled = rawDesired.Disabled
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	rawNew.Parent = rawDesired.Parent

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffLogExclusion(c *Client, desired, actual *LogExclusion, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogExclusionUpdateExclusionOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Filter, actual.Filter, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogExclusionUpdateExclusionOperation")}, fn.AddNest("Filter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Disabled, actual.Disabled, dcl.Info{OperationSelector: dcl.TriggersOperation("updateLogExclusionUpdateExclusionOperation")}, fn.AddNest("Disabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *LogExclusion) urlNormalized() *LogExclusion {
	normalized := dcl.Copy(*r).(LogExclusion)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Filter = dcl.SelfLinkToName(r.Filter)
	normalized.Parent = r.Parent
	return &normalized
}

func (r *LogExclusion) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateExclusion" {
		fields := map[string]interface{}{
			"parent": dcl.ValueOrEmptyString(nr.Parent),
			"name":   dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("{{parent}}/exclusions/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the LogExclusion resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *LogExclusion) marshal(c *Client) ([]byte, error) {
	m, err := expandLogExclusion(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling LogExclusion: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalLogExclusion decodes JSON responses into the LogExclusion resource schema.
func unmarshalLogExclusion(b []byte, c *Client) (*LogExclusion, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapLogExclusion(m, c)
}

func unmarshalMapLogExclusion(m map[string]interface{}, c *Client) (*LogExclusion, error) {

	flattened := flattenLogExclusion(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandLogExclusion expands LogExclusion into a JSON request object.
func expandLogExclusion(c *Client, f *LogExclusion) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Filter; dcl.ValueShouldBeSent(v) {
		m["filter"] = v
	}
	if v := f.Disabled; dcl.ValueShouldBeSent(v) {
		m["disabled"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if v != nil {
		m["parent"] = v
	}

	return m, nil
}

// flattenLogExclusion flattens LogExclusion from a JSON request object into the
// LogExclusion type.
func flattenLogExclusion(c *Client, i interface{}) *LogExclusion {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &LogExclusion{}
	res.Name = dcl.FlattenString(m["name"])
	res.Description = dcl.FlattenString(m["description"])
	res.Filter = dcl.FlattenString(m["filter"])
	res.Disabled = dcl.FlattenBool(m["disabled"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Parent = dcl.FlattenString(m["parent"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *LogExclusion) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalLogExclusion(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Parent == nil && ncr.Parent == nil {
			c.Config.Logger.Info("Both Parent fields null - considering equal.")
		} else if nr.Parent == nil || ncr.Parent == nil {
			c.Config.Logger.Info("Only one Parent field is null - considering unequal.")
			return false
		} else if *nr.Parent != *ncr.Parent {
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

type logExclusionDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         logExclusionApiOperation
}

func convertFieldDiffsToLogExclusionDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]logExclusionDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []logExclusionDiff
	// For each operation name, create a logExclusionDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := logExclusionDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToLogExclusionApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToLogExclusionApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (logExclusionApiOperation, error) {
	switch opName {

	case "updateLogExclusionUpdateExclusionOperation":
		return &updateLogExclusionUpdateExclusionOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractLogExclusionFields(r *LogExclusion) error {
	return nil
}
