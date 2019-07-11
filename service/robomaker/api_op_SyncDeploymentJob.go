// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/SyncDeploymentJobRequest
type SyncDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The target fleet for the synchronization.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SyncDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SyncDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SyncDeploymentJobInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SyncDeploymentJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/SyncDeploymentJobResponse
type SyncDeploymentJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the synchronization request.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// Information about the deployment application configurations.
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list"`

	// Information about the deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The failure code if the job fails:
	//
	// InternalServiceError
	//
	// Internal service error.
	//
	// RobotApplicationCrash
	//
	// Robot application exited abnormally.
	//
	// SimulationApplicationCrash
	//
	// Simulation application exited abnormally.
	//
	// BadPermissionsRobotApplication
	//
	// Robot application bundle could not be downloaded.
	//
	// BadPermissionsSimulationApplication
	//
	// Simulation application bundle could not be downloaded.
	//
	// BadPermissionsS3Output
	//
	// Unable to publish outputs to customer-provided S3 bucket.
	//
	// BadPermissionsCloudwatchLogs
	//
	// Unable to publish logs to customer-provided CloudWatch Logs resource.
	//
	// SubnetIpLimitExceeded
	//
	// Subnet IP limit exceeded.
	//
	// ENILimitExceeded
	//
	// ENI limit exceeded.
	//
	// BadPermissionsUserCredentials
	//
	// Unable to use the Role provided.
	//
	// InvalidBundleRobotApplication
	//
	// Robot bundle cannot be extracted (invalid format, bundling error, or other
	// issue).
	//
	// InvalidBundleSimulationApplication
	//
	// Simulation bundle cannot be extracted (invalid format, bundling error, or
	// other issue).
	//
	// RobotApplicationVersionMismatchedEtag
	//
	// Etag for RobotApplication does not match value during version creation.
	//
	// SimulationApplicationVersionMismatchedEtag
	//
	// Etag for SimulationApplication does not match value during version creation.
	FailureCode DeploymentJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// The failure reason if the job fails.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// The status of the synchronization job.
	Status DeploymentStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s SyncDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SyncDeploymentJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.DeploymentApplicationConfigs != nil {
		v := s.DeploymentApplicationConfigs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "deploymentApplicationConfigs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DeploymentConfig != nil {
		v := s.DeploymentConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "deploymentConfig", v, metadata)
	}
	if len(s.FailureCode) > 0 {
		v := s.FailureCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FailureReason != nil {
		v := *s.FailureReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opSyncDeploymentJob = "SyncDeploymentJob"

// SyncDeploymentJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Syncrhonizes robots in a fleet to the latest deployment. This is helpful
// if robots were added after a deployment.
//
//    // Example sending a request using SyncDeploymentJobRequest.
//    req := client.SyncDeploymentJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/SyncDeploymentJob
func (c *Client) SyncDeploymentJobRequest(input *SyncDeploymentJobInput) SyncDeploymentJobRequest {
	op := &aws.Operation{
		Name:       opSyncDeploymentJob,
		HTTPMethod: "POST",
		HTTPPath:   "/syncDeploymentJob",
	}

	if input == nil {
		input = &SyncDeploymentJobInput{}
	}

	req := c.newRequest(op, input, &SyncDeploymentJobOutput{})
	return SyncDeploymentJobRequest{Request: req, Input: input, Copy: c.SyncDeploymentJobRequest}
}

// SyncDeploymentJobRequest is the request type for the
// SyncDeploymentJob API operation.
type SyncDeploymentJobRequest struct {
	*aws.Request
	Input *SyncDeploymentJobInput
	Copy  func(*SyncDeploymentJobInput) SyncDeploymentJobRequest
}

// Send marshals and sends the SyncDeploymentJob API request.
func (r SyncDeploymentJobRequest) Send(ctx context.Context) (*SyncDeploymentJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SyncDeploymentJobResponse{
		SyncDeploymentJobOutput: r.Request.Data.(*SyncDeploymentJobOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SyncDeploymentJobResponse is the response type for the
// SyncDeploymentJob API operation.
type SyncDeploymentJobResponse struct {
	*SyncDeploymentJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SyncDeploymentJob request.
func (r *SyncDeploymentJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
