// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Simulate the invocation of a Method in your RestApi with headers, parameters,
// and an incoming request body.
func (c *Client) TestInvokeMethod(ctx context.Context, params *TestInvokeMethodInput, optFns ...func(*Options)) (*TestInvokeMethodOutput, error) {
	if params == nil {
		params = &TestInvokeMethodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestInvokeMethod", params, optFns, c.addOperationTestInvokeMethodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestInvokeMethodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Make a request to simulate the invocation of a Method.
type TestInvokeMethodInput struct {

	// Specifies a test invoke method request's HTTP method.
	//
	// This member is required.
	HttpMethod *string

	// Specifies a test invoke method request's resource ID.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The simulated request body of an incoming invocation request.
	Body *string

	// A ClientCertificate identifier to use in the test invocation. API Gateway will
	// use the certificate when making the HTTPS request to the defined back-end
	// endpoint.
	ClientCertificateId *string

	// A key-value map of headers to simulate an incoming invocation request.
	Headers map[string]string

	// The headers as a map from string to list of values to simulate an incoming
	// invocation request.
	MultiValueHeaders map[string][]string

	// The URI path, including query string, of the simulated invocation request. Use
	// this to specify path parameters and query string parameters.
	PathWithQueryString *string

	// A key-value map of stage variables to simulate an invocation on a deployed
	// Stage.
	StageVariables map[string]string

	noSmithyDocumentSerde
}

// Represents the response of the test invoke request in the HTTP method.
type TestInvokeMethodOutput struct {

	// The body of the HTTP response.
	Body *string

	// The headers of the HTTP response.
	Headers map[string]string

	// The execution latency, in ms, of the test invoke request.
	Latency int64

	// The API Gateway execution log for the test invoke request.
	Log *string

	// The headers of the HTTP response as a map from string to list of values.
	MultiValueHeaders map[string][]string

	// The HTTP status code.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestInvokeMethodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestInvokeMethod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestInvokeMethod{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TestInvokeMethod"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpTestInvokeMethodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestInvokeMethod(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opTestInvokeMethod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TestInvokeMethod",
	}
}
