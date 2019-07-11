// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to describe an existing Method resource.
type GetMethodInput struct {
	_ struct{} `type:"structure"`

	// [Required] Specifies the method request's HTTP method type.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// [Required] The Resource identifier for the Method resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMethodInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMethodInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMethodInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMethodInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "http_method", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a client-facing interface by which the client calls the API to
// access back-end resources. A Method resource is integrated with an Integration
// resource. Both consist of a request and one or more responses. The method
// request takes the client input that is passed to the back end through the
// integration request. A method response returns the output from the back end
// to the client through an integration response. A method request is embodied
// in a Method resource, whereas an integration request is embodied in an Integration
// resource. On the other hand, a method response is represented by a MethodResponse
// resource, whereas an integration response is represented by an IntegrationResponse
// resource.
//
// Example: Retrive the GET method on a specified resource
//
// Request
//
// The following example request retrieves the information about the GET method
// on an API resource (3kzxbg5sa2) of an API (fugvjdxtri).
//  GET /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET HTTP/1.1 Content-Type:
//  application/json Host: apigateway.us-east-1.amazonaws.com X-Amz-Date: 20160603T210259Z
//  Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160603/us-east-1/apigateway/aws4_request,
//  SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
// Response
//
// The successful response returns a 200 OK status code and a payload similar
// to the following:
//  { "_links": { "curies": [ { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-{rel}.html",
//  "name": "integration", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
//  "name": "integrationresponse", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
//  "name": "method", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
//  "name": "methodresponse", "templated": true } ], "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET",
//  "name": "GET", "title": "GET" }, "integration:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
//  }, "method:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET"
//  }, "method:integration": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
//  }, "method:responses": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
//  "name": "200", "title": "200" }, "method:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET"
//  }, "methodresponse:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/{status_code}",
//  "templated": true } }, "apiKeyRequired": true, "authorizationType": "NONE",
//  "httpMethod": "GET", "_embedded": { "method:integration": { "_links": {
//  "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
//  }, "integration:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
//  }, "integration:responses": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
//  "name": "200", "title": "200" }, "integration:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration"
//  }, "integrationresponse:put": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/{status_code}",
//  "templated": true } }, "cacheKeyParameters": [], "cacheNamespace": "3kzxbg5sa2",
//  "credentials": "arn:aws:iam::123456789012:role/apigAwsProxyRole", "httpMethod":
//  "POST", "passthroughBehavior": "WHEN_NO_MATCH", "requestParameters": { "integration.request.header.Content-Type":
//  "'application/x-amz-json-1.1'" }, "requestTemplates": { "application/json":
//  "{\n}" }, "type": "AWS", "uri": "arn:aws:apigateway:us-east-1:kinesis:action/ListStreams",
//  "_embedded": { "integration:responses": { "_links": { "self": { "href":
//  "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
//  "name": "200", "title": "200" }, "integrationresponse:delete": { "href":
//  "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
//  }, "integrationresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
//  } }, "responseParameters": { "method.response.header.Content-Type": "'application/xml'"
//  }, "responseTemplates": { "application/json": "$util.urlDecode(\"%3CkinesisStreams%3E%23foreach(%24stream%20in%20%24input.path(%27%24.StreamNames%27))%3Cstream%3E%3Cname%3E%24stream%3C%2Fname%3E%3C%2Fstream%3E%23end%3C%2FkinesisStreams%3E\")"
//  }, "statusCode": "200" } } }, "method:responses": { "_links": { "self":
//  { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
//  "name": "200", "title": "200" }, "methodresponse:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
//  }, "methodresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
//  } }, "responseModels": { "application/json": "Empty" }, "responseParameters":
//  { "method.response.header.Content-Type": false }, "statusCode": "200" }
//  } }
// In the example above, the response template for the 200 OK response maps
// the JSON output from the ListStreams action in the back end to an XML output.
// The mapping template is URL-encoded as %3CkinesisStreams%3E%23foreach(%24stream%20in%20%24input.path(%27%24.StreamNames%27))%3Cstream%3E%3Cname%3E%24stream%3C%2Fname%3E%3C%2Fstream%3E%23end%3C%2FkinesisStreams%3E
// and the output is decoded using the $util.urlDecode() (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#util-templat-reference)
// helper function.
//
// MethodResponse, Integration, IntegrationResponse, Resource, Set up an API's
// method (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-method-settings.html)
type GetMethodOutput struct {
	_ struct{} `type:"structure"`

	// A boolean flag specifying whether a valid ApiKey is required to invoke this
	// method.
	ApiKeyRequired *bool `locationName:"apiKeyRequired" type:"boolean"`

	// A list of authorization scopes configured on the method. The scopes are used
	// with a COGNITO_USER_POOLS authorizer to authorize the method invocation.
	// The authorization works by matching the method scopes against the scopes
	// parsed from the access token in the incoming request. The method invocation
	// is authorized if any method scopes matches a claimed scope in the access
	// token. Otherwise, the invocation is not authorized. When the method scope
	// is configured, the client must provide an access token instead of an identity
	// token for authorization purposes.
	AuthorizationScopes []string `locationName:"authorizationScopes" type:"list"`

	// The method's authorization type. Valid values are NONE for open access, AWS_IAM
	// for using AWS IAM permissions, CUSTOM for using a custom authorizer, or COGNITO_USER_POOLS
	// for using a Cognito user pool.
	AuthorizationType *string `locationName:"authorizationType" type:"string"`

	// The identifier of an Authorizer to use on this method. The authorizationType
	// must be CUSTOM.
	AuthorizerId *string `locationName:"authorizerId" type:"string"`

	// The method's HTTP verb.
	HttpMethod *string `locationName:"httpMethod" type:"string"`

	// Gets the method's integration responsible for passing the client-submitted
	// request to the back end and performing necessary transformations to make
	// the request compliant with the back end.
	//
	// Example:
	//
	// Request
	//   GET /restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration HTTP/1.1
	//   Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
	//   Content-Length: 117 X-Amz-Date: 20160613T213210Z Authorization: AWS4-HMAC-SHA256
	//   Credential={access_key_ID}/20160613/us-east-1/apigateway/aws4_request,
	//   SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	// Response
	//
	// The successful response returns a 200 OK status code and a payload similar
	// to the following:
	//  { "_links": { "curies": [ { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-{rel}.html",
	//  "name": "integration", "templated": true }, { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
	//  "name": "integrationresponse", "templated": true } ], "self": { "href":
	//  "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration" }, "integration:delete":
	//  { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
	//  }, "integration:responses": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200",
	//  "name": "200", "title": "200" }, "integration:update": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration"
	//  }, "integrationresponse:put": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/{status_code}",
	//  "templated": true } }, "cacheKeyParameters": [], "cacheNamespace": "0cjtch",
	//  "credentials": "arn:aws:iam::123456789012:role/apigAwsProxyRole", "httpMethod":
	//  "POST", "passthroughBehavior": "WHEN_NO_MATCH", "requestTemplates": { "application/json":
	//  "{\n \"a\": \"$input.params('operand1')\",\n \"b\": \"$input.params('operand2')\",
	//  \n \"op\": \"$input.params('operator')\" \n}" }, "type": "AWS", "uri": "arn:aws:apigateway:us-west-2:lambda:path//2015-03-31/functions/arn:aws:lambda:us-west-2:123456789012:function:Calc/invocations",
	//  "_embedded": { "integration:responses": { "_links": { "self": { "href":
	//  "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200",
	//  "name": "200", "title": "200" }, "integrationresponse:delete": { "href":
	//  "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200"
	//  }, "integrationresponse:update": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/integration/responses/200"
	//  } }, "responseParameters": { "method.response.header.operator": "integration.response.body.op",
	//  "method.response.header.operand_2": "integration.response.body.b", "method.response.header.operand_1":
	//  "integration.response.body.a" }, "responseTemplates": { "application/json":
	//  "#set($res = $input.path('$'))\n{\n \"result\": \"$res.a, $res.b, $res.op
	//  => $res.c\",\n \"a\" : \"$res.a\",\n \"b\" : \"$res.b\",\n \"op\" : \"$res.op\",\n
	//  \"c\" : \"$res.c\"\n}" }, "selectionPattern": "", "statusCode": "200" }
	//  } }
	//
	// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-integration.html)
	MethodIntegration *Integration `locationName:"methodIntegration" type:"structure"`

	// Gets a method response associated with a given HTTP status code.
	//
	// The collection of method responses are encapsulated in a key-value map, where
	// the key is a response's HTTP status code and the value is a MethodResponse
	// resource that specifies the response returned to the caller from the back
	// end through the integration response.
	//
	// Example: Get a 200 OK response of a GET method
	//
	// Request
	//   GET /restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200 HTTP/1.1
	//   Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
	//   Content-Length: 117 X-Amz-Date: 20160613T215008Z Authorization: AWS4-HMAC-SHA256
	//   Credential={access_key_ID}/20160613/us-east-1/apigateway/aws4_request,
	//   SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	// Response
	//
	// The successful response returns a 200 OK status code and a payload similar
	// to the following:
	//  { "_links": { "curies": { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
	//  "name": "methodresponse", "templated": true }, "self": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200",
	//  "title": "200" }, "methodresponse:delete": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200"
	//  }, "methodresponse:update": { "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200"
	//  } }, "responseModels": { "application/json": "Empty" }, "responseParameters":
	//  { "method.response.header.operator": false, "method.response.header.operand_2":
	//  false, "method.response.header.operand_1": false }, "statusCode": "200"
	//  }
	//
	// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-method-response.html)
	MethodResponses map[string]MethodResponse `locationName:"methodResponses" type:"map"`

	// A human-friendly operation identifier for the method. For example, you can
	// assign the operationName of ListPets for the GET /pets method in PetStore
	// (https://petstore-demo-endpoint.execute-api.com/petstore/pets) example.
	OperationName *string `locationName:"operationName" type:"string"`

	// A key-value map specifying data schemas, represented by Model resources,
	// (as the mapped value) of the request payloads of given content types (as
	// the mapping key).
	RequestModels map[string]string `locationName:"requestModels" type:"map"`

	// A key-value map defining required or optional method request parameters that
	// can be accepted by API Gateway. A key is a method request parameter name
	// matching the pattern of method.request.{location}.{name}, where location
	// is querystring, path, or header and name is a valid and unique parameter
	// name. The value associated with the key is a Boolean flag indicating whether
	// the parameter is required (true) or optional (false). The method request
	// parameter names defined here are available in Integration to be mapped to
	// integration request parameters or templates.
	RequestParameters map[string]bool `locationName:"requestParameters" type:"map"`

	// The identifier of a RequestValidator for request validation.
	RequestValidatorId *string `locationName:"requestValidatorId" type:"string"`
}

// String returns the string representation
func (s GetMethodOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMethodOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiKeyRequired != nil {
		v := *s.ApiKeyRequired

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeyRequired", protocol.BoolValue(v), metadata)
	}
	if s.AuthorizationScopes != nil {
		v := s.AuthorizationScopes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "authorizationScopes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AuthorizationType != nil {
		v := *s.AuthorizationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizationType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerId != nil {
		v := *s.AuthorizerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "httpMethod", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MethodIntegration != nil {
		v := s.MethodIntegration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "methodIntegration", v, metadata)
	}
	if s.MethodResponses != nil {
		v := s.MethodResponses

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "methodResponses", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.OperationName != nil {
		v := *s.OperationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "operationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestModels != nil {
		v := s.RequestModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RequestParameters != nil {
		v := s.RequestParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.BoolValue(v1))
		}
		ms0.End()

	}
	if s.RequestValidatorId != nil {
		v := *s.RequestValidatorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestValidatorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetMethod = "GetMethod"

// GetMethodRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describe an existing Method resource.
//
//    // Example sending a request using GetMethodRequest.
//    req := client.GetMethodRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetMethodRequest(input *GetMethodInput) GetMethodRequest {
	op := &aws.Operation{
		Name:       opGetMethod,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}",
	}

	if input == nil {
		input = &GetMethodInput{}
	}

	req := c.newRequest(op, input, &GetMethodOutput{})
	return GetMethodRequest{Request: req, Input: input, Copy: c.GetMethodRequest}
}

// GetMethodRequest is the request type for the
// GetMethod API operation.
type GetMethodRequest struct {
	*aws.Request
	Input *GetMethodInput
	Copy  func(*GetMethodInput) GetMethodRequest
}

// Send marshals and sends the GetMethod API request.
func (r GetMethodRequest) Send(ctx context.Context) (*GetMethodResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMethodResponse{
		GetMethodOutput: r.Request.Data.(*GetMethodOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMethodResponse is the response type for the
// GetMethod API operation.
type GetMethodResponse struct {
	*GetMethodOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMethod request.
func (r *GetMethodResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
