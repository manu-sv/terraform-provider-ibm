// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIkePoliciesIDParams creates a new GetIkePoliciesIDParams object
// with the default values initialized.
func NewGetIkePoliciesIDParams() *GetIkePoliciesIDParams {
	var ()
	return &GetIkePoliciesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIkePoliciesIDParamsWithTimeout creates a new GetIkePoliciesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIkePoliciesIDParamsWithTimeout(timeout time.Duration) *GetIkePoliciesIDParams {
	var ()
	return &GetIkePoliciesIDParams{

		timeout: timeout,
	}
}

// NewGetIkePoliciesIDParamsWithContext creates a new GetIkePoliciesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIkePoliciesIDParamsWithContext(ctx context.Context) *GetIkePoliciesIDParams {
	var ()
	return &GetIkePoliciesIDParams{

		Context: ctx,
	}
}

// NewGetIkePoliciesIDParamsWithHTTPClient creates a new GetIkePoliciesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIkePoliciesIDParamsWithHTTPClient(client *http.Client) *GetIkePoliciesIDParams {
	var ()
	return &GetIkePoliciesIDParams{
		HTTPClient: client,
	}
}

/*GetIkePoliciesIDParams contains all the parameters to send to the API endpoint
for the get ike policies ID operation typically these are written to a http.Request
*/
type GetIkePoliciesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The IKE policy idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ike policies ID params
func (o *GetIkePoliciesIDParams) WithTimeout(timeout time.Duration) *GetIkePoliciesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ike policies ID params
func (o *GetIkePoliciesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ike policies ID params
func (o *GetIkePoliciesIDParams) WithContext(ctx context.Context) *GetIkePoliciesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ike policies ID params
func (o *GetIkePoliciesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ike policies ID params
func (o *GetIkePoliciesIDParams) WithHTTPClient(client *http.Client) *GetIkePoliciesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ike policies ID params
func (o *GetIkePoliciesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get ike policies ID params
func (o *GetIkePoliciesIDParams) WithGeneration(generation int64) *GetIkePoliciesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get ike policies ID params
func (o *GetIkePoliciesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get ike policies ID params
func (o *GetIkePoliciesIDParams) WithID(id string) *GetIkePoliciesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ike policies ID params
func (o *GetIkePoliciesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get ike policies ID params
func (o *GetIkePoliciesIDParams) WithVersion(version string) *GetIkePoliciesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get ike policies ID params
func (o *GetIkePoliciesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetIkePoliciesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}