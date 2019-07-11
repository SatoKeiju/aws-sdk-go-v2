// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDomainNameRequest
type GetDomainNameInput struct {
	_ struct{} `type:"structure"`

	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainNameInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainNameInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDomainNameResponse
type GetDomainNameOutput struct {
	_ struct{} `type:"structure"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiMappingSelectionExpression *string `locationName:"apiMappingSelectionExpression" type:"string"`

	// A string with a length between [1-512].
	DomainName *string `locationName:"domainName" type:"string"`

	// The domain name configurations.
	DomainNameConfigurations []DomainNameConfiguration `locationName:"domainNameConfigurations" type:"list"`
}

// String returns the string representation
func (s GetDomainNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainNameOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiMappingSelectionExpression != nil {
		v := *s.ApiMappingSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiMappingSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainNameConfigurations != nil {
		v := s.DomainNameConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "domainNameConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetDomainName = "GetDomainName"

// GetDomainNameRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a domain name.
//
//    // Example sending a request using GetDomainNameRequest.
//    req := client.GetDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDomainName
func (c *Client) GetDomainNameRequest(input *GetDomainNameInput) GetDomainNameRequest {
	op := &aws.Operation{
		Name:       opGetDomainName,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/domainnames/{domainName}",
	}

	if input == nil {
		input = &GetDomainNameInput{}
	}

	req := c.newRequest(op, input, &GetDomainNameOutput{})
	return GetDomainNameRequest{Request: req, Input: input, Copy: c.GetDomainNameRequest}
}

// GetDomainNameRequest is the request type for the
// GetDomainName API operation.
type GetDomainNameRequest struct {
	*aws.Request
	Input *GetDomainNameInput
	Copy  func(*GetDomainNameInput) GetDomainNameRequest
}

// Send marshals and sends the GetDomainName API request.
func (r GetDomainNameRequest) Send(ctx context.Context) (*GetDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainNameResponse{
		GetDomainNameOutput: r.Request.Data.(*GetDomainNameOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainNameResponse is the response type for the
// GetDomainName API operation.
type GetDomainNameResponse struct {
	*GetDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainName request.
func (r *GetDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
