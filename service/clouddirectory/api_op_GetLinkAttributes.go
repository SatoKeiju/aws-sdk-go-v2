// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetLinkAttributesRequest
type GetLinkAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of attribute names whose values will be retrieved.
	//
	// AttributeNames is a required field
	AttributeNames []string `type:"list" required:"true"`

	// The consistency level at which to retrieve the attributes on a typed link.
	ConsistencyLevel ConsistencyLevel `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the typed link resides. For more information, see arns or Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Allows a typed link specifier to be accepted as input.
	//
	// TypedLinkSpecifier is a required field
	TypedLinkSpecifier *TypedLinkSpecifier `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetLinkAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLinkAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLinkAttributesInput"}

	if s.AttributeNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeNames"))
	}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.TypedLinkSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypedLinkSpecifier"))
	}
	if s.TypedLinkSpecifier != nil {
		if err := s.TypedLinkSpecifier.Validate(); err != nil {
			invalidParams.AddNested("TypedLinkSpecifier", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLinkAttributesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AttributeNames != nil {
		v := s.AttributeNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AttributeNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConsistencyLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TypedLinkSpecifier != nil {
		v := s.TypedLinkSpecifier

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TypedLinkSpecifier", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetLinkAttributesResponse
type GetLinkAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The attributes that are associated with the typed link.
	Attributes []AttributeKeyAndValue `type:"list"`
}

// String returns the string representation
func (s GetLinkAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLinkAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attributes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetLinkAttributes = "GetLinkAttributes"

// GetLinkAttributesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves attributes that are associated with a typed link.
//
//    // Example sending a request using GetLinkAttributesRequest.
//    req := client.GetLinkAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetLinkAttributes
func (c *Client) GetLinkAttributesRequest(input *GetLinkAttributesInput) GetLinkAttributesRequest {
	op := &aws.Operation{
		Name:       opGetLinkAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/attributes/get",
	}

	if input == nil {
		input = &GetLinkAttributesInput{}
	}

	req := c.newRequest(op, input, &GetLinkAttributesOutput{})
	return GetLinkAttributesRequest{Request: req, Input: input, Copy: c.GetLinkAttributesRequest}
}

// GetLinkAttributesRequest is the request type for the
// GetLinkAttributes API operation.
type GetLinkAttributesRequest struct {
	*aws.Request
	Input *GetLinkAttributesInput
	Copy  func(*GetLinkAttributesInput) GetLinkAttributesRequest
}

// Send marshals and sends the GetLinkAttributes API request.
func (r GetLinkAttributesRequest) Send(ctx context.Context) (*GetLinkAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLinkAttributesResponse{
		GetLinkAttributesOutput: r.Request.Data.(*GetLinkAttributesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLinkAttributesResponse is the response type for the
// GetLinkAttributes API operation.
type GetLinkAttributesResponse struct {
	*GetLinkAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLinkAttributes request.
func (r *GetLinkAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
