// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration for a Lambda function URL.
func (c *Client) UpdateFunctionUrlConfig(ctx context.Context, params *UpdateFunctionUrlConfigInput, optFns ...func(*Options)) (*UpdateFunctionUrlConfigOutput, error) {
	if params == nil {
		params = &UpdateFunctionUrlConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunctionUrlConfig", params, optFns, c.addOperationUpdateFunctionUrlConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionUrlConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionUrlConfigInput struct {

	// The name of the Lambda function. Name formats
	//   - Function name – my-function .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The type of authentication that your function URL uses. Set to AWS_IAM if you
	// want to restrict access to authenticated users only. Set to NONE if you want to
	// bypass IAM authentication to create a public endpoint. For more information, see
	// Security and auth model for Lambda function URLs (https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html)
	// .
	AuthType types.FunctionUrlAuthType

	// The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
	// settings for your function URL.
	Cors *types.Cors

	// Use one of the following options:
	//   - BUFFERED – This is the default option. Lambda invokes your function using
	//   the Invoke API operation. Invocation results are available when the payload is
	//   complete. The maximum payload size is 6 MB.
	//   - RESPONSE_STREAM – Your function streams payload results as they become
	//   available. Lambda invokes your function using the InvokeWithResponseStream API
	//   operation. The maximum response payload size is 20 MB, however, you can
	//   request a quota increase (https://docs.aws.amazon.com/servicequotas/latest/userguide/request-quota-increase.html)
	//   .
	InvokeMode types.InvokeMode

	// The alias name.
	Qualifier *string

	noSmithyDocumentSerde
}

type UpdateFunctionUrlConfigOutput struct {

	// The type of authentication that your function URL uses. Set to AWS_IAM if you
	// want to restrict access to authenticated users only. Set to NONE if you want to
	// bypass IAM authentication to create a public endpoint. For more information, see
	// Security and auth model for Lambda function URLs (https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html)
	// .
	//
	// This member is required.
	AuthType types.FunctionUrlAuthType

	// When the function URL was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// This member is required.
	CreationTime *string

	// The Amazon Resource Name (ARN) of your function.
	//
	// This member is required.
	FunctionArn *string

	// The HTTP URL endpoint for your function.
	//
	// This member is required.
	FunctionUrl *string

	// When the function URL configuration was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// This member is required.
	LastModifiedTime *string

	// The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
	// settings for your function URL.
	Cors *types.Cors

	// Use one of the following options:
	//   - BUFFERED – This is the default option. Lambda invokes your function using
	//   the Invoke API operation. Invocation results are available when the payload is
	//   complete. The maximum payload size is 6 MB.
	//   - RESPONSE_STREAM – Your function streams payload results as they become
	//   available. Lambda invokes your function using the InvokeWithResponseStream API
	//   operation. The maximum response payload size is 20 MB, however, you can
	//   request a quota increase (https://docs.aws.amazon.com/servicequotas/latest/userguide/request-quota-increase.html)
	//   .
	InvokeMode types.InvokeMode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFunctionUrlConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFunctionUrlConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFunctionUrlConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFunctionUrlConfig"); err != nil {
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
	if err = addOpUpdateFunctionUrlConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunctionUrlConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFunctionUrlConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFunctionUrlConfig",
	}
}
