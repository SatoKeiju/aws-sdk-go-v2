// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the authorizer's Lambda function.
	//
	// AuthorizerFunctionArn is a required field
	AuthorizerFunctionArn *string `locationName:"authorizerFunctionArn" type:"string" required:"true"`

	// The authorizer name.
	//
	// AuthorizerName is a required field
	AuthorizerName *string `location:"uri" locationName:"authorizerName" min:"1" type:"string" required:"true"`

	// The status of the create authorizer request.
	Status AuthorizerStatus `locationName:"status" type:"string" enum:"true"`

	// The name of the token key used to extract the token from the HTTP headers.
	//
	// TokenKeyName is a required field
	TokenKeyName *string `locationName:"tokenKeyName" min:"1" type:"string" required:"true"`

	// The public keys used to verify the digital signature returned by your custom
	// authentication service.
	//
	// TokenSigningPublicKeys is a required field
	TokenSigningPublicKeys map[string]string `locationName:"tokenSigningPublicKeys" type:"map" required:"true"`
}

// String returns the string representation
func (s CreateAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAuthorizerInput"}

	if s.AuthorizerFunctionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerFunctionArn"))
	}

	if s.AuthorizerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerName"))
	}
	if s.AuthorizerName != nil && len(*s.AuthorizerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizerName", 1))
	}

	if s.TokenKeyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TokenKeyName"))
	}
	if s.TokenKeyName != nil && len(*s.TokenKeyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TokenKeyName", 1))
	}

	if s.TokenSigningPublicKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TokenSigningPublicKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AuthorizerFunctionArn != nil {
		v := *s.AuthorizerFunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerFunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TokenKeyName != nil {
		v := *s.TokenKeyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "tokenKeyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TokenSigningPublicKeys != nil {
		v := s.TokenSigningPublicKeys

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tokenSigningPublicKeys", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.AuthorizerName != nil {
		v := *s.AuthorizerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "authorizerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	// The authorizer ARN.
	AuthorizerArn *string `locationName:"authorizerArn" type:"string"`

	// The authorizer's name.
	AuthorizerName *string `locationName:"authorizerName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthorizerArn != nil {
		v := *s.AuthorizerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerName != nil {
		v := *s.AuthorizerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateAuthorizer = "CreateAuthorizer"

// CreateAuthorizerRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates an authorizer.
//
//    // Example sending a request using CreateAuthorizerRequest.
//    req := client.CreateAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateAuthorizerRequest(input *CreateAuthorizerInput) CreateAuthorizerRequest {
	op := &aws.Operation{
		Name:       opCreateAuthorizer,
		HTTPMethod: "POST",
		HTTPPath:   "/authorizer/{authorizerName}",
	}

	if input == nil {
		input = &CreateAuthorizerInput{}
	}

	req := c.newRequest(op, input, &CreateAuthorizerOutput{})
	return CreateAuthorizerRequest{Request: req, Input: input, Copy: c.CreateAuthorizerRequest}
}

// CreateAuthorizerRequest is the request type for the
// CreateAuthorizer API operation.
type CreateAuthorizerRequest struct {
	*aws.Request
	Input *CreateAuthorizerInput
	Copy  func(*CreateAuthorizerInput) CreateAuthorizerRequest
}

// Send marshals and sends the CreateAuthorizer API request.
func (r CreateAuthorizerRequest) Send(ctx context.Context) (*CreateAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAuthorizerResponse{
		CreateAuthorizerOutput: r.Request.Data.(*CreateAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAuthorizerResponse is the response type for the
// CreateAuthorizer API operation.
type CreateAuthorizerResponse struct {
	*CreateAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAuthorizer request.
func (r *CreateAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
