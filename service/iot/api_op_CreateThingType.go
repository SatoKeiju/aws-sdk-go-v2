// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the CreateThingType operation.
type CreateThingTypeInput struct {
	_ struct{} `type:"structure"`

	// Metadata which can be used to manage the thing type.
	Tags []Tag `locationName:"tags" type:"list"`

	// The name of the thing type.
	//
	// ThingTypeName is a required field
	ThingTypeName *string `location:"uri" locationName:"thingTypeName" min:"1" type:"string" required:"true"`

	// The ThingTypeProperties for the thing type to create. It contains information
	// about the new thing type including a description, and a list of searchable
	// thing attribute names.
	ThingTypeProperties *ThingTypeProperties `locationName:"thingTypeProperties" type:"structure"`
}

// String returns the string representation
func (s CreateThingTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateThingTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateThingTypeInput"}

	if s.ThingTypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingTypeName"))
	}
	if s.ThingTypeName != nil && len(*s.ThingTypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingTypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThingTypeInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ThingTypeProperties != nil {
		v := s.ThingTypeProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingTypeProperties", v, metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output of the CreateThingType operation.
type CreateThingTypeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the thing type.
	ThingTypeArn *string `locationName:"thingTypeArn" type:"string"`

	// The thing type ID.
	ThingTypeId *string `locationName:"thingTypeId" type:"string"`

	// The name of the thing type.
	ThingTypeName *string `locationName:"thingTypeName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateThingTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThingTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ThingTypeArn != nil {
		v := *s.ThingTypeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeId != nil {
		v := *s.ThingTypeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateThingType = "CreateThingType"

// CreateThingTypeRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a new thing type.
//
//    // Example sending a request using CreateThingTypeRequest.
//    req := client.CreateThingTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateThingTypeRequest(input *CreateThingTypeInput) CreateThingTypeRequest {
	op := &aws.Operation{
		Name:       opCreateThingType,
		HTTPMethod: "POST",
		HTTPPath:   "/thing-types/{thingTypeName}",
	}

	if input == nil {
		input = &CreateThingTypeInput{}
	}

	req := c.newRequest(op, input, &CreateThingTypeOutput{})
	return CreateThingTypeRequest{Request: req, Input: input, Copy: c.CreateThingTypeRequest}
}

// CreateThingTypeRequest is the request type for the
// CreateThingType API operation.
type CreateThingTypeRequest struct {
	*aws.Request
	Input *CreateThingTypeInput
	Copy  func(*CreateThingTypeInput) CreateThingTypeRequest
}

// Send marshals and sends the CreateThingType API request.
func (r CreateThingTypeRequest) Send(ctx context.Context) (*CreateThingTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateThingTypeResponse{
		CreateThingTypeOutput: r.Request.Data.(*CreateThingTypeOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateThingTypeResponse is the response type for the
// CreateThingType API operation.
type CreateThingTypeResponse struct {
	*CreateThingTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateThingType request.
func (r *CreateThingTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
