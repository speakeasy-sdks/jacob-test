// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/jacob-test/v3/pkg/models/shared"
	"net/http"
)

type ShowPetByIDRequest struct {
	// The id of the pet to retrieve
	PetID string `pathParam:"style=simple,explode=false,name=petId"`
}

func (o *ShowPetByIDRequest) GetPetID() string {
	if o == nil {
		return ""
	}
	return o.PetID
}

type ShowPetByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// unexpected error
	Error *shared.Error
	// Expected response to a valid request
	Pet *shared.Pet
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ShowPetByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ShowPetByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ShowPetByIDResponse) GetPet() *shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}

func (o *ShowPetByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ShowPetByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
