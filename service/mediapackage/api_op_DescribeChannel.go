// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeChannelRequest
type DescribeChannelInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeChannelInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeChannelResponse
type DescribeChannelOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *HlsIngest `locationName:"hlsIngest" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsIngest != nil {
		v := s.HlsIngest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsIngest", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opDescribeChannel = "DescribeChannel"

// DescribeChannelRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Gets details about a Channel.
//
//    // Example sending a request using DescribeChannelRequest.
//    req := client.DescribeChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeChannel
func (c *Client) DescribeChannelRequest(input *DescribeChannelInput) DescribeChannelRequest {
	op := &aws.Operation{
		Name:       opDescribeChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/channels/{id}",
	}

	if input == nil {
		input = &DescribeChannelInput{}
	}

	req := c.newRequest(op, input, &DescribeChannelOutput{})
	return DescribeChannelRequest{Request: req, Input: input, Copy: c.DescribeChannelRequest}
}

// DescribeChannelRequest is the request type for the
// DescribeChannel API operation.
type DescribeChannelRequest struct {
	*aws.Request
	Input *DescribeChannelInput
	Copy  func(*DescribeChannelInput) DescribeChannelRequest
}

// Send marshals and sends the DescribeChannel API request.
func (r DescribeChannelRequest) Send(ctx context.Context) (*DescribeChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChannelResponse{
		DescribeChannelOutput: r.Request.Data.(*DescribeChannelOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChannelResponse is the response type for the
// DescribeChannel API operation.
type DescribeChannelResponse struct {
	*DescribeChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChannel request.
func (r *DescribeChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
