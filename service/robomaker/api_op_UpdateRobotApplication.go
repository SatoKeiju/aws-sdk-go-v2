// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/UpdateRobotApplicationRequest
type UpdateRobotApplicationInput struct {
	_ struct{} `type:"structure"`

	// The application information for the robot application.
	//
	// Application is a required field
	Application *string `locationName:"application" min:"1" type:"string" required:"true"`

	// The revision id for the robot application.
	CurrentRevisionId *string `locationName:"currentRevisionId" min:"1" type:"string"`

	// The robot software suite used by the robot application.
	//
	// RobotSoftwareSuite is a required field
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure" required:"true"`

	// The sources of the robot application.
	//
	// Sources is a required field
	Sources []SourceConfig `locationName:"sources" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateRobotApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRobotApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRobotApplicationInput"}

	if s.Application == nil {
		invalidParams.Add(aws.NewErrParamRequired("Application"))
	}
	if s.Application != nil && len(*s.Application) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Application", 1))
	}
	if s.CurrentRevisionId != nil && len(*s.CurrentRevisionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentRevisionId", 1))
	}

	if s.RobotSoftwareSuite == nil {
		invalidParams.Add(aws.NewErrParamRequired("RobotSoftwareSuite"))
	}

	if s.Sources == nil {
		invalidParams.Add(aws.NewErrParamRequired("Sources"))
	}
	if s.Sources != nil {
		for i, v := range s.Sources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Sources", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRobotApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Application != nil {
		v := *s.Application

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "application", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CurrentRevisionId != nil {
		v := *s.CurrentRevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentRevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotSoftwareSuite != nil {
		v := s.RobotSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "robotSoftwareSuite", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/UpdateRobotApplicationResponse
type UpdateRobotApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the updated robot application.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the robot application was
	// last updated.
	LastUpdatedAt *time.Time `locationName:"lastUpdatedAt" type:"timestamp" timestampFormat:"unix"`

	// The name of the robot application.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The revision id of the robot application.
	RevisionId *string `locationName:"revisionId" min:"1" type:"string"`

	// The robot software suite used by the robot application.
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure"`

	// The sources of the robot application.
	Sources []Source `locationName:"sources" type:"list"`

	// The version of the robot application.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateRobotApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRobotApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotSoftwareSuite != nil {
		v := s.RobotSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "robotSoftwareSuite", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateRobotApplication = "UpdateRobotApplication"

// UpdateRobotApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Updates a robot application.
//
//    // Example sending a request using UpdateRobotApplicationRequest.
//    req := client.UpdateRobotApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/UpdateRobotApplication
func (c *Client) UpdateRobotApplicationRequest(input *UpdateRobotApplicationInput) UpdateRobotApplicationRequest {
	op := &aws.Operation{
		Name:       opUpdateRobotApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/updateRobotApplication",
	}

	if input == nil {
		input = &UpdateRobotApplicationInput{}
	}

	req := c.newRequest(op, input, &UpdateRobotApplicationOutput{})
	return UpdateRobotApplicationRequest{Request: req, Input: input, Copy: c.UpdateRobotApplicationRequest}
}

// UpdateRobotApplicationRequest is the request type for the
// UpdateRobotApplication API operation.
type UpdateRobotApplicationRequest struct {
	*aws.Request
	Input *UpdateRobotApplicationInput
	Copy  func(*UpdateRobotApplicationInput) UpdateRobotApplicationRequest
}

// Send marshals and sends the UpdateRobotApplication API request.
func (r UpdateRobotApplicationRequest) Send(ctx context.Context) (*UpdateRobotApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRobotApplicationResponse{
		UpdateRobotApplicationOutput: r.Request.Data.(*UpdateRobotApplicationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRobotApplicationResponse is the response type for the
// UpdateRobotApplication API operation.
type UpdateRobotApplicationResponse struct {
	*UpdateRobotApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRobotApplication request.
func (r *UpdateRobotApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
