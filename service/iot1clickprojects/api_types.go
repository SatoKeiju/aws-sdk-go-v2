// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// An object representing a device for a placement template (see PlacementTemplate).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/DeviceTemplate
type DeviceTemplate struct {
	_ struct{} `type:"structure"`

	// An optional Lambda function to invoke instead of the default Lambda function
	// provided by the placement template.
	CallbackOverrides map[string]string `locationName:"callbackOverrides" type:"map"`

	// The device type, which currently must be "button".
	DeviceType *string `locationName:"deviceType" type:"string"`
}

// String returns the string representation
func (s DeviceTemplate) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeviceTemplate) MarshalFields(e protocol.FieldEncoder) error {
	if s.CallbackOverrides != nil {
		v := s.CallbackOverrides

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "callbackOverrides", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.DeviceType != nil {
		v := *s.DeviceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object describing a project's placement.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/PlacementDescription
type PlacementDescription struct {
	_ struct{} `type:"structure"`

	// The user-defined attributes associated with the placement.
	//
	// Attributes is a required field
	Attributes map[string]string `locationName:"attributes" type:"map" required:"true"`

	// The date when the placement was initially created, in UNIX epoch time format.
	//
	// CreatedDate is a required field
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The name of the placement.
	//
	// PlacementName is a required field
	PlacementName *string `locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project containing the placement.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// The date when the placement was last updated, in UNIX epoch time format.
	// If the placement was not updated, then createdDate and updatedDate are the
	// same.
	//
	// UpdatedDate is a required field
	UpdatedDate *time.Time `locationName:"updatedDate" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s PlacementDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PlacementDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.PlacementName != nil {
		v := *s.PlacementName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "placementName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdatedDate != nil {
		v := *s.UpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

// An object providing summary information for a particular placement.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/PlacementSummary
type PlacementSummary struct {
	_ struct{} `type:"structure"`

	// The date when the placement was originally created, in UNIX epoch time format.
	//
	// CreatedDate is a required field
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The name of the placement being summarized.
	//
	// PlacementName is a required field
	PlacementName *string `locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project containing the placement.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// The date when the placement was last updated, in UNIX epoch time format.
	// If the placement was not updated, then createdDate and updatedDate are the
	// same.
	//
	// UpdatedDate is a required field
	UpdatedDate *time.Time `locationName:"updatedDate" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s PlacementSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PlacementSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.PlacementName != nil {
		v := *s.PlacementName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "placementName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdatedDate != nil {
		v := *s.UpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

// An object defining the template for a placement.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/PlacementTemplate
type PlacementTemplate struct {
	_ struct{} `type:"structure"`

	// The default attributes (key/value pairs) to be applied to all placements
	// using this template.
	DefaultAttributes map[string]string `locationName:"defaultAttributes" type:"map"`

	// An object specifying the DeviceTemplate for all placements using this (PlacementTemplate)
	// template.
	DeviceTemplates map[string]DeviceTemplate `locationName:"deviceTemplates" type:"map"`
}

// String returns the string representation
func (s PlacementTemplate) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PlacementTemplate) MarshalFields(e protocol.FieldEncoder) error {
	if s.DefaultAttributes != nil {
		v := s.DefaultAttributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "defaultAttributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.DeviceTemplates != nil {
		v := s.DeviceTemplates

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "deviceTemplates", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	return nil
}

// An object providing detailed information for a particular project associated
// with an AWS account and region.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/ProjectDescription
type ProjectDescription struct {
	_ struct{} `type:"structure"`

	// The ARN of the project.
	Arn *string `locationName:"arn" type:"string"`

	// The date when the project was originally created, in UNIX epoch time format.
	//
	// CreatedDate is a required field
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The description of the project.
	Description *string `locationName:"description" type:"string"`

	// An object describing the project's placement specifications.
	PlacementTemplate *PlacementTemplate `locationName:"placementTemplate" type:"structure"`

	// The name of the project for which to obtain information from.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// The tags (metadata key/value pairs) associated with the project.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`

	// The date when the project was last updated, in UNIX epoch time format. If
	// the project was not updated, then createdDate and updatedDate are the same.
	//
	// UpdatedDate is a required field
	UpdatedDate *time.Time `locationName:"updatedDate" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s ProjectDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ProjectDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlacementTemplate != nil {
		v := s.PlacementTemplate

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "placementTemplate", v, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.UpdatedDate != nil {
		v := *s.UpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

// An object providing summary information for a particular project for an associated
// AWS account and region.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/ProjectSummary
type ProjectSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the project.
	Arn *string `locationName:"arn" type:"string"`

	// The date when the project was originally created, in UNIX epoch time format.
	//
	// CreatedDate is a required field
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The name of the project being summarized.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// The tags (metadata key/value pairs) associated with the project.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`

	// The date when the project was last updated, in UNIX epoch time format. If
	// the project was not updated, then createdDate and updatedDate are the same.
	//
	// UpdatedDate is a required field
	UpdatedDate *time.Time `locationName:"updatedDate" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s ProjectSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ProjectSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.UpdatedDate != nil {
		v := *s.UpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}
