// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeploymentsRequest
type ListDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentsInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeploymentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeploymentsResponse
type ListDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of deployments for the requested groups.
	Deployments []Deployment `type:"list"`

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeploymentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Deployments != nil {
		v := s.Deployments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Deployments", metadata)
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

const opListDeployments = "ListDeployments"

// ListDeploymentsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Returns a history of deployments for the group.
//
//    // Example sending a request using ListDeploymentsRequest.
//    req := client.ListDeploymentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeployments
func (c *Client) ListDeploymentsRequest(input *ListDeploymentsInput) ListDeploymentsRequest {
	op := &aws.Operation{
		Name:       opListDeployments,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/deployments",
	}

	if input == nil {
		input = &ListDeploymentsInput{}
	}

	req := c.newRequest(op, input, &ListDeploymentsOutput{})
	return ListDeploymentsRequest{Request: req, Input: input, Copy: c.ListDeploymentsRequest}
}

// ListDeploymentsRequest is the request type for the
// ListDeployments API operation.
type ListDeploymentsRequest struct {
	*aws.Request
	Input *ListDeploymentsInput
	Copy  func(*ListDeploymentsInput) ListDeploymentsRequest
}

// Send marshals and sends the ListDeployments API request.
func (r ListDeploymentsRequest) Send(ctx context.Context) (*ListDeploymentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentsResponse{
		ListDeploymentsOutput: r.Request.Data.(*ListDeploymentsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDeploymentsResponse is the response type for the
// ListDeployments API operation.
type ListDeploymentsResponse struct {
	*ListDeploymentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeployments request.
func (r *ListDeploymentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
