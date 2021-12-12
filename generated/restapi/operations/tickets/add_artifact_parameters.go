package tickets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/SecurityBrewery/catalyst/generated/models"
	"github.com/SecurityBrewery/catalyst/generated/restapi/api"
)

// AddArtifactEndpoint executes the core logic of the related
// route endpoint.
func AddArtifactEndpoint(handler func(ctx context.Context, params *AddArtifactParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewAddArtifactParams()
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

// NewAddArtifactParams creates a new AddArtifactParams object
// with the default values initialized.
func NewAddArtifactParams() *AddArtifactParams {
	var ()
	return &AddArtifactParams{}
}

// AddArtifactParams contains all the bound params for the add artifact operation
// typically these are obtained from a http.Request
//
// swagger:parameters addArtifact
type AddArtifactParams struct {

	/*Artifact object that needs to be added
	  Required: true
	  In: body
	*/
	Artifact *models.Artifact
	/*Ticket ID
	  Required: true
	  In: path
	*/
	ID int64
}

// ReadRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddArtifactParams) ReadRequest(ctx *gin.Context) error {
	var res []error

	if runtime.HasBody(ctx.Request) {
		var body models.Artifact
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("artifact", "body", ""))
			} else {
				res = append(res, errors.NewParseError("artifact", "body", "", err))
			}

		} else {
			o.Artifact = &body
		}
	} else {
		res = append(res, errors.Required("artifact", "body", ""))
	}

	rID := []string{ctx.Param("id")}
	if err := o.bindID(rID, true); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddArtifactParams) bindID(rawData []string, hasKey bool) error {
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
