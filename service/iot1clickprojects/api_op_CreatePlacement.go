// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/CreatePlacementRequest
type CreatePlacementInput struct {
	_ struct{} `type:"structure"`

	// Optional user-defined key/value pairs providing contextual data (such as
	// location or function) for the placement.
	Attributes map[string]string `locationName:"attributes" type:"map"`

	// The name of the placement to be created.
	//
	// PlacementName is a required field
	PlacementName *string `locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project in which to create the placement.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePlacementInput"}

	if s.PlacementName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementName"))
	}
	if s.PlacementName != nil && len(*s.PlacementName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementName", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePlacementInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.PlacementName != nil {
		v := *s.PlacementName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "placementName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/CreatePlacementResponse
type CreatePlacementOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreatePlacementOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePlacementOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreatePlacement = "CreatePlacement"

// CreatePlacementRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Creates an empty placement.
//
//    // Example sending a request using CreatePlacementRequest.
//    req := client.CreatePlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/CreatePlacement
func (c *Client) CreatePlacementRequest(input *CreatePlacementInput) CreatePlacementRequest {
	op := &aws.Operation{
		Name:       opCreatePlacement,
		HTTPMethod: "POST",
		HTTPPath:   "/projects/{projectName}/placements",
	}

	if input == nil {
		input = &CreatePlacementInput{}
	}

	req := c.newRequest(op, input, &CreatePlacementOutput{})
	return CreatePlacementRequest{Request: req, Input: input, Copy: c.CreatePlacementRequest}
}

// CreatePlacementRequest is the request type for the
// CreatePlacement API operation.
type CreatePlacementRequest struct {
	*aws.Request
	Input *CreatePlacementInput
	Copy  func(*CreatePlacementInput) CreatePlacementRequest
}

// Send marshals and sends the CreatePlacement API request.
func (r CreatePlacementRequest) Send(ctx context.Context) (*CreatePlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlacementResponse{
		CreatePlacementOutput: r.Request.Data.(*CreatePlacementOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlacementResponse is the response type for the
// CreatePlacement API operation.
type CreatePlacementResponse struct {
	*CreatePlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlacement request.
func (r *CreatePlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
