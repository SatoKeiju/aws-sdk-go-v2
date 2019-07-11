// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Unrchive Findings Request
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UnarchiveFindingsRequest
type UnarchiveFindingsInput struct {
	_ struct{} `type:"structure"`

	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" type:"string" required:"true"`

	// IDs of the findings that you want to unarchive.
	//
	// FindingIds is a required field
	FindingIds []string `locationName:"findingIds" type:"list" required:"true"`
}

// String returns the string representation
func (s UnarchiveFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnarchiveFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnarchiveFindingsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}

	if s.FindingIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UnarchiveFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FindingIds != nil {
		v := s.FindingIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findingIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UnarchiveFindingsResponse
type UnarchiveFindingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UnarchiveFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UnarchiveFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUnarchiveFindings = "UnarchiveFindings"

// UnarchiveFindingsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Unarchives Amazon GuardDuty findings specified by the list of finding IDs.
//
//    // Example sending a request using UnarchiveFindingsRequest.
//    req := client.UnarchiveFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UnarchiveFindings
func (c *Client) UnarchiveFindingsRequest(input *UnarchiveFindingsInput) UnarchiveFindingsRequest {
	op := &aws.Operation{
		Name:       opUnarchiveFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/findings/unarchive",
	}

	if input == nil {
		input = &UnarchiveFindingsInput{}
	}

	req := c.newRequest(op, input, &UnarchiveFindingsOutput{})
	return UnarchiveFindingsRequest{Request: req, Input: input, Copy: c.UnarchiveFindingsRequest}
}

// UnarchiveFindingsRequest is the request type for the
// UnarchiveFindings API operation.
type UnarchiveFindingsRequest struct {
	*aws.Request
	Input *UnarchiveFindingsInput
	Copy  func(*UnarchiveFindingsInput) UnarchiveFindingsRequest
}

// Send marshals and sends the UnarchiveFindings API request.
func (r UnarchiveFindingsRequest) Send(ctx context.Context) (*UnarchiveFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UnarchiveFindingsResponse{
		UnarchiveFindingsOutput: r.Request.Data.(*UnarchiveFindingsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UnarchiveFindingsResponse is the response type for the
// UnarchiveFindings API operation.
type UnarchiveFindingsResponse struct {
	*UnarchiveFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UnarchiveFindings request.
func (r *UnarchiveFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
