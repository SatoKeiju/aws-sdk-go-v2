// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/UpdateFileSystemRequest
type UpdateFileSystemInput struct {
	_ struct{} `type:"structure"`

	// The ID of the file system that you want to update.
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	// (Optional) The amount of throughput, in MiB/s, that you want to provision
	// for your file system. Valid values are 1-1024. Required if ThroughputMode
	// is changed to provisioned on update. If you're not updating the amount of
	// provisioned throughput for your file system, you don't need to provide this
	// value in your request.
	ProvisionedThroughputInMibps *float64 `min:"1" type:"double"`

	// (Optional) The throughput mode that you want your file system to use. If
	// you're not updating your throughput mode, you don't need to provide this
	// value in your request. If you are changing the ThroughputMode to provisioned,
	// you must also set a value for ProvisionedThroughputInMibps.
	ThroughputMode ThroughputMode `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateFileSystemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFileSystemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFileSystemInput"}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}
	if s.ProvisionedThroughputInMibps != nil && *s.ProvisionedThroughputInMibps < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ProvisionedThroughputInMibps", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFileSystemInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ProvisionedThroughputInMibps != nil {
		v := *s.ProvisionedThroughputInMibps

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProvisionedThroughputInMibps", protocol.Float64Value(v), metadata)
	}
	if len(s.ThroughputMode) > 0 {
		v := s.ThroughputMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThroughputMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A description of the file system.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/FileSystemDescription
type UpdateFileSystemOutput struct {
	_ struct{} `type:"structure"`

	// The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z).
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The opaque string specified in the request.
	//
	// CreationToken is a required field
	CreationToken *string `min:"1" type:"string" required:"true"`

	// A Boolean value that, if true, indicates that the file system is encrypted.
	Encrypted *bool `type:"boolean"`

	// The ID of the file system, assigned by Amazon EFS.
	//
	// FileSystemId is a required field
	FileSystemId *string `type:"string" required:"true"`

	// The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK)
	// that was used to protect the encrypted file system.
	KmsKeyId *string `min:"1" type:"string"`

	// The lifecycle phase of the file system.
	//
	// LifeCycleState is a required field
	LifeCycleState LifeCycleState `type:"string" required:"true" enum:"true"`

	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem. If the file system has a Name tag, Amazon EFS returns
	// the value in this field.
	Name *string `type:"string"`

	// The current number of mount targets that the file system has. For more information,
	// see CreateMountTarget.
	//
	// NumberOfMountTargets is a required field
	NumberOfMountTargets *int64 `type:"integer" required:"true"`

	// The AWS account that created the file system. If the file system was created
	// by an IAM user, the parent account to which the user belongs is the owner.
	//
	// OwnerId is a required field
	OwnerId *string `type:"string" required:"true"`

	// The performance mode of the file system.
	//
	// PerformanceMode is a required field
	PerformanceMode PerformanceMode `type:"string" required:"true" enum:"true"`

	// The throughput, measured in MiB/s, that you want to provision for a file
	// system. Valid values are 1-1024. Required if ThroughputMode is set to provisioned.
	// The limit on throughput is 1024 MiB/s. You can get these limits increased
	// by contacting AWS Support. For more information, see Amazon EFS Limits That
	// You Can Increase (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits)
	// in the Amazon EFS User Guide.
	ProvisionedThroughputInMibps *float64 `min:"1" type:"double"`

	// The latest known metered size (in bytes) of data stored in the file system,
	// in its Value field, and the time at which that size was determined in its
	// Timestamp field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of
	// a consistent snapshot of the file system, but it is eventually consistent
	// when there are no writes to the file system. That is, SizeInBytes represents
	// actual size only if the file system is not modified for a period longer than
	// a couple of hours. Otherwise, the value is not the exact size that the file
	// system was at any point in time.
	//
	// SizeInBytes is a required field
	SizeInBytes *FileSystemSize `type:"structure" required:"true"`

	// The tags associated with the file system, presented as an array of Tag objects.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`

	// The throughput mode for a file system. There are two throughput modes to
	// choose from for your file system: bursting and provisioned. If you set ThroughputMode
	// to provisioned, you must also set a value for ProvisionedThroughPutInMibps.
	// You can decrease your file system's throughput in Provisioned Throughput
	// mode or change between the throughput modes as long as it’s been more than
	// 24 hours since the last decrease or throughput mode change.
	ThroughputMode ThroughputMode `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateFileSystemOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFileSystemOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.CreationToken != nil {
		v := *s.CreationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Encrypted != nil {
		v := *s.Encrypted

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Encrypted", protocol.BoolValue(v), metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KmsKeyId != nil {
		v := *s.KmsKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KmsKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LifeCycleState) > 0 {
		v := s.LifeCycleState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LifeCycleState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NumberOfMountTargets != nil {
		v := *s.NumberOfMountTargets

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NumberOfMountTargets", protocol.Int64Value(v), metadata)
	}
	if s.OwnerId != nil {
		v := *s.OwnerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OwnerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PerformanceMode) > 0 {
		v := s.PerformanceMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PerformanceMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ProvisionedThroughputInMibps != nil {
		v := *s.ProvisionedThroughputInMibps

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProvisionedThroughputInMibps", protocol.Float64Value(v), metadata)
	}
	if s.SizeInBytes != nil {
		v := s.SizeInBytes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SizeInBytes", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ThroughputMode) > 0 {
		v := s.ThroughputMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThroughputMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opUpdateFileSystem = "UpdateFileSystem"

// UpdateFileSystemRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Updates the throughput mode or the amount of provisioned throughput of an
// existing file system.
//
//    // Example sending a request using UpdateFileSystemRequest.
//    req := client.UpdateFileSystemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/UpdateFileSystem
func (c *Client) UpdateFileSystemRequest(input *UpdateFileSystemInput) UpdateFileSystemRequest {
	op := &aws.Operation{
		Name:       opUpdateFileSystem,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}",
	}

	if input == nil {
		input = &UpdateFileSystemInput{}
	}

	req := c.newRequest(op, input, &UpdateFileSystemOutput{})
	return UpdateFileSystemRequest{Request: req, Input: input, Copy: c.UpdateFileSystemRequest}
}

// UpdateFileSystemRequest is the request type for the
// UpdateFileSystem API operation.
type UpdateFileSystemRequest struct {
	*aws.Request
	Input *UpdateFileSystemInput
	Copy  func(*UpdateFileSystemInput) UpdateFileSystemRequest
}

// Send marshals and sends the UpdateFileSystem API request.
func (r UpdateFileSystemRequest) Send(ctx context.Context) (*UpdateFileSystemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFileSystemResponse{
		UpdateFileSystemOutput: r.Request.Data.(*UpdateFileSystemOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFileSystemResponse is the response type for the
// UpdateFileSystem API operation.
type UpdateFileSystemResponse struct {
	*UpdateFileSystemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFileSystem request.
func (r *UpdateFileSystemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
