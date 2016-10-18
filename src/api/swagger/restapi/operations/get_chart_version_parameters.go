package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetChartVersionParams creates a new GetChartVersionParams object
// with the default values initialized.
func NewGetChartVersionParams() GetChartVersionParams {
	var ()
	return GetChartVersionParams{}
}

// GetChartVersionParams contains all the bound params for the get chart version operation
// typically these are obtained from a http.Request
//
// swagger:parameters getChartVersion
type GetChartVersionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	ChartName string
	/*
	  Required: true
	  In: path
	*/
	Repo string
	/*
	  Required: true
	  In: path
	*/
	Version string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetChartVersionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rChartName, rhkChartName, _ := route.Params.GetOK("chartName")
	if err := o.bindChartName(rChartName, rhkChartName, route.Formats); err != nil {
		res = append(res, err)
	}

	rRepo, rhkRepo, _ := route.Params.GetOK("repo")
	if err := o.bindRepo(rRepo, rhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}

	rVersion, rhkVersion, _ := route.Params.GetOK("version")
	if err := o.bindVersion(rVersion, rhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetChartVersionParams) bindChartName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ChartName = raw

	return nil
}

func (o *GetChartVersionParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Repo = raw

	return nil
}

func (o *GetChartVersionParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Version = raw

	return nil
}