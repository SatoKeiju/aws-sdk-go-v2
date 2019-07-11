// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupsRequest
type ListGroupsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupsResponse
type ListGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Information about a group.
	Groups []GroupInformation `type:"list"`

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Groups != nil {
		v := s.Groups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Groups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListGroups = "ListGroups"

// ListGroupsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of groups.
//
//    // Example sending a request using ListGroupsRequest.
//    req := client.ListGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroups
func (c *Client) ListGroupsRequest(input *ListGroupsInput) ListGroupsRequest {
	op := &aws.Operation{
		Name:       opListGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups",
	}

	if input == nil {
		input = &ListGroupsInput{}
	}

	req := c.newRequest(op, input, &ListGroupsOutput{})
	return ListGroupsRequest{Request: req, Input: input, Copy: c.ListGroupsRequest}
}

// ListGroupsRequest is the request type for the
// ListGroups API operation.
type ListGroupsRequest struct {
	*aws.Request
	Input *ListGroupsInput
	Copy  func(*ListGroupsInput) ListGroupsRequest
}

// Send marshals and sends the ListGroups API request.
func (r ListGroupsRequest) Send(ctx context.Context) (*ListGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupsResponse{
		ListGroupsOutput: r.Request.Data.(*ListGroupsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGroupsResponse is the response type for the
// ListGroups API operation.
type ListGroupsResponse struct {
	*ListGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroups request.
func (r *ListGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
