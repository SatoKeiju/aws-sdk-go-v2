// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The updates that you want to make to a specific entitlement.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/UpdateFlowEntitlementRequest
type UpdateFlowEntitlementInput struct {
	_ struct{} `type:"structure"`

	// A description of the entitlement. This description appears only on the AWS
	// Elemental MediaConnect console and will not be seen by the subscriber or
	// end user.
	Description *string `locationName:"description" type:"string"`

	// The type of encryption that will be used on the output associated with this
	// entitlement.
	Encryption *UpdateEncryption `locationName:"encryption" type:"structure"`

	// EntitlementArn is a required field
	EntitlementArn *string `location:"uri" locationName:"entitlementArn" type:"string" required:"true"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// The AWS account IDs that you want to share your content with. The receiving
	// accounts (subscribers) will be allowed to create their own flow using your
	// content as the source.
	Subscribers []string `locationName:"subscribers" type:"list"`
}

// String returns the string representation
func (s UpdateFlowEntitlementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFlowEntitlementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFlowEntitlementInput"}

	if s.EntitlementArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntitlementArn"))
	}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFlowEntitlementInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Encryption != nil {
		v := s.Encryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryption", v, metadata)
	}
	if s.Subscribers != nil {
		v := s.Subscribers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subscribers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.EntitlementArn != nil {
		v := *s.EntitlementArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "entitlementArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful UpdateFlowEntitlement request. The response includes
// the ARN of the flow that was updated and the updated entitlement configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/UpdateFlowEntitlementResponse
type UpdateFlowEntitlementOutput struct {
	_ struct{} `type:"structure"`

	// The settings for a flow entitlement.
	Entitlement *Entitlement `locationName:"entitlement" type:"structure"`

	// The ARN of the flow that this entitlement was granted on.
	FlowArn *string `locationName:"flowArn" type:"string"`
}

// String returns the string representation
func (s UpdateFlowEntitlementOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFlowEntitlementOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Entitlement != nil {
		v := s.Entitlement

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "entitlement", v, metadata)
	}
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateFlowEntitlement = "UpdateFlowEntitlement"

// UpdateFlowEntitlementRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// You can change an entitlement's description, subscribers, and encryption.
// If you change the subscribers, the service will remove the outputs that are
// are used by the subscribers that are removed.
//
//    // Example sending a request using UpdateFlowEntitlementRequest.
//    req := client.UpdateFlowEntitlementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/UpdateFlowEntitlement
func (c *Client) UpdateFlowEntitlementRequest(input *UpdateFlowEntitlementInput) UpdateFlowEntitlementRequest {
	op := &aws.Operation{
		Name:       opUpdateFlowEntitlement,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/flows/{flowArn}/entitlements/{entitlementArn}",
	}

	if input == nil {
		input = &UpdateFlowEntitlementInput{}
	}

	req := c.newRequest(op, input, &UpdateFlowEntitlementOutput{})
	return UpdateFlowEntitlementRequest{Request: req, Input: input, Copy: c.UpdateFlowEntitlementRequest}
}

// UpdateFlowEntitlementRequest is the request type for the
// UpdateFlowEntitlement API operation.
type UpdateFlowEntitlementRequest struct {
	*aws.Request
	Input *UpdateFlowEntitlementInput
	Copy  func(*UpdateFlowEntitlementInput) UpdateFlowEntitlementRequest
}

// Send marshals and sends the UpdateFlowEntitlement API request.
func (r UpdateFlowEntitlementRequest) Send(ctx context.Context) (*UpdateFlowEntitlementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFlowEntitlementResponse{
		UpdateFlowEntitlementOutput: r.Request.Data.(*UpdateFlowEntitlementOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFlowEntitlementResponse is the response type for the
// UpdateFlowEntitlement API operation.
type UpdateFlowEntitlementResponse struct {
	*UpdateFlowEntitlementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFlowEntitlement request.
func (r *UpdateFlowEntitlementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
