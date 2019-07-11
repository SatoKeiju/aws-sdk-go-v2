// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateAccountAuditConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Specifies which audit checks are enabled and disabled for this account. Use
	// DescribeAccountAuditConfiguration to see the list of all checks including
	// those that are currently enabled.
	//
	// Note that some data collection may begin immediately when certain checks
	// are enabled. When a check is disabled, any data collected so far in relation
	// to the check is deleted.
	//
	// You cannot disable a check if it is used by any scheduled audit. You must
	// first delete the check from the scheduled audit or delete the scheduled audit
	// itself.
	//
	// On the first call to UpdateAccountAuditConfiguration this parameter is required
	// and must specify at least one enabled check.
	AuditCheckConfigurations map[string]AuditCheckConfiguration `locationName:"auditCheckConfigurations" type:"map"`

	// Information about the targets to which audit notifications are sent.
	AuditNotificationTargetConfigurations map[string]AuditNotificationTarget `locationName:"auditNotificationTargetConfigurations" type:"map"`

	// The ARN of the role that grants permission to AWS IoT to access information
	// about your devices, policies, certificates and other items as necessary when
	// performing an audit.
	RoleArn *string `locationName:"roleArn" min:"20" type:"string"`
}

// String returns the string representation
func (s UpdateAccountAuditConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAccountAuditConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAccountAuditConfigurationInput"}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.AuditNotificationTargetConfigurations != nil {
		for i, v := range s.AuditNotificationTargetConfigurations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AuditNotificationTargetConfigurations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountAuditConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AuditCheckConfigurations != nil {
		v := s.AuditCheckConfigurations

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "auditCheckConfigurations", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.AuditNotificationTargetConfigurations != nil {
		v := s.AuditNotificationTargetConfigurations

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "auditNotificationTargetConfigurations", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateAccountAuditConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAccountAuditConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountAuditConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateAccountAuditConfiguration = "UpdateAccountAuditConfiguration"

// UpdateAccountAuditConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Configures or reconfigures the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks
// are enabled or disabled.
//
//    // Example sending a request using UpdateAccountAuditConfigurationRequest.
//    req := client.UpdateAccountAuditConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateAccountAuditConfigurationRequest(input *UpdateAccountAuditConfigurationInput) UpdateAccountAuditConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateAccountAuditConfiguration,
		HTTPMethod: "PATCH",
		HTTPPath:   "/audit/configuration",
	}

	if input == nil {
		input = &UpdateAccountAuditConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateAccountAuditConfigurationOutput{})
	return UpdateAccountAuditConfigurationRequest{Request: req, Input: input, Copy: c.UpdateAccountAuditConfigurationRequest}
}

// UpdateAccountAuditConfigurationRequest is the request type for the
// UpdateAccountAuditConfiguration API operation.
type UpdateAccountAuditConfigurationRequest struct {
	*aws.Request
	Input *UpdateAccountAuditConfigurationInput
	Copy  func(*UpdateAccountAuditConfigurationInput) UpdateAccountAuditConfigurationRequest
}

// Send marshals and sends the UpdateAccountAuditConfiguration API request.
func (r UpdateAccountAuditConfigurationRequest) Send(ctx context.Context) (*UpdateAccountAuditConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccountAuditConfigurationResponse{
		UpdateAccountAuditConfigurationOutput: r.Request.Data.(*UpdateAccountAuditConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccountAuditConfigurationResponse is the response type for the
// UpdateAccountAuditConfiguration API operation.
type UpdateAccountAuditConfigurationResponse struct {
	*UpdateAccountAuditConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccountAuditConfiguration request.
func (r *UpdateAccountAuditConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
