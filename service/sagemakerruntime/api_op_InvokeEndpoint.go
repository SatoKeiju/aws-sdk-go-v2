// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// After you deploy a model into production using Amazon SageMaker hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint. For an overview of Amazon SageMaker, see How
// It Works (https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html) .
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax. Calls to
// InvokeEndpoint are authenticated by using Amazon Web Services Signature Version
// 4. For information, see Authenticating Requests (Amazon Web Services Signature
// Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference. A customer's model containers must respond to
// requests within 60 seconds. The model itself can have a maximum processing time
// of 60 seconds before responding to invocations. If your model is going to take
// 50-60 seconds of processing time, the SDK socket timeout should be set to be 70
// seconds. Endpoints are scoped to an individual account, and are not public. The
// URL does not contain the account ID, but Amazon SageMaker determines the account
// ID from the authentication token that is supplied by the caller.
func (c *Client) InvokeEndpoint(ctx context.Context, params *InvokeEndpointInput, optFns ...func(*Options)) (*InvokeEndpointOutput, error) {
	if params == nil {
		params = &InvokeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeEndpoint", params, optFns, c.addOperationInvokeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeEndpointInput struct {

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model. For
	// information about the format of the request body, see Common Data
	// Formats-Inference (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html)
	// .
	//
	// This member is required.
	Body []byte

	// The name of the endpoint that you specified when you created the endpoint using
	// the CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html)
	// API.
	//
	// This member is required.
	EndpointName *string

	// The desired MIME type of the inference response from the model container.
	Accept *string

	// The MIME type of the input data in the request body.
	ContentType *string

	// Provides additional information about a request for an inference submitted to a
	// model hosted at an Amazon SageMaker endpoint. The information is an opaque value
	// that is forwarded verbatim. You could use this value, for example, to provide an
	// ID that you can use to track a request or to provide other metadata that a
	// service endpoint was programmed to process. The value must consist of no more
	// than 1024 visible US-ASCII characters as specified in Section 3.3.6. Field
	// Value Components (https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.6)
	// of the Hypertext Transfer Protocol (HTTP/1.1). The code in your model is
	// responsible for setting or updating any custom attributes in the response. If
	// your code does not set this value in the response, an empty value is returned.
	// For example, if a custom attribute represents the trace ID, your model can
	// prepend the custom attribute with Trace ID: in your post-processing function.
	// This feature is currently supported in the Amazon Web Services SDKs but not in
	// the Amazon SageMaker Python SDK.
	CustomAttributes *string

	// An optional JMESPath expression used to override the EnableExplanations
	// parameter of the ClarifyExplainerConfig API. See the EnableExplanations (https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-online-explainability-create-endpoint.html#clarify-online-explainability-create-endpoint-enable)
	// section in the developer guide for more information.
	EnableExplanations *string

	// If you provide a value, it is added to the captured data when you enable data
	// capture on the endpoint. For information about data capture, see Capture Data (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor-data-capture.html)
	// .
	InferenceId *string

	// If the endpoint hosts multiple containers and is configured to use direct
	// invocation, this parameter specifies the host name of the container to invoke.
	TargetContainerHostname *string

	// The model to request for inference when invoking a multi-model endpoint.
	TargetModel *string

	// Specify the production variant to send the inference request to when invoking
	// an endpoint that is running two or more variants. Note that this parameter
	// overrides the default behavior for the endpoint, which is to distribute the
	// invocation traffic based on the variant weights. For information about how to
	// use variant targeting to perform a/b testing, see Test models in production (https://docs.aws.amazon.com/sagemaker/latest/dg/model-ab-testing.html)
	TargetVariant *string

	noSmithyDocumentSerde
}

type InvokeEndpointOutput struct {

	// Includes the inference provided by the model. For information about the format
	// of the response body, see Common Data Formats-Inference (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html)
	// . If the explainer is activated, the body includes the explanations provided by
	// the model. For more information, see the Response section under Invoke the
	// Endpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-online-explainability-invoke-endpoint.html#clarify-online-explainability-response)
	// in the Developer Guide.
	//
	// This member is required.
	Body []byte

	// The MIME type of the inference returned from the model container.
	ContentType *string

	// Provides additional information in the response about the inference returned by
	// a model hosted at an Amazon SageMaker endpoint. The information is an opaque
	// value that is forwarded verbatim. You could use this value, for example, to
	// return an ID received in the CustomAttributes header of a request or other
	// metadata that a service endpoint was programmed to produce. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in
	// Section 3.3.6. Field Value Components (https://tools.ietf.org/html/rfc7230#section-3.2.6)
	// of the Hypertext Transfer Protocol (HTTP/1.1). If the customer wants the custom
	// attribute returned, the model must set the custom attribute to be included on
	// the way back. The code in your model is responsible for setting or updating any
	// custom attributes in the response. If your code does not set this value in the
	// response, an empty value is returned. For example, if a custom attribute
	// represents the trace ID, your model can prepend the custom attribute with Trace
	// ID: in your post-processing function. This feature is currently supported in the
	// Amazon Web Services SDKs but not in the Amazon SageMaker Python SDK.
	CustomAttributes *string

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInvokeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeEndpoint"); err != nil {
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
	if err = addOpInvokeEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeEndpoint(options.Region), middleware.Before); err != nil {
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
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opInvokeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeEndpoint",
	}
}
