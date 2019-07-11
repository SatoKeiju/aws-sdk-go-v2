// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/TagInput
type TagInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource to which to add tags.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"Arn" min:"12" type:"string" required:"true"`

	// The tags to add to the specified resource. A tag is a string-to-string map
	// of key-value pairs. Tag keys can have a maximum character length of 128 characters,
	// and tag values can have a maximum length of 256 characters.
	//
	// Tags is a required field
	Tags map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s TagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 12))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/TagOutput
type TagOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the tagged resource.
	Arn *string `min:"12" type:"string"`

	// The tags that have been added to the specified resource.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s TagOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opTag = "Tag"

// TagRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Adds tags to a resource group with the specified ARN. Existing tags on a
// resource group are not changed if they are not specified in the request parameters.
//
//    // Example sending a request using TagRequest.
//    req := client.TagRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/Tag
func (c *Client) TagRequest(input *TagInput) TagRequest {
	op := &aws.Operation{
		Name:       opTag,
		HTTPMethod: "PUT",
		HTTPPath:   "/resources/{Arn}/tags",
	}

	if input == nil {
		input = &TagInput{}
	}

	req := c.newRequest(op, input, &TagOutput{})
	return TagRequest{Request: req, Input: input, Copy: c.TagRequest}
}

// TagRequest is the request type for the
// Tag API operation.
type TagRequest struct {
	*aws.Request
	Input *TagInput
	Copy  func(*TagInput) TagRequest
}

// Send marshals and sends the Tag API request.
func (r TagRequest) Send(ctx context.Context) (*TagResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResponse{
		TagOutput: r.Request.Data.(*TagOutput),
		response:  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResponse is the response type for the
// Tag API operation.
type TagResponse struct {
	*TagOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Tag request.
func (r *TagResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
