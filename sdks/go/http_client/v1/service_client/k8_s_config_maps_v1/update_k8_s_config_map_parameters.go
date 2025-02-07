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

package k8_s_config_maps_v1

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

// NewUpdateK8SConfigMapParams creates a new UpdateK8SConfigMapParams object
// with the default values initialized.
func NewUpdateK8SConfigMapParams() *UpdateK8SConfigMapParams {
	var ()
	return &UpdateK8SConfigMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateK8SConfigMapParamsWithTimeout creates a new UpdateK8SConfigMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateK8SConfigMapParamsWithTimeout(timeout time.Duration) *UpdateK8SConfigMapParams {
	var ()
	return &UpdateK8SConfigMapParams{

		timeout: timeout,
	}
}

// NewUpdateK8SConfigMapParamsWithContext creates a new UpdateK8SConfigMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateK8SConfigMapParamsWithContext(ctx context.Context) *UpdateK8SConfigMapParams {
	var ()
	return &UpdateK8SConfigMapParams{

		Context: ctx,
	}
}

// NewUpdateK8SConfigMapParamsWithHTTPClient creates a new UpdateK8SConfigMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateK8SConfigMapParamsWithHTTPClient(client *http.Client) *UpdateK8SConfigMapParams {
	var ()
	return &UpdateK8SConfigMapParams{
		HTTPClient: client,
	}
}

/*UpdateK8SConfigMapParams contains all the parameters to send to the API endpoint
for the update k8 s config map operation typically these are written to a http.Request
*/
type UpdateK8SConfigMapParams struct {

	/*Body
	  Artifact store body

	*/
	Body *service_model.V1K8SResource
	/*K8sResourceUUID
	  UUID

	*/
	K8sResourceUUID string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) WithTimeout(timeout time.Duration) *UpdateK8SConfigMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) WithContext(ctx context.Context) *UpdateK8SConfigMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) WithHTTPClient(client *http.Client) *UpdateK8SConfigMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) WithBody(body *service_model.V1K8SResource) *UpdateK8SConfigMapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) SetBody(body *service_model.V1K8SResource) {
	o.Body = body
}

// WithK8sResourceUUID adds the k8sResourceUUID to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) WithK8sResourceUUID(k8sResourceUUID string) *UpdateK8SConfigMapParams {
	o.SetK8sResourceUUID(k8sResourceUUID)
	return o
}

// SetK8sResourceUUID adds the k8sResourceUuid to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) SetK8sResourceUUID(k8sResourceUUID string) {
	o.K8sResourceUUID = k8sResourceUUID
}

// WithOwner adds the owner to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) WithOwner(owner string) *UpdateK8SConfigMapParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update k8 s config map params
func (o *UpdateK8SConfigMapParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateK8SConfigMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param k8s_resource.uuid
	if err := r.SetPathParam("k8s_resource.uuid", o.K8sResourceUUID); err != nil {
		return err
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
