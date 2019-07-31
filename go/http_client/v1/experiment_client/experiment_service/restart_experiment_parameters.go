// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	experiment_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/experiment_model"
)

// NewRestartExperimentParams creates a new RestartExperimentParams object
// with the default values initialized.
func NewRestartExperimentParams() *RestartExperimentParams {
	var ()
	return &RestartExperimentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestartExperimentParamsWithTimeout creates a new RestartExperimentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestartExperimentParamsWithTimeout(timeout time.Duration) *RestartExperimentParams {
	var ()
	return &RestartExperimentParams{

		timeout: timeout,
	}
}

// NewRestartExperimentParamsWithContext creates a new RestartExperimentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestartExperimentParamsWithContext(ctx context.Context) *RestartExperimentParams {
	var ()
	return &RestartExperimentParams{

		Context: ctx,
	}
}

// NewRestartExperimentParamsWithHTTPClient creates a new RestartExperimentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestartExperimentParamsWithHTTPClient(client *http.Client) *RestartExperimentParams {
	var ()
	return &RestartExperimentParams{
		HTTPClient: client,
	}
}

/*RestartExperimentParams contains all the parameters to send to the API endpoint
for the restart experiment operation typically these are written to a http.Request
*/
type RestartExperimentParams struct {

	/*Body*/
	Body *experiment_model.V1OwnedEntityIDRequest
	/*ID
	  Unique integer identifier of the entity

	*/
	ID string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the experiement will be assigned

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restart experiment params
func (o *RestartExperimentParams) WithTimeout(timeout time.Duration) *RestartExperimentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart experiment params
func (o *RestartExperimentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart experiment params
func (o *RestartExperimentParams) WithContext(ctx context.Context) *RestartExperimentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart experiment params
func (o *RestartExperimentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart experiment params
func (o *RestartExperimentParams) WithHTTPClient(client *http.Client) *RestartExperimentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart experiment params
func (o *RestartExperimentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the restart experiment params
func (o *RestartExperimentParams) WithBody(body *experiment_model.V1OwnedEntityIDRequest) *RestartExperimentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the restart experiment params
func (o *RestartExperimentParams) SetBody(body *experiment_model.V1OwnedEntityIDRequest) {
	o.Body = body
}

// WithID adds the id to the restart experiment params
func (o *RestartExperimentParams) WithID(id string) *RestartExperimentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the restart experiment params
func (o *RestartExperimentParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the restart experiment params
func (o *RestartExperimentParams) WithOwner(owner string) *RestartExperimentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the restart experiment params
func (o *RestartExperimentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the restart experiment params
func (o *RestartExperimentParams) WithProject(project string) *RestartExperimentParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the restart experiment params
func (o *RestartExperimentParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *RestartExperimentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
