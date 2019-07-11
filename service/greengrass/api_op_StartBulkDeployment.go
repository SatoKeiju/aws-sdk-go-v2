// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Information about a bulk deployment. You cannot start a new bulk deployment
// while another one is still running or in a non-terminal state.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/StartBulkDeploymentRequest
type StartBulkDeploymentInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// The ARN of the execution role to associate with the bulk deployment operation.
	// This IAM role must allow the ''greengrass:CreateDeployment'' action for all
	// group versions that are listed in the input file. This IAM role must have
	// access to the S3 bucket containing the input file.
	ExecutionRoleArn *string `type:"string"`

	// The URI of the input file contained in the S3 bucket. The execution role
	// must have ''getObject'' permissions on this bucket to access the input file.
	// The input file is a JSON-serialized, line delimited file with UTF-8 encoding
	// that provides a list of group and version IDs and the deployment type. This
	// file must be less than 100 MB. Currently, AWS IoT Greengrass supports only
	// ''NewDeployment'' deployment types.
	InputFileUri *string `type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s StartBulkDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartBulkDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ExecutionRoleArn != nil {
		v := *s.ExecutionRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExecutionRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InputFileUri != nil {
		v := *s.InputFileUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InputFileUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/StartBulkDeploymentResponse
type StartBulkDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the bulk deployment.
	BulkDeploymentArn *string `type:"string"`

	// The ID of the bulk deployment.
	BulkDeploymentId *string `type:"string"`
}

// String returns the string representation
func (s StartBulkDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartBulkDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BulkDeploymentArn != nil {
		v := *s.BulkDeploymentArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BulkDeploymentArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BulkDeploymentId != nil {
		v := *s.BulkDeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BulkDeploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartBulkDeployment = "StartBulkDeployment"

// StartBulkDeploymentRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Deploys multiple groups in one operation. This action starts the bulk deployment
// of a specified set of group versions. Each group version deployment will
// be triggered with an adaptive rate that has a fixed upper limit. We recommend
// that you include an ''X-Amzn-Client-Token'' token in every ''StartBulkDeployment''
// request. These requests are idempotent with respect to the token and the
// request parameters.
//
//    // Example sending a request using StartBulkDeploymentRequest.
//    req := client.StartBulkDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/StartBulkDeployment
func (c *Client) StartBulkDeploymentRequest(input *StartBulkDeploymentInput) StartBulkDeploymentRequest {
	op := &aws.Operation{
		Name:       opStartBulkDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/bulk/deployments",
	}

	if input == nil {
		input = &StartBulkDeploymentInput{}
	}

	req := c.newRequest(op, input, &StartBulkDeploymentOutput{})
	return StartBulkDeploymentRequest{Request: req, Input: input, Copy: c.StartBulkDeploymentRequest}
}

// StartBulkDeploymentRequest is the request type for the
// StartBulkDeployment API operation.
type StartBulkDeploymentRequest struct {
	*aws.Request
	Input *StartBulkDeploymentInput
	Copy  func(*StartBulkDeploymentInput) StartBulkDeploymentRequest
}

// Send marshals and sends the StartBulkDeployment API request.
func (r StartBulkDeploymentRequest) Send(ctx context.Context) (*StartBulkDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartBulkDeploymentResponse{
		StartBulkDeploymentOutput: r.Request.Data.(*StartBulkDeploymentOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartBulkDeploymentResponse is the response type for the
// StartBulkDeployment API operation.
type StartBulkDeploymentResponse struct {
	*StartBulkDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartBulkDeployment request.
func (r *StartBulkDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
