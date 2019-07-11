// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to begin the verification process for an email identity (an email
// address or domain).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateEmailIdentityRequest
type CreateEmailIdentityInput struct {
	_ struct{} `type:"structure"`

	// The email address or domain that you want to verify.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `type:"string" required:"true"`

	// An array of objects that define the tags (keys and values) that you want
	// to associate with the email identity.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateEmailIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEmailIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEmailIdentityInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateEmailIdentityInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.EmailIdentity != nil {
		v := *s.EmailIdentity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmailIdentity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

// If the email identity is a domain, this object contains tokens that you can
// use to create a set of CNAME records. To sucessfully verify your domain,
// you have to add these records to the DNS configuration for your domain.
//
// If the email identity is an email address, this object is empty.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateEmailIdentityResponse
type CreateEmailIdentityOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about the DKIM attributes for the identity.
	// This object includes the tokens that you use to create the CNAME records
	// that are required to complete the DKIM verification process.
	DkimAttributes *DkimAttributes `type:"structure"`

	// The email identity type.
	IdentityType IdentityType `type:"string" enum:"true"`

	// Specifies whether or not the identity is verified. In Amazon Pinpoint, you
	// can only send email from verified email addresses or domains. For more information
	// about verifying identities, see the Amazon Pinpoint User Guide (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus *bool `type:"boolean"`
}

// String returns the string representation
func (s CreateEmailIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateEmailIdentityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DkimAttributes != nil {
		v := s.DkimAttributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DkimAttributes", v, metadata)
	}
	if len(s.IdentityType) > 0 {
		v := s.IdentityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.VerifiedForSendingStatus != nil {
		v := *s.VerifiedForSendingStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VerifiedForSendingStatus", protocol.BoolValue(v), metadata)
	}
	return nil
}

const opCreateEmailIdentity = "CreateEmailIdentity"

// CreateEmailIdentityRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Verifies an email identity for use with Amazon Pinpoint. In Amazon Pinpoint,
// an identity is an email address or domain that you use when you send email.
// Before you can use an identity to send email with Amazon Pinpoint, you first
// have to verify it. By verifying an address, you demonstrate that you're the
// owner of the address, and that you've given Amazon Pinpoint permission to
// send email from the address.
//
// When you verify an email address, Amazon Pinpoint sends an email to the address.
// Your email address is verified as soon as you follow the link in the verification
// email.
//
// When you verify a domain, this operation provides a set of DKIM tokens, which
// you can convert into CNAME tokens. You add these CNAME tokens to the DNS
// configuration for your domain. Your domain is verified when Amazon Pinpoint
// detects these records in the DNS configuration for your domain. It usually
// takes around 72 hours to complete the domain verification process.
//
//    // Example sending a request using CreateEmailIdentityRequest.
//    req := client.CreateEmailIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateEmailIdentity
func (c *Client) CreateEmailIdentityRequest(input *CreateEmailIdentityInput) CreateEmailIdentityRequest {
	op := &aws.Operation{
		Name:       opCreateEmailIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/email/identities",
	}

	if input == nil {
		input = &CreateEmailIdentityInput{}
	}

	req := c.newRequest(op, input, &CreateEmailIdentityOutput{})
	return CreateEmailIdentityRequest{Request: req, Input: input, Copy: c.CreateEmailIdentityRequest}
}

// CreateEmailIdentityRequest is the request type for the
// CreateEmailIdentity API operation.
type CreateEmailIdentityRequest struct {
	*aws.Request
	Input *CreateEmailIdentityInput
	Copy  func(*CreateEmailIdentityInput) CreateEmailIdentityRequest
}

// Send marshals and sends the CreateEmailIdentity API request.
func (r CreateEmailIdentityRequest) Send(ctx context.Context) (*CreateEmailIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEmailIdentityResponse{
		CreateEmailIdentityOutput: r.Request.Data.(*CreateEmailIdentityOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEmailIdentityResponse is the response type for the
// CreateEmailIdentity API operation.
type CreateEmailIdentityResponse struct {
	*CreateEmailIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEmailIdentity request.
func (r *CreateEmailIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
