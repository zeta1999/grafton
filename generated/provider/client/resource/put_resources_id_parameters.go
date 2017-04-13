package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/manifoldco/grafton/generated/provider/models"
)

// NewPutResourcesIDParams creates a new PutResourcesIDParams object
// with the default values initialized.
func NewPutResourcesIDParams() *PutResourcesIDParams {
	var ()
	return &PutResourcesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutResourcesIDParamsWithTimeout creates a new PutResourcesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutResourcesIDParamsWithTimeout(timeout time.Duration) *PutResourcesIDParams {
	var ()
	return &PutResourcesIDParams{

		timeout: timeout,
	}
}

// NewPutResourcesIDParamsWithContext creates a new PutResourcesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutResourcesIDParamsWithContext(ctx context.Context) *PutResourcesIDParams {
	var ()
	return &PutResourcesIDParams{

		Context: ctx,
	}
}

// NewPutResourcesIDParamsWithHTTPClient creates a new PutResourcesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutResourcesIDParamsWithHTTPClient(client *http.Client) *PutResourcesIDParams {
	var ()
	return &PutResourcesIDParams{
		HTTPClient: client,
	}
}

/*PutResourcesIDParams contains all the parameters to send to the API endpoint
for the put resources ID operation typically these are written to a http.Request
*/
type PutResourcesIDParams struct {

	/*Date
	  Timestamp of when the request was issued from Manifold in UTC.

	*/
	Date strfmt.DateTime
	/*XCallbackID
	  ID of the Callback for completing this request if the provider returns a
	`202 Accepted`, stored as a base 32 encoded 18 byte identifier.


	*/
	XCallbackID string
	/*XCallbackURL
	  The URL the provider calls to complete the request if a `202 Accepted` is
	returned.


	*/
	XCallbackURL string
	/*XSignature
	  Header containing the signature, ephemeral public key, and
	signature of the used public key signed by the Manifold root
	signing key to prove authenticity of the request.

	```
	X-Signature: [request-signature] [public-key-value] [signature-of-public-key]
	```

	examples:

	```
	X-Signature: 96afMc5FVZjhGQ4YLPyRW9tcYoPKyj1EMUxkzma_Q4c WydEYGQb7Y4ER6KPAc-YuIwAqFUpII5P9U3MAZ3wxAQ ozhcosOmuWltP8r1BAs-0kccV1AkbHcKYLSjU0dGUHY
	```


	*/
	XSignature string
	/*XSignedHeaders
	  Comma delimited ordered list of header fields used in generating
	the request signature.


	*/
	XSignedHeaders string
	/*Body
	  Resource Provisioning Request

	*/
	Body *models.ResourceRequest
	/*ID
	  ID of a Resource object, stored as a base32 encoded 18 byte identifier.


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put resources ID params
func (o *PutResourcesIDParams) WithTimeout(timeout time.Duration) *PutResourcesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put resources ID params
func (o *PutResourcesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put resources ID params
func (o *PutResourcesIDParams) WithContext(ctx context.Context) *PutResourcesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put resources ID params
func (o *PutResourcesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put resources ID params
func (o *PutResourcesIDParams) WithHTTPClient(client *http.Client) *PutResourcesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put resources ID params
func (o *PutResourcesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the put resources ID params
func (o *PutResourcesIDParams) WithDate(date strfmt.DateTime) *PutResourcesIDParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the put resources ID params
func (o *PutResourcesIDParams) SetDate(date strfmt.DateTime) {
	o.Date = date
}

// WithXCallbackID adds the xCallbackID to the put resources ID params
func (o *PutResourcesIDParams) WithXCallbackID(xCallbackID string) *PutResourcesIDParams {
	o.SetXCallbackID(xCallbackID)
	return o
}

// SetXCallbackID adds the xCallbackId to the put resources ID params
func (o *PutResourcesIDParams) SetXCallbackID(xCallbackID string) {
	o.XCallbackID = xCallbackID
}

// WithXCallbackURL adds the xCallbackURL to the put resources ID params
func (o *PutResourcesIDParams) WithXCallbackURL(xCallbackURL string) *PutResourcesIDParams {
	o.SetXCallbackURL(xCallbackURL)
	return o
}

// SetXCallbackURL adds the xCallbackUrl to the put resources ID params
func (o *PutResourcesIDParams) SetXCallbackURL(xCallbackURL string) {
	o.XCallbackURL = xCallbackURL
}

// WithXSignature adds the xSignature to the put resources ID params
func (o *PutResourcesIDParams) WithXSignature(xSignature string) *PutResourcesIDParams {
	o.SetXSignature(xSignature)
	return o
}

// SetXSignature adds the xSignature to the put resources ID params
func (o *PutResourcesIDParams) SetXSignature(xSignature string) {
	o.XSignature = xSignature
}

// WithXSignedHeaders adds the xSignedHeaders to the put resources ID params
func (o *PutResourcesIDParams) WithXSignedHeaders(xSignedHeaders string) *PutResourcesIDParams {
	o.SetXSignedHeaders(xSignedHeaders)
	return o
}

// SetXSignedHeaders adds the xSignedHeaders to the put resources ID params
func (o *PutResourcesIDParams) SetXSignedHeaders(xSignedHeaders string) {
	o.XSignedHeaders = xSignedHeaders
}

// WithBody adds the body to the put resources ID params
func (o *PutResourcesIDParams) WithBody(body *models.ResourceRequest) *PutResourcesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put resources ID params
func (o *PutResourcesIDParams) SetBody(body *models.ResourceRequest) {
	o.Body = body
}

// WithID adds the id to the put resources ID params
func (o *PutResourcesIDParams) WithID(id string) *PutResourcesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put resources ID params
func (o *PutResourcesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutResourcesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Date
	if err := r.SetHeaderParam("Date", o.Date.String()); err != nil {
		return err
	}

	// header param X-Callback-ID
	if err := r.SetHeaderParam("X-Callback-ID", o.XCallbackID); err != nil {
		return err
	}

	// header param X-Callback-URL
	if err := r.SetHeaderParam("X-Callback-URL", o.XCallbackURL); err != nil {
		return err
	}

	// header param X-Signature
	if err := r.SetHeaderParam("X-Signature", o.XSignature); err != nil {
		return err
	}

	// header param X-Signed-Headers
	if err := r.SetHeaderParam("X-Signed-Headers", o.XSignedHeaders); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models.ResourceRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
