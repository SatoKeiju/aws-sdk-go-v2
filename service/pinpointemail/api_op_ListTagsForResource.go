// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListTagsForResourceRequest
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource that you want to retrieve
	// tag information for.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"querystring" locationName:"ResourceArn" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForResourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "ResourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListTagsForResourceResponse
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// An array that lists all the tags that are associated with the resource. Each
	// tag consists of a required tag key (Key) and an associated tag value (Value)
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Retrieve a list of the tags (keys and values) that are associated with a
// specified resource. A tag is a label that you optionally define and associate
// with a resource in Amazon Pinpoint. Each tag consists of a required tag key
// and an optional associated tag value. A tag key is a general label that acts
// as a category for more specific tag values. A tag value acts as a descriptor
// within a tag key.
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/tags",
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})
	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
