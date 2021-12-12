package tickets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/SecurityBrewery/catalyst/generated/restapi/api"
)

// SetSchemaEndpoint executes the core logic of the related
// route endpoint.
func SetSchemaEndpoint(handler func(ctx context.Context, params *SetSchemaParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewSetSchemaParams()
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

// NewSetSchemaParams creates a new SetSchemaParams object
// with the default values initialized.
func NewSetSchemaParams() *SetSchemaParams {
	var ()
	return &SetSchemaParams{}
}

// SetSchemaParams contains all the bound params for the set schema operation
// typically these are obtained from a http.Request
//
// swagger:parameters setSchema
type SetSchemaParams struct {

	/*Ticket ID
	  Required: true
	  In: path
	*/
	ID int64
	/*New ticket schema
	  In: body
	*/
	Schema string
}

// ReadRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *SetSchemaParams) ReadRequest(ctx *gin.Context) error {
	var res []error

	rID := []string{ctx.Param("id")}
	if err := o.bindID(rID, true); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(ctx.Request) {
		var body string
		if err := ctx.BindJSON(&body); err != nil {
			res = append(res, errors.NewParseError("schema", "body", "", err))
		} else {
			o.Schema = body
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetSchemaParams) bindID(rawData []string, hasKey bool) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}
