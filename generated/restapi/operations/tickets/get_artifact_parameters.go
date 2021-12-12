package tickets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	"github.com/SecurityBrewery/catalyst/generated/restapi/api"
)

// GetArtifactEndpoint executes the core logic of the related
// route endpoint.
func GetArtifactEndpoint(handler func(ctx context.Context, params *GetArtifactParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewGetArtifactParams()
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

// NewGetArtifactParams creates a new GetArtifactParams object
// with the default values initialized.
func NewGetArtifactParams() *GetArtifactParams {
	var ()
	return &GetArtifactParams{}
}

// GetArtifactParams contains all the bound params for the get artifact operation
// typically these are obtained from a http.Request
//
// swagger:parameters getArtifact
type GetArtifactParams struct {

	/*Ticket ID
	  Required: true
	  In: path
	*/
	ID int64
	/*
	  Required: true
	  In: path
	*/
	Name string
}

// ReadRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetArtifactParams) ReadRequest(ctx *gin.Context) error {
	var res []error

	rID := []string{ctx.Param("id")}
	if err := o.bindID(rID, true); err != nil {
		res = append(res, err)
	}

	rName := []string{ctx.Param("name")}
	if err := o.bindName(rName, true); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetArtifactParams) bindID(rawData []string, hasKey bool) error {
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

func (o *GetArtifactParams) bindName(rawData []string, hasKey bool) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Name = raw

	return nil
}
