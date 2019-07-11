// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListBulkDeploymentDetailedReportsRequest
type ListBulkDeploymentDetailedReportsInput struct {
	_ struct{} `type:"structure"`

	// BulkDeploymentId is a required field
	BulkDeploymentId *string `location:"uri" locationName:"BulkDeploymentId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListBulkDeploymentDetailedReportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBulkDeploymentDetailedReportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBulkDeploymentDetailedReportsInput"}

	if s.BulkDeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BulkDeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBulkDeploymentDetailedReportsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BulkDeploymentId != nil {
		v := *s.BulkDeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "BulkDeploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListBulkDeploymentDetailedReportsResponse
type ListBulkDeploymentDetailedReportsOutput struct {
	_ struct{} `type:"structure"`

	// A list of the individual group deployments in the bulk deployment operation.
	Deployments []BulkDeploymentResult `type:"list"`

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListBulkDeploymentDetailedReportsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBulkDeploymentDetailedReportsOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opListBulkDeploymentDetailedReports = "ListBulkDeploymentDetailedReports"

// ListBulkDeploymentDetailedReportsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Gets a paginated list of the deployments that have been started in a bulk
// deployment operation, and their current deployment status.
//
//    // Example sending a request using ListBulkDeploymentDetailedReportsRequest.
//    req := client.ListBulkDeploymentDetailedReportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListBulkDeploymentDetailedReports
func (c *Client) ListBulkDeploymentDetailedReportsRequest(input *ListBulkDeploymentDetailedReportsInput) ListBulkDeploymentDetailedReportsRequest {
	op := &aws.Operation{
		Name:       opListBulkDeploymentDetailedReports,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/bulk/deployments/{BulkDeploymentId}/detailed-reports",
	}

	if input == nil {
		input = &ListBulkDeploymentDetailedReportsInput{}
	}

	req := c.newRequest(op, input, &ListBulkDeploymentDetailedReportsOutput{})
	return ListBulkDeploymentDetailedReportsRequest{Request: req, Input: input, Copy: c.ListBulkDeploymentDetailedReportsRequest}
}

// ListBulkDeploymentDetailedReportsRequest is the request type for the
// ListBulkDeploymentDetailedReports API operation.
type ListBulkDeploymentDetailedReportsRequest struct {
	*aws.Request
	Input *ListBulkDeploymentDetailedReportsInput
	Copy  func(*ListBulkDeploymentDetailedReportsInput) ListBulkDeploymentDetailedReportsRequest
}

// Send marshals and sends the ListBulkDeploymentDetailedReports API request.
func (r ListBulkDeploymentDetailedReportsRequest) Send(ctx context.Context) (*ListBulkDeploymentDetailedReportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBulkDeploymentDetailedReportsResponse{
		ListBulkDeploymentDetailedReportsOutput: r.Request.Data.(*ListBulkDeploymentDetailedReportsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBulkDeploymentDetailedReportsResponse is the response type for the
// ListBulkDeploymentDetailedReports API operation.
type ListBulkDeploymentDetailedReportsResponse struct {
	*ListBulkDeploymentDetailedReportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBulkDeploymentDetailedReports request.
func (r *ListBulkDeploymentDetailedReportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
