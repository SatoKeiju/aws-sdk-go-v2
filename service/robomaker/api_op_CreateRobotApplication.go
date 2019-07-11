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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobotApplicationRequest
type CreateRobotApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the robot application.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The robot software suite used by the robot application.
	//
	// RobotSoftwareSuite is a required field
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure" required:"true"`

	// The sources of the robot application.
	//
	// Sources is a required field
	Sources []SourceConfig `locationName:"sources" type:"list" required:"true"`

	// A map that contains tag keys and tag values that are attached to the robot
	// application.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateRobotApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRobotApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRobotApplicationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
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
func (s CreateRobotApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobotApplicationResponse
type CreateRobotApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the robot application.
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

	// The list of all tags added to the robot application.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The version of the robot application.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateRobotApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRobotApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateRobotApplication = "CreateRobotApplication"

// CreateRobotApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Creates a robot application.
//
//    // Example sending a request using CreateRobotApplicationRequest.
//    req := client.CreateRobotApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobotApplication
func (c *Client) CreateRobotApplicationRequest(input *CreateRobotApplicationInput) CreateRobotApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateRobotApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/createRobotApplication",
	}

	if input == nil {
		input = &CreateRobotApplicationInput{}
	}

	req := c.newRequest(op, input, &CreateRobotApplicationOutput{})
	return CreateRobotApplicationRequest{Request: req, Input: input, Copy: c.CreateRobotApplicationRequest}
}

// CreateRobotApplicationRequest is the request type for the
// CreateRobotApplication API operation.
type CreateRobotApplicationRequest struct {
	*aws.Request
	Input *CreateRobotApplicationInput
	Copy  func(*CreateRobotApplicationInput) CreateRobotApplicationRequest
}

// Send marshals and sends the CreateRobotApplication API request.
func (r CreateRobotApplicationRequest) Send(ctx context.Context) (*CreateRobotApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRobotApplicationResponse{
		CreateRobotApplicationOutput: r.Request.Data.(*CreateRobotApplicationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRobotApplicationResponse is the response type for the
// CreateRobotApplication API operation.
type CreateRobotApplicationResponse struct {
	*CreateRobotApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRobotApplication request.
func (r *CreateRobotApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
