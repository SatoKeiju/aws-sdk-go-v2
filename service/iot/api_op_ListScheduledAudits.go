// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListScheduledAuditsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListScheduledAuditsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListScheduledAuditsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListScheduledAuditsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListScheduledAuditsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListScheduledAuditsOutput struct {
	_ struct{} `type:"structure"`

	// A token that can be used to retrieve the next set of results, or null if
	// there are no additional results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of scheduled audits.
	ScheduledAudits []ScheduledAuditMetadata `locationName:"scheduledAudits" type:"list"`
}

// String returns the string representation
func (s ListScheduledAuditsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListScheduledAuditsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ScheduledAudits != nil {
		v := s.ScheduledAudits

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "scheduledAudits", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListScheduledAudits = "ListScheduledAudits"

// ListScheduledAuditsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists all of your scheduled audits.
//
//    // Example sending a request using ListScheduledAuditsRequest.
//    req := client.ListScheduledAuditsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListScheduledAuditsRequest(input *ListScheduledAuditsInput) ListScheduledAuditsRequest {
	op := &aws.Operation{
		Name:       opListScheduledAudits,
		HTTPMethod: "GET",
		HTTPPath:   "/audit/scheduledaudits",
	}

	if input == nil {
		input = &ListScheduledAuditsInput{}
	}

	req := c.newRequest(op, input, &ListScheduledAuditsOutput{})
	return ListScheduledAuditsRequest{Request: req, Input: input, Copy: c.ListScheduledAuditsRequest}
}

// ListScheduledAuditsRequest is the request type for the
// ListScheduledAudits API operation.
type ListScheduledAuditsRequest struct {
	*aws.Request
	Input *ListScheduledAuditsInput
	Copy  func(*ListScheduledAuditsInput) ListScheduledAuditsRequest
}

// Send marshals and sends the ListScheduledAudits API request.
func (r ListScheduledAuditsRequest) Send(ctx context.Context) (*ListScheduledAuditsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListScheduledAuditsResponse{
		ListScheduledAuditsOutput: r.Request.Data.(*ListScheduledAuditsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListScheduledAuditsResponse is the response type for the
// ListScheduledAudits API operation.
type ListScheduledAuditsResponse struct {
	*ListScheduledAuditsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListScheduledAudits request.
func (r *ListScheduledAuditsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
