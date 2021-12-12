package playbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"

	"github.com/SecurityBrewery/catalyst/generated/restapi/api"
)

// GetPlaybookEndpoint executes the core logic of the related
// route endpoint.
func GetPlaybookEndpoint(handler func(ctx context.Context, params *GetPlaybookParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewGetPlaybookParams()
		err := params.ReadRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(int(errObj.Code()), gin.H{"error": errObj.Error()})
			return
		}

		resp := handler(ctx, params)

		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// NewGetPlaybookParams creates a new GetPlaybookParams object
// with the default values initialized.
func NewGetPlaybookParams() *GetPlaybookParams {
	var ()
	return &GetPlaybookParams{}
}

// GetPlaybookParams contains all the bound params for the get playbook operation
// typically these are obtained from a http.Request
//
// swagger:parameters getPlaybook
type GetPlaybookParams struct {

	/*Playbook name
	  Required: true
	  In: path
	*/
	ID string
}

// ReadRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetPlaybookParams) ReadRequest(ctx *gin.Context) error {
	var res []error

	rID := []string{ctx.Param("id")}
	if err := o.bindID(rID, true); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPlaybookParams) bindID(rawData []string, hasKey bool) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}
