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

package k8_s_secrets_v1

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

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateK8SSecretsParams creates a new CreateK8SSecretsParams object
// with the default values initialized.
func NewCreateK8SSecretsParams() *CreateK8SSecretsParams {
	var ()
	return &CreateK8SSecretsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateK8SSecretsParamsWithTimeout creates a new CreateK8SSecretsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateK8SSecretsParamsWithTimeout(timeout time.Duration) *CreateK8SSecretsParams {
	var ()
	return &CreateK8SSecretsParams{

		timeout: timeout,
	}
}

// NewCreateK8SSecretsParamsWithContext creates a new CreateK8SSecretsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateK8SSecretsParamsWithContext(ctx context.Context) *CreateK8SSecretsParams {
	var ()
	return &CreateK8SSecretsParams{

		Context: ctx,
	}
}

// NewCreateK8SSecretsParamsWithHTTPClient creates a new CreateK8SSecretsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateK8SSecretsParamsWithHTTPClient(client *http.Client) *CreateK8SSecretsParams {
	var ()
	return &CreateK8SSecretsParams{
		HTTPClient: client,
	}
}

/*CreateK8SSecretsParams contains all the parameters to send to the API endpoint
for the create k8 s secrets operation typically these are written to a http.Request
*/
type CreateK8SSecretsParams struct {

	/*Body
	  Artifact store body

	*/
	Body *service_model.V1K8SResource
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create k8 s secrets params
func (o *CreateK8SSecretsParams) WithTimeout(timeout time.Duration) *CreateK8SSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create k8 s secrets params
func (o *CreateK8SSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create k8 s secrets params
func (o *CreateK8SSecretsParams) WithContext(ctx context.Context) *CreateK8SSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create k8 s secrets params
func (o *CreateK8SSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create k8 s secrets params
func (o *CreateK8SSecretsParams) WithHTTPClient(client *http.Client) *CreateK8SSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create k8 s secrets params
func (o *CreateK8SSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create k8 s secrets params
func (o *CreateK8SSecretsParams) WithBody(body *service_model.V1K8SResource) *CreateK8SSecretsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create k8 s secrets params
func (o *CreateK8SSecretsParams) SetBody(body *service_model.V1K8SResource) {
	o.Body = body
}

// WithOwner adds the owner to the create k8 s secrets params
func (o *CreateK8SSecretsParams) WithOwner(owner string) *CreateK8SSecretsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create k8 s secrets params
func (o *CreateK8SSecretsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateK8SSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}