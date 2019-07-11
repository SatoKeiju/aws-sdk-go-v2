// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request a new generated client SDK for a RestApi and Stage.
type GetSdkInput struct {
	_ struct{} `type:"structure"`

	// A string-to-string key-value map of query parameters sdkType-dependent properties
	// of the SDK. For sdkType of objectivec or swift, a parameter named classPrefix
	// is required. For sdkType of android, parameters named groupId, artifactId,
	// artifactVersion, and invokerPackage are required. For sdkType of java, parameters
	// named serviceName and javaPackageName are required.
	Parameters map[string]string `location:"querystring" locationName:"parameters" type:"map"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The language for the generated SDK. Currently java, javascript,
	// android, objectivec (for iOS), swift (for iOS), and ruby are supported.
	//
	// SdkType is a required field
	SdkType *string `location:"uri" locationName:"sdk_type" type:"string" required:"true"`

	// [Required] The name of the Stage that the SDK will use.
	//
	// StageName is a required field
	StageName *string `location:"uri" locationName:"stage_name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSdkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSdkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSdkInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.SdkType == nil {
		invalidParams.Add(aws.NewErrParamRequired("SdkType"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSdkInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SdkType != nil {
		v := *s.SdkType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "sdk_type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "stage_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.QueryTarget, "parameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// The binary blob response to GetSdk, which contains the generated SDK.
type GetSdkOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The binary blob response to GetSdk, which contains the generated SDK.
	Body []byte `locationName:"body" type:"blob"`

	// The content-disposition header value in the HTTP response.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// The content-type header value in the HTTP response.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
}

// String returns the string representation
func (s GetSdkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSdkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetSdk = "GetSdk"

// GetSdkRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Generates a client SDK for a RestApi and Stage.
//
//    // Example sending a request using GetSdkRequest.
//    req := client.GetSdkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetSdkRequest(input *GetSdkInput) GetSdkRequest {
	op := &aws.Operation{
		Name:       opGetSdk,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}/sdks/{sdk_type}",
	}

	if input == nil {
		input = &GetSdkInput{}
	}

	req := c.newRequest(op, input, &GetSdkOutput{})
	return GetSdkRequest{Request: req, Input: input, Copy: c.GetSdkRequest}
}

// GetSdkRequest is the request type for the
// GetSdk API operation.
type GetSdkRequest struct {
	*aws.Request
	Input *GetSdkInput
	Copy  func(*GetSdkInput) GetSdkRequest
}

// Send marshals and sends the GetSdk API request.
func (r GetSdkRequest) Send(ctx context.Context) (*GetSdkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSdkResponse{
		GetSdkOutput: r.Request.Data.(*GetSdkOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSdkResponse is the response type for the
// GetSdk API operation.
type GetSdkResponse struct {
	*GetSdkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSdk request.
func (r *GetSdkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
