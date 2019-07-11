// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/UpdateConfigurationRequest
type UpdateConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ConfigurationId is a required field
	ConfigurationId *string `location:"uri" locationName:"configuration-id" type:"string" required:"true"`

	Data *string `locationName:"data" type:"string"`

	Description *string `locationName:"description" type:"string"`
}

// String returns the string representation
func (s UpdateConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationInput"}

	if s.ConfigurationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Data != nil {
		v := *s.Data

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "data", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationId != nil {
		v := *s.ConfigurationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "configuration-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/UpdateConfigurationResponse
type UpdateConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Created *time.Time `locationName:"created" type:"timestamp" timestampFormat:"unix"`

	Id *string `locationName:"id" type:"string"`

	// Returns information about the specified configuration revision.
	LatestRevision *ConfigurationRevision `locationName:"latestRevision" type:"structure"`

	Name *string `locationName:"name" type:"string"`

	Warnings []SanitizationWarning `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s UpdateConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestRevision != nil {
		v := s.LatestRevision

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "latestRevision", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opUpdateConfiguration = "UpdateConfiguration"

// UpdateConfigurationRequest returns a request value for making API operation for
// AmazonMQ.
//
// Updates the specified configuration.
//
//    // Example sending a request using UpdateConfigurationRequest.
//    req := client.UpdateConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/UpdateConfiguration
func (c *Client) UpdateConfigurationRequest(input *UpdateConfigurationInput) UpdateConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/configurations/{configuration-id}",
	}

	if input == nil {
		input = &UpdateConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateConfigurationOutput{})
	return UpdateConfigurationRequest{Request: req, Input: input, Copy: c.UpdateConfigurationRequest}
}

// UpdateConfigurationRequest is the request type for the
// UpdateConfiguration API operation.
type UpdateConfigurationRequest struct {
	*aws.Request
	Input *UpdateConfigurationInput
	Copy  func(*UpdateConfigurationInput) UpdateConfigurationRequest
}

// Send marshals and sends the UpdateConfiguration API request.
func (r UpdateConfigurationRequest) Send(ctx context.Context) (*UpdateConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConfigurationResponse{
		UpdateConfigurationOutput: r.Request.Data.(*UpdateConfigurationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConfigurationResponse is the response type for the
// UpdateConfiguration API operation.
type UpdateConfigurationResponse struct {
	*UpdateConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConfiguration request.
func (r *UpdateConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
